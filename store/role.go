package store

import (
	"auth/model"
	"auth/model/entity"
	"context"
)

type IRole interface {
	// Create
	Create(ctx context.Context, e *entity.Role) (int64, error)
	// Find
	Find(ctx context.Context, in *model.RoleInfoRequest) (*entity.Role, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.RoleListRequest) (int, []*entity.Role, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
