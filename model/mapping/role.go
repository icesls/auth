package mapping

import (
	"auth/model"
	"auth/model/entity"
)

// RolesEntityToDto entity data transfer
func RolesEntityToDto(roles []*entity.Role) []*model.RoleInfo {
	out := make([]*model.RoleInfo, 0, len(roles))
	for _, c := range roles {
		out = append(out, RoleEntityToDto(c))
	}
	return out
}

// RoleEntityToDto entity data transfer
func RoleEntityToDto(e *entity.Role) *model.RoleInfo {
	return &model.RoleInfo{

		Id: e.Id,

		Name: e.Name,

		Resources: e.Resources,

		CreatedAt: e.CreatedAt.Unix(),

		UpdatedAt: e.UpdatedAt.Unix(),
	}
}
