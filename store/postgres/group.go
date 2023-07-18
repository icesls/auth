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

var Group = &group{}

type group struct{}

func init() {
	Register(Group)
}

func (a *group) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.Group{}
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
func (a *group) Create(ctx context.Context, m *entity.Group) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *group) Find(ctx context.Context, in *model.GroupInfoRequest) (*entity.Group, error) {
	e := &entity.Group{}

	q := GetDB(ctx).Model(&entity.Group{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update
func (a *group) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.Group{}).Where("id = ?", id).Updates(dict).Error
}

// Delete
func (a *group) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.Group{}, id).Error
}

// List query list
func (a *group) List(ctx context.Context, in *model.GroupListRequest) (int, []*entity.Group, error) {
	var (
		q      = GetDB(ctx).Model(&entity.Group{})
		err    error
		total  int64
		groups []*entity.Group
	)

	if in.ParentId != nil {

		q = q.Where("parent_id = ?", in.ParentId)

	}

	if in.CreatedAt != nil {

		q = q.Where("created_at = ?", in.CreatedAt)

	}

	if in.UpdatedAt != nil {

		q = q.Where("updated_at = ?", in.UpdatedAt)

	}

	if err = q.Count(&total).Error; err != nil {
		return 0, nil, err
	}
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&groups).Error; err != nil {
		return 0, nil, err
	}
	return int(total), groups, nil
}

// ExecTransaction execute database transaction
func (a *group) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
