package bll

import (
	"context"

	"auth/model"
	"auth/model/entity"
	"auth/model/mapping"
	"auth/store"
	"auth/store/postgres"
	"time"
)

type role struct {
	iRole store.IRole
}

var Role = &role{
	iRole: postgres.Role,
}

func init() {
	Register(Role)
}

func (a *role) init() func() {
	return func() {}
}

// Create
func (a *role) Create(ctx context.Context, in *model.RoleCreateRequest) error {
	var (
		err error
	)
	c := buildRole(in)
	_, err = a.iRole.Create(ctx, c)
	return err
}

// Update
func (a *role) Update(ctx context.Context, in *model.RoleUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.CreatedAt != nil {
		dict["created_at"] = in.CreatedAt
	}

	if in.UpdatedAt != nil {
		dict["updated_at"] = in.UpdatedAt
	}

	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.iRole.Update(ctx, in.Id, dict)
}

// Delete
func (a *role) Delete(ctx context.Context, in *model.RoleDeleteRequest) error {
	return a.iRole.Delete(ctx, in.Id)
}

// List
func (a *role) List(ctx context.Context, in *model.RoleListRequest) (*model.RoleListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.Role
		out   = &model.RoleListResponse{}
	)

	if total, list, err = a.iRole.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.RolesEntityToDto(list)

	return out, nil
}

// Find
func (a *role) Find(ctx context.Context, in *model.RoleInfoRequest) (*model.RoleInfo, error) {
	var (
		err  error
		data *entity.Role
		out  = &model.RoleInfo{}
	)

	if data, err = a.iRole.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.RoleEntityToDto(data)
	return out, nil
}

// buildRole build entity
func buildRole(in *model.RoleCreateRequest) *entity.Role {

	now := time.Now()

	ety := &entity.Role{

		Name: in.Name,

		Resources: in.Resources,

		CreatedAt: now,

		UpdatedAt: now,
	}

	return ety
}
