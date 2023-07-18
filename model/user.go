package model

import "github.com/lib/pq"

type UserCreateRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Phone     *string       `json:"phone"`
    WxId      *string       `json:"wx_id"`
    Password  string        `binding:"required" json:"password"`
    Type      int           `binding:"required" json:"type"`
    Avator    *string       `json:"avator"`
    Roles     pq.Int64Array `json:"roles"`
    Groups    pq.Int64Array `binding:"required" json:"groups"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
}
type UserUpdateRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Phone     *string       `json:"phone"`
    WxId      *string       `json:"wx_id"`
    Password  string        `binding:"required" json:"password"`
    Type      int           `binding:"required" json:"type"`
    Avator    *string       `json:"avator"`
    Roles     pq.Int64Array `json:"roles"`
    Groups    pq.Int64Array `binding:"required" json:"groups"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
}
type UserListRequest struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Phone     *string       `json:"phone"`
    WxId      *string       `json:"wx_id"`
    Password  string        `binding:"required" json:"password"`
    Type      int           `binding:"required" json:"type"`
    Avator    *string       `json:"avator"`
    Roles     pq.Int64Array `json:"roles"`
    Groups    pq.Int64Array `binding:"required" json:"groups"`
    CreatedAt *int64        `json:"created_at"`
    UpdatedAt *int64        `json:"updated_at"`
    Index     int           `json:"index"`
    Size      int           `json:"size"`
}
type UserListResponse struct {
    Total int         `json:"total"`
    List  []*UserInfo `json:"list"`
}
type UserInfoRequest struct {
    Id int64 `json:"id"`
}
type UserInfo struct {
    Id        int64         `json:"id"`
    Name      string        `binding:"required" json:"name"`
    Phone     string        `json:"phone"`
    WxId      string        `json:"wx_id"`
    Password  string        `binding:"required" json:"password"`
    Salt      string        `json:"salt"`
    Type      int           `binding:"required" json:"type"`
    Avator    string        `json:"avator"`
    Roles     pq.Int64Array `json:"roles"`
    Groups    pq.Int64Array `binding:"required" json:"groups"`
    CreatedAt int64         `json:"created_at"`
    UpdatedAt int64         `json:"updated_at"`
}
type UserDeleteRequest struct {
    Id int64 `json:"id"`
}
