package mapping

import (
	"auth/model"
	"auth/model/entity"
)

// GroupsEntityToDto entity data transfer
func GroupsEntityToDto(groups []*entity.Group) []*model.GroupInfo {
	out := make([]*model.GroupInfo, 0, len(groups))
	for _, c := range groups {
		out = append(out, GroupEntityToDto(c))
	}
	return out
}

// GroupEntityToDto entity data transfer
func GroupEntityToDto(e *entity.Group) *model.GroupInfo {
	return &model.GroupInfo{

		Id: e.Id,

		Name: e.Name,

		ParentId: e.ParentId,

		Roles: e.Roles,

		CreatedAt: e.CreatedAt.Unix(),

		UpdatedAt: e.UpdatedAt.Unix(),
	}
}
