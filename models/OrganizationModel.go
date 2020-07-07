package models

type Organization struct {
	OrganizationId int64  `json:"organizationId" xorm:"pk autoincr organizationId"`
	TOPLayer       string `json:"topLayer" xorm:"varchar(64) topLayer"`
	TwoLayer       string `json:"twoLayer" xorm:"varchar(64) twoLayer"`
	ThreeLayer     string `json:"threeLayer"xorm:"varchar(64) threeLayer"`
}
type UserOrganization struct {
	User         `xorm:"extends"`
	Organization `xorm:"extends"`
}
