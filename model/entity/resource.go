package entity

import (
	"time"
)

type Resource struct {
	Id        int64     `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Type      string    `gorm:"column:type;type:varchar(255);not null" json:"type"`
	Path      string    `gorm:"column:path;type:varchar(255);not null" json:"path"`
	Icon      string    `gorm:"column:icon;type:varchar(255);not null" json:"icon"`
	ParentId  int64     `gorm:"column:parent_id;type:bigint;not null" json:"parent_id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp with time zone;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp with time zone;not null" json:"updated_at"`
}

func (r *Resource) TableName() string {
	return "resources"
}
