package entity

import (
	"time"

	"github.com/lib/pq"
)

type Group struct {
	Id        int64         `gorm:"column:id;primary_key" json:"id"`
	Name      string        `gorm:"column:name;type:varchar(255);not null" json:"name"`
	ParentId  int64         `gorm:"column:parent_id;type:bigint;not null" json:"parent_id"`
	Roles     pq.Int64Array `gorm:"column:roles;type:bigint[];not null" json:"roles"`
	CreatedAt time.Time     `gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:timestamp with time zone;not null" json:"updated_at"`
}

func (g *Group) TableName() string {
	return "groups"
}
