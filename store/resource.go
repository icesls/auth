package store

import (
	"context"

	"auth/model"
	"auth/model/entity"
)

type IResource interface {
	// Create
	Create(ctx context.Context, e *entity.Resource) (int64, error)

	CreateMany(ctx context.Context, m ...*entity.Resource) error
	// Find
	Find(ctx context.Context, in *model.ResourceInfoRequest) (*entity.Resource, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.ResourceListRequest) (int, []*entity.Resource, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
