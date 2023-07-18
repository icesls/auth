package bll

import (
	"context"

	"github.com/lib/pq"

	"time"

	"auth/model"
	"auth/model/entity"
	"auth/model/mapping"
	"auth/store"
	"auth/store/postgres"
)

type user struct {
	iUser store.IUser
}

var User = &user{
	iUser: postgres.User,
}

func init() {
	Register(User)
}

func (a *user) init() func() {
	return func() {}
}

// Create
func (a *user) Create(ctx context.Context, in *model.UserCreateRequest) error {
	var (
		err error
	)
	c := buildUser(in)
	_, err = a.iUser.Create(ctx, c)
	return err
}

// Update
func (a *user) Update(ctx context.Context, in *model.UserUpdateRequest) error {
	var (
		dict = make(map[string]interface{})
	)

	if in.Phone != nil {
		dict["phone"] = in.Phone
	}

	if in.WxId != nil {
		dict["wx_id"] = in.WxId
	}

	if in.Avator != nil {
		dict["avator"] = in.Avator
	}

	if in.Roles != nil {
		dict["roles"] = in.Roles
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
	return a.iUser.Update(ctx, in.Id, dict)
}

// Delete
func (a *user) Delete(ctx context.Context, in *model.UserDeleteRequest) error {
	return a.iUser.Delete(ctx, in.Id)
}

// List
func (a *user) List(ctx context.Context, in *model.UserListRequest) (*model.UserListResponse, error) {
	var (
		err   error
		total int
		list  []*entity.User
		out   = &model.UserListResponse{}
	)

	if total, list, err = a.iUser.List(ctx, in); err != nil {
		return nil, err
	}

	out.Total = total
	out.List = mapping.UsersEntityToDto(list)

	return out, nil
}

// Find
func (a *user) Find(ctx context.Context, in *model.UserInfoRequest) (*model.UserInfo, error) {
	var (
		err  error
		data *entity.User
		out  = &model.UserInfo{}
	)

	if data, err = a.iUser.Find(ctx, in); err != nil {
		return nil, err
	}

	out = mapping.UserEntityToDto(data)
	return out, nil
}

// buildUser build entity
func buildUser(in *model.UserCreateRequest) *entity.User {

	now := time.Now()

	ety := &entity.User{

		Name: in.Name,

		Password: in.Password,

		Salt: "",

		Type: in.Type,

		Groups: in.Groups,

		CreatedAt: now,

		UpdatedAt: now,
	}

	if in.Phone != nil {
		ety.Phone = *in.Phone
	}

	if in.WxId != nil {
		ety.WxId = *in.WxId
	}

	if in.Avator != nil {
		ety.Avator = *in.Avator
	}

	if len(in.Roles) != 0 {
		ety.Roles = in.Roles
	} else {
		ety.Roles = pq.Int64Array{}
	}

	return ety
}
