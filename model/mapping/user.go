package mapping

import (
	"auth/model"
	"auth/model/entity"
)

// UsersEntityToDto entity data transfer
func UsersEntityToDto(users []*entity.User) []*model.UserInfo {
	out := make([]*model.UserInfo, 0, len(users))
	for _, c := range users {
		out = append(out, UserEntityToDto(c))
	}
	return out
}

// UserEntityToDto entity data transfer
func UserEntityToDto(e *entity.User) *model.UserInfo {
	return &model.UserInfo{

		Id: e.Id,

		Name: e.Name,

		Phone: e.Phone,

		WxId: e.WxId,

		Password: e.Password,

		Salt: e.Salt,

		Type: e.Type,

		Avator: e.Avator,

		Roles: e.Roles,

		Groups: e.Groups,

		CreatedAt: e.CreatedAt.Unix(),

		UpdatedAt: e.UpdatedAt.Unix(),
	}
}
