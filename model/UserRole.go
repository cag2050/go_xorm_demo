package model

type UserRole struct {
	Id     int64 `json:"id"`
	UserId int64 `json:"userId"`
	RoleId int64 `json:"roleId"`
}

func (UserRole) TableName() string {
	return "exporter_user_role"
}
