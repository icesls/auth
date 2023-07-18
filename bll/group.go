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

type group struct {
	iGroup store.IGroup
}

var Group = &group{
	iGroup: postgres.Group,
}

func init() {
	Register(Group)
}

func (a *group) init() func() {
	return func() {}
}

// Create
func (a *group) Create(ctx context.Context, in *model.GroupCreateRequest) error {
	var (
		err error
	)
	c := buildGroup(in)
	_, err = a.iGroup.Create(ctx, c)
	return err
}

// Update
func (a *group) Update(ctx context.Context, in *model.GroupUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.ParentId != nil {
		dict["parent_id"] = in.ParentId
	}

	if in.CreatedAt != nil {
		dict["created_at"] = in.CreatedAt
	}

	if in.UpdatedAt != nil {
		dict["updated_at"] = in.UpdatedAt
	}

	// do other update here
	updateAt := time.Now().Unix()
	in.UpdatedAt = &updateAt
	return a.iGroup.Update(ctx, in.Id, dict)
}

// Delete
func (a *group) Delete(ctx context.Context, in *model.GroupDeleteRequest) error {
	return a.iGroup.Delete(ctx, in.Id)
}

// List
func (a *group) List(ctx context.Context, in *model.GroupListRequest) (*model.GroupListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.Group
		out   = &model.GroupListResponse{}
	)

	if total, list, err = a.iGroup.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.GroupsEntityToDto(list)

	return out, nil
}

// Find
func (a *group) Find(ctx context.Context, in *model.GroupInfoRequest) (*model.GroupInfo, error) {
	var (
		err  error
		data *entity.Group
		out  = &model.GroupInfo{}
	)

	if data, err = a.iGroup.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.GroupEntityToDto(data)
	return out, nil
}

// buildGroup build entity
func buildGroup(in *model.GroupCreateRequest) *entity.Group {

	now := time.Now()

	ety := &entity.Group{

		Name: in.Name,

		Roles: in.Roles,

		CreatedAt: now,

		UpdatedAt: now,
	}

	if in.ParentId != nil {
		ety.ParentId = *in.ParentId
	}

	return ety
}
