package model

import "github.com/lib/pq"

type RoleCreateRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Resources pq.Int64Array `binding:"required" json:"resources"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
}
type RoleUpdateRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Resources pq.Int64Array `binding:"required" json:"resources"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
}
type RoleListRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Resources pq.Int64Array `binding:"required" json:"resources"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
    Index     int           `json:"index"`
    Size      int           `json:"size"`
}
type RoleListResponse struct {
    Total int         `json:"total"`
    List  []*RoleInfo `json:"list"`
}
type RoleInfoRequest struct {
    Id int64 `json:"id"`
}
type RoleInfo struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Resources pq.Int64Array `binding:"required" json:"resources"`
    CreatedAt int64         `json:"created_at"`
    UpdatedAt int64         `json:"updated_at"`
}
type RoleDeleteRequest struct {
    Id int64 `json:"id"`
}
