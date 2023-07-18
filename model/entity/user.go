package entity

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	Id        int64         `gorm:"column:id;primary_key" json:"id"`
	Name      string        `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Phone     string        `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	WxId      string        `gorm:"column:wx_id;type:varchar(255);not null" json:"wx_id"`
	Password  string        `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Salt      string        `gorm:"column:salt;type:varchar(255);not null" json:"salt"`
	Type      int           `gorm:"column:type;type:integer;not null" json:"type"`
	Avator    string        `gorm:"column:avator;type:varchar(255);not null" json:"avator"`
	Roles     pq.Int64Array `gorm:"column:roles;type:bigint[];not null" json:"roles"`
	Groups    pq.Int64Array `gorm:"column:groups;type:bigint[];not null" json:"groups"`
	CreatedAt time.Time     `gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:timestamp with time zone;not null" json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
