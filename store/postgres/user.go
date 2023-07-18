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

var User = &user{}

type user struct{}

func init() {
	Register(User)
}

func (a *user) Init() {
	if config.Conf.AutoMigrate {
		p := &entity.User{}
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
func (a *user) Create(ctx context.Context, m *entity.User) (int64, error) {
	err := GetDB(ctx).Create(m).Error
	return m.Id, err
}

// Find detail
func (a *user) Find(ctx context.Context, in *model.UserInfoRequest) (*entity.User, error) {
	e := &entity.User{}

	q := GetDB(ctx).Model(&entity.User{})

	if in.Id == 0 {
		return e, errors.New("condition illegal")
	}
	err := q.First(&e).Error
	return e, err
}

// Update
func (a *user) Update(ctx context.Context, id int64, dict map[string]interface{}) error {
	return GetDB(ctx).Model(&entity.User{}).Where("id = ?", id).Updates(dict).Error
}

// Delete
func (a *user) Delete(ctx context.Context, id int64) error {
	return GetDB(ctx).Delete(&entity.User{}, id).Error
}

// List query list
func (a *user) List(ctx context.Context, in *model.UserListRequest) (int, []*entity.User, error) {
	var (
		q     = GetDB(ctx).Model(&entity.User{})
		err   error
		total int64
		users []*entity.User
	)

	if in.Phone != nil {

		q = q.Where("phone like ?", in.Phone)

	}

	if in.WxId != nil {

		q = q.Where("wx_id like ?", in.WxId)

	}

	if in.Avator != nil {

		q = q.Where("avator like ?", in.Avator)

	}

	if in.Roles != nil {

		q = q.Where("roles = ?", in.Roles)

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
	if err = q.Limit(in.Size).Offset((in.Index - 1) * in.Size).Find(&users).Error; err != nil {
		return 0, nil, err
	}
	return int(total), users, nil
}

// ExecTransaction execute database transaction
func (a *user) ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error {
	return GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ContextTxKey, tx)
		return callback(ctx)
	})
}
