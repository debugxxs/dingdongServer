package models

type User struct {
	UserId         int64  `json:"userId" xorm:"pk autoincr userId"`
	UserName       string `json:"userName"xorm:"varchar(16) userName"`
	Phone          string `json:"phone"xorm:"varchar(12) phone"`
	Email          string `json:"eMail" xorm:"varchar(64) eMail"`
	Avatar         string `json:"avatar" xorm:"varchar(64) avatar"`
	Position       string `json:"position" xorm:"varchar(64) position"`
	LoginTime      string `json:"loginTime"xorm:"updated"`
	Token		 	string	`json:"token" xorm:"varchar(254) taken"`
	PassId         int64  `json:"passId" xorm:"default 1 passId"`
	RoleId         int64  `json:"roleId" xorm:"default 2 roleId"`
	OrganizationId int64  `json:"organizationId"xorm:"default 0 organizationId"`
	DeletedAt      int64  `xorm:"deleted"`
}
type Password struct {
	PassId   int64  `json:"passId" xorm:"pk autoincr passId"`
	PassWord string `json:"passWord"xorm:"varchar(64) passWord"`
}
type UserPass struct {
	User     `xorm:"extends"`
	Password `xorm:"extends"`
}

//UserAllData 查询所有用户返回的信息实体
type UserAllData struct {
	UserId           int64        `json:"userId"`
	UserName         string       `json:"userName"`
	Phone            string       `json:"phone"`
	Email            string       `json:"eMail"`
	Avatar           string       `json:"avatar"`
	Position         string       `json:"position"`
	RoleName         string       `json:"roleName"`
	OrganizationName Organization `json:"organizationName"`
	DeletedAt        int64        `json:"deletedAt"`
}


type UserListData struct {
	UserId int64	`json:"userId"`
	UserName string	`json:"userName"`
}