package mapping

import (
	"auth/model"
	"auth/model/entity"
)

// ResourcesEntityToDto entity data transfer
func ResourcesEntityToDto(resources []*entity.Resource) []*model.ResourceInfo {
	out := make([]*model.ResourceInfo, 0, len(resources))
	for _, c := range resources {
		out = append(out, ResourceEntityToDto(c))
	}
	return out
}

// ResourceEntityToDto entity data transfer
func ResourceEntityToDto(e *entity.Resource) *model.ResourceInfo {
	return &model.ResourceInfo{

		Id: e.Id,

		Name: e.Name,

		Type: e.Type,

		Path: e.Path,

		Icon: e.Icon,

		ParentId: e.ParentId,

		CreatedAt: e.CreatedAt.Unix(),

		UpdatedAt: e.UpdatedAt.Unix(),
	}
}
