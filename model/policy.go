package model

type PolicyCreateRequest struct {
	Id         int64   `json:"id"`
	Ptype      string  `binding:"required" json:"ptype"`
	RoleId     int64   `binding:"required" json:"role_id"`
	ResourceId int64   `binding:"required" json:"resource_id"`
	Operate    int     `binding:"required" json:"operate"`
	V3         *string `json:"v3"`
	V4         *string `json:"v4"`
	V5         *string `json:"v5"`
}
type PolicyUpdateRequest struct {
	Id         int64   `json:"id"`
	Ptype      string  `binding:"required" json:"ptype"`
	RoleId     int64   `binding:"required" json:"role_id"`
	ResourceId int64   `binding:"required" json:"resource_id"`
	Operate    int     `binding:"required" json:"operate"`
	V3         *string `json:"v3"`
	V4         *string `json:"v4"`
	V5         *string `json:"v5"`
}
type PolicyListRequest struct {
	Id         int64   `json:"id"`
	Ptype      string  `binding:"required" json:"ptype"`
	RoleId     int64   `binding:"required" json:"role_id"`
	ResourceId int64   `binding:"required" json:"resource_id"`
	Operate    int     `binding:"required" json:"operate"`
	V3         *string `json:"v3"`
	V4         *string `json:"v4"`
	V5         *string `json:"v5"`
	Index      int     `json:"index"`
	Size       int     `json:"size"`
}
type PolicyListResponse struct {
	Total int           `json:"total"`
	List  []*PolicyInfo `json:"list"`
}
type PolicyInfoRequest struct {
	Id int64 `json:"id"`
}
type PolicyInfo struct {
	Id         int64  `json:"id"`
	Ptype      string `binding:"required" json:"ptype"`
	RoleId     int64  `binding:"required" json:"role_id"`
	ResourceId int64  `binding:"required" json:"resource_id"`
	Operate    int    `binding:"required" json:"operate"`
	V3         string `json:"v3"`
	V4         string `json:"v4"`
	V5         string `json:"v5"`
}
type PolicyDeleteRequest struct {
	Id int64 `json:"id"`
}
