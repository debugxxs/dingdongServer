package models

type Role struct {
	RoleId   int64  `json:"roleId" xorm:"pk autoincr roleId"`
	RoleName string `json:"roleName" xorm:"varchar(64) roleName"`
}

type UserRole struct {
	User `xorm:"extends"`
	Role `xorm:"extends"`
}
