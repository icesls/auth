package postgres

import (
	log "auth/collector/logger"
	"auth/config"
	"auth/errors"
	"auth/model"
	"auth/model/entity"
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Role = &role{}

type role struct{}

func init() {
	Register(Role)
}

func (a *role) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.Role{}
		if db.Migrator().HasTable(p) {
			log.Debug("table already exist: ", zap.String("table", p.TableName()))
			return
		}
		if err := db.AutoMigrate(p); err != nil {
			log.Error("filed to create table please check config or manually create", zap.String("table", p.TableName()), zap.String("err", err.Error()))
		} else {
			log.Info("create table successfully", zap.String("table", p.TableName()))
		}
	}
}

// Create
func (a *role) Create(ctx context.Context, m *entity.Role) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *role) Find(ctx context.Context, in *model.RoleInfoRequest) (*entity.Role, error) {
	e := &entity.Role{}

	q := GetDB(ctx).Model(&entity.Role{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update
func (a *role) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.Role{}).Where("id = ?", id).Updates(dict).Error
}

// Delete
func (a *role) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.Role{}, id).Error
}

// List query list
func (a *role) List(ctx context.Context, in *model.RoleListRequest) (int, []*entity.Role, error) {
	var (
		q     = GetDB(ctx).Model(&entity.Role{})
		err   error
		total int64
		roles []*entity.Role
	)

	if in.CreatedAt != nil {

		q = q.Where("created_at = ?", in.CreatedAt)

	}

	if in.UpdatedAt != nil {

		q = q.Where("updated_at = ?", in.UpdatedAt)

	}

	if err = q.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&roles).Error; err != nil {
		return 0, nil, err
	}
	return int(total), roles, nil
}

// ExecTransaction execute database transaction
func (a *role) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
