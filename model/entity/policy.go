package entity

type Policy struct {
	Id         int64  `gorm:"column:id;primary_key" json:"id"`
	Ptype      string `gorm:"column:ptype;type:varchar(255);not null" json:"ptype"`
	RoleId     int64  `gorm:"column:role_id;type:bigint;not null" json:"role_id"`
	ResourceId int64  `gorm:"column:resource_id;type:bigint;not null" json:"resource_id"`
	Operate    int    `gorm:"column:operate;type:integer;not null" json:"operate"`
	V3         string `gorm:"column:v3;type:varchar(255);not null" json:"v3"`
	V4         string `gorm:"column:v4;type:varchar(255);not null" json:"v4"`
	V5         string `gorm:"column:v5;type:varchar(255);not null" json:"v5"`
}

func (p *Policy) TableName() string {
	return "policies"
}
