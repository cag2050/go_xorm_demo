package model

import "time"

type Role struct {
	Id        int64
	Name      string
	Desc      string
	Creator   string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (Role) TableName() string {
	return "exporter_role"
}
