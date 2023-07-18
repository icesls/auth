package mapping

import (
	"auth/model"
	"auth/model/entity"
)

// PoliciesEntityToDto entity data transfer
func PoliciesEntityToDto(policies []*entity.Policy) []*model.PolicyInfo {
	out := make([]*model.PolicyInfo, 0, len(policies))
	for _, c := range policies {
		out = append(out, PolicyEntityToDto(c))
	}
	return out
}

// PolicyEntityToDto entity data transfer
func PolicyEntityToDto(e *entity.Policy) *model.PolicyInfo {
	return &model.PolicyInfo{
		Id:         e.Id,
		Ptype:      e.Ptype,
		RoleId:     e.RoleId,
		ResourceId: e.ResourceId,
		Operate:    e.Operate,
		V3:         e.V3,
		V4:         e.V4,
		V5:         e.V5,
	}
}
