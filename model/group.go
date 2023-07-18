package model

import "github.com/lib/pq"

type GroupCreateRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    ParentId  *int64        `json:"parent_id"`
    Roles     pq.Int64Array `binding:"required" json:"roles"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
}
type GroupUpdateRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    ParentId  *int64        `json:"parent_id"`
    Roles     pq.Int64Array `binding:"required" json:"roles"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
}
type GroupListRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    ParentId  *int64        `json:"parent_id"`
    Roles     pq.Int64Array `binding:"required" json:"roles"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
    Index     int           `json:"index"`
    Size      int           `json:"size"`
}
type GroupListResponse struct {
    Total int          `json:"total"`
    List  []*GroupInfo `json:"list"`
}
type GroupInfoRequest struct {
    Id int64 `json:"id"`
}
type GroupInfo struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    ParentId  int64         `json:"parent_id"`
    Roles     pq.Int64Array `binding:"required" json:"roles"`
    CreatedAt int64         `json:"created_at"`
    UpdatedAt int64         `json:"updated_at"`
}
type GroupDeleteRequest struct {
    Id int64 `json:"id"`
}
