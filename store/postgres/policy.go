package postgres

import (
	"context"

	log "auth/collector/logger"
	"auth/config"
	"auth/errors"
	"auth/model"
	"auth/model/entity"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Policy = &policy{}

type policy struct{}

func init() {
	Register(Policy)
}

func (a *policy) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.Policy{}
		if db.Migrator().HasTable(p) {
			log.Debug("table already exist: ", zap.String("table", p.TableName()))
			return
		}

		if err := db.AutoMigrate(p); err != nil {
			log.Error(
				"filed to create table please check config or manually create",
				zap.String("table", p.TableName()),
				zap.String("err", err.Error()),
			)
		} else {
			log.Info("create table successfully", zap.String("table", p.TableName()))
		}
	}
}

func (a *policy) Adapter() (*gormadapter.Adapter, error) {
	return gormadapter.NewAdapterByDBWithCustomTable(db, &entity.Policy{}, "policies")
}

func (a *policy) Create(ctx context.Context, m *entity.Policy) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *policy) Find(ctx context.Context, in *model.PolicyInfoRequest) (*entity.Policy, error) {
	e := &entity.Policy{}

	q := GetDB(ctx).Model(&entity.Policy{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

func (a *policy) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.Policy{}).Where("id = ?", id).Updates(dict).Error
}

func (a *policy) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.Policy{}, id).Error
}

// List query list
func (a *policy) List(ctx context.Context, in *model.PolicyListRequest) (int, []*entity.Policy, error) {
	var (
		q       = GetDB(ctx).Model(&entity.Policy{})
		err     error
		total   int64
		policys []*entity.Policy
	)

	if in.V3 != nil {
		q = q.Where("v3 like ?", in.V3)
	}

	if in.V4 != nil {
		q = q.Where("v4 like ?", in.V4)
	}

	if in.V5 != nil {
		q = q.Where("v5 like ?", in.V5)
	}

	if err = q.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&policys).Error; err != nil {
		return 0, nil, err
	}
	return int(total), policys, nil
}

// ExecTransaction execute database transaction
func (a *policy) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
