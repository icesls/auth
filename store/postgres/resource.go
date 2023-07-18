package postgres

import (
	"context"

	log "auth/collector/logger"
	"auth/config"
	"auth/errors"
	"auth/model"
	"auth/model/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Resource = &resource{}

type resource struct{}

func init() {
	Register(Resource)
}

func (a *resource) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.Resource{}
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
func (a *resource) Create(ctx context.Context, m *entity.Resource) (int64, error) {
	err := GetDB(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(m).Error

	return m.Id, err
}

func (a *resource) CreateMany(ctx context.Context, m ...*entity.Resource) error {
	err := GetDB(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoNothing: true,
	}).Create(m).Error

	return err
}

// Find detail
func (a *resource) Find(ctx context.Context, in *model.ResourceInfoRequest) (*entity.Resource, error) {
	e := &entity.Resource{}

	q := GetDB(ctx).Model(&entity.Resource{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update
func (a *resource) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.Resource{}).Where("id = ?", id).Updates(dict).Error
}

// Delete
func (a *resource) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.Resource{}, id).Error
}

// List query list
func (a *resource) List(ctx context.Context, in *model.ResourceListRequest) (int, []*entity.Resource, error) {
	var (
		q         = GetDB(ctx).Model(&entity.Resource{})
		err       error
		total     int64
		resources []*entity.Resource
	)

	if in.Icon != nil {

		q = q.Where("icon like ?", in.Icon)

	}

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
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&resources).Error; err != nil {
		return 0, nil, err
	}
	return int(total), resources, nil
}

// ExecTransaction execute database transaction
func (a *resource) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
