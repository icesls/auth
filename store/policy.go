package store

import (
	"auth/model"
	"auth/model/entity"
	"context"
)

type IPolicy interface {
	// Create
	Create(ctx context.Context, e *entity.Policy) (int64, error)
	// Find
	Find(ctx context.Context, in *model.PolicyInfoRequest) (*entity.Policy, error)
	// Update
	Update(ctx context.Context, id int64, updates map[string]interface{}) error
	// Delete
	Delete(ctx context.Context, id int64) error
	// List
	List(ctx context.Context, in *model.PolicyListRequest) (int, []*entity.Policy, error)
	// ExecTransaction
	ExecTransaction(ctx context.Context, callback func(ctx context.Context) error) error
}
