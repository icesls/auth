package model

type ResourceCreateRequest struct {
	Id       int64   `json:"id"`
	IsParent bool    `json:"is_parent"`
	Name     string  `binding:"required" json:"name"`
	Type     string  `binding:"required" json:"type"`
	Path     string  `binding:"required" json:"path"`
	Icon     *string `json:"icon"`
	ParentId *int64  `json:"parent_id"`
}

type ResourceUpdateRequest struct {
	Id        int64   `json:"id"`
	Name      string  `binding:"required" json:"name"`
	Type      string  `binding:"required" json:"type"`
	Path      string  `binding:"required" json:"path"`
	Icon      *string `json:"icon"`
	ParentId  *int64  `json:"parent_id"`
	CreatedAt *int64  `json:"created_at"`
	UpdatedAt *int64  `json:"updated_at"`
}

type ResourceListRequest struct {
	Id        int64   `json:"id"`
	Name      string  `binding:"required" json:"name"`
	Type      string  `binding:"required" json:"type"`
	Path      string  `binding:"required" json:"path"`
	Icon      *string `json:"icon"`
	ParentId  *int64  `json:"parent_id"`
	CreatedAt *int64  `json:"created_at"`
	UpdatedAt *int64  `json:"updated_at"`
	Index     int     `json:"index"`
	Size      int     `json:"size"`
}

type ResourceListResponse struct {
	Total int             `json:"total"`
	List  []*ResourceInfo `json:"list"`
}

type ResourceInfoRequest struct {
	Id int64 `json:"id"`
}

type ResourceInfo struct {
	Id        int64  `json:"id"`
	Name      string `binding:"required" json:"name"`
	Type      string `binding:"required" json:"type"`
	Path      string `binding:"required" json:"path"`
	Icon      string `json:"icon"`
	ParentId  int64  `json:"parent_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ResourceDeleteRequest struct {
	Id int64 `json:"id"`
}
