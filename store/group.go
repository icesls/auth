package store

import (
	"auth/model"
	"auth/model/entity"
	"context"
)

type IGroup interface {
	// Create
	Create(ctx context.Context, e *entity.Group) (int64, error)
	// Find
	Find(ctx context.Context, in *model.GroupInfoRequest) (*entity.Group, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.GroupListRequest) (int, []*entity.Group, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
