package bll

import (
	"context"

	"auth/model"
	"auth/model/entity"
	"auth/model/mapping"
	"auth/store"
	"auth/store/postgres"
)

type policy struct {
	iPolicy store.IPolicy
}

var Policy = &policy{
	iPolicy: postgres.Policy,
}

func init() {
	Register(Policy)
}

func (a *policy) init() func() {
	return func() {}
}

// Create
func (a *policy) Create(ctx context.Context, in *model.PolicyCreateRequest) error {
	var (
		err error
	)
	c := buildPolicy(in)
	_, err = a.iPolicy.Create(ctx, c)
	return err
}

// Update
func (a *policy) Update(ctx context.Context, in *model.PolicyUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.V3 != nil {
		dict["v3"] = in.V3
	}

	if in.V4 != nil {
		dict["v4"] = in.V4
	}

	if in.V5 != nil {
		dict["v5"] = in.V5
	}

	// do other update here
	return a.iPolicy.Update(ctx, in.Id, dict)
}

// Delete
func (a *policy) Delete(ctx context.Context, in *model.PolicyDeleteRequest) error {
	return a.iPolicy.Delete(ctx, in.Id)
}

// List
func (a *policy) List(ctx context.Context, in *model.PolicyListRequest) (*model.PolicyListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.Policy
		out   = &model.PolicyListResponse{}
	)

	if total, list, err = a.iPolicy.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.PoliciesEntityToDto(list)

	return out, nil
}

// Find
func (a *policy) Find(ctx context.Context, in *model.PolicyInfoRequest) (*model.PolicyInfo, error) {
	var (
		err  error
		data *entity.Policy
		out  = &model.PolicyInfo{}
	)

	if data, err = a.iPolicy.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.PolicyEntityToDto(data)
	return out, nil
}

// buildPolicy build entity
func buildPolicy(in *model.PolicyCreateRequest) *entity.Policy {

	ety := &entity.Policy{

		Ptype: in.Ptype,

		RoleId: in.RoleId,

		ResourceId: in.ResourceId,

		Operate: in.Operate,
	}

	if in.V3 != nil {
		ety.V3 = *in.V3
	}

	if in.V4 != nil {
		ety.V4 = *in.V4
	}

	if in.V5 != nil {
		ety.V5 = *in.V5
	}

	return ety
}
