package store

import (
	"auth/model"
	"auth/model/entity"
	"context"
)

type IUser interface {
	// Create
	Create(ctx context.Context, e *entity.User) (int64, error)
	// Find
	Find(ctx context.Context, in *model.UserInfoRequest) (*entity.User, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.UserListRequest) (int, []*entity.User, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
