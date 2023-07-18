package entity

import (
	"time"

	"github.com/lib/pq"
)

type Role struct {
	Id        int64         `gorm:"column:id;primary_key" json:"id"`
	Name      string        `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Resources pq.Int64Array `gorm:"column:resources;type:bigint[];not null" json:"resources"`
	CreatedAt time.Time     `gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:timestamp with time zone;not null" json:"updated_at"`
}

func (r *Role) TableName() string {
	return "roles"
}
