package db

import (
	"dindongwork/common"
	"dindongwork/models"
	"dindongwork/tools"
	"fmt"
)

type UserDb struct {
}
//QueryUserPass 查询用户密码
func (ud UserDb)QueryUserPass(userName string)models.UserPass{
	/*
	1.定义映射实体
	2.查询数据
	3.返回查询到的数据
	*/
	userPass := models.UserPass{}
	tools.DbEngine.Table("user").Join("INNER","password","user.passId = password.passId and user.userName = ?",userName).Get(&userPass)
	return userPass
}

//QueryUser 根据用户名返回用户id
func (ud UserDb)QueryUserId(userName string)int64{
	user := models.User{}
	tools.DbEngine.Table("user").Where("userName = ?",userName).Get(&user)
	return user.UserId
}

//QueryUserPassId 根据用户名查询passId
func (ud UserDb)QueryUserPassId(userName string)int64  {
	user := models.User{}
	tools.DbEngine.Table("user").Where("userName = ?",userName).Get(&user)
	return user.PassId
}

//QueryUserRole 根据用户名返回用户角色信息
func (ud UserDb)QueryUserRole(userName string)models.UserRole{
	var userRole = models.UserRole{}
	tools.DbEngine.Table("user").Join("INNER","role","role.roleId = user.roleId and user.userName = ? ",userName).Get(&userRole)
	return userRole
}


//SaveUserToken 保存token到数据库
func (ud UserDb) SaveToken(userName , token string){
	/*
	1.先根据用户名查询用户id
	2.再根据用户id保存到数据库
	*/
	userId := ud.QueryUserId(userName)
	userToken := models.User{Token: token}
	tools.DbEngine.Table("user").ID(userId).Update(&userToken)
}

//QueryUserToken 查询用户token
func (ud UserDb)QueryUserToken(userName ,token string )(string,models.User)  {
	user := models.User{}
	_, err := tools.DbEngine.Table("user").Where("userName = ?", userName).Get(&user)
	if err !=nil{
		return common.QueryDataFail,user
	}
	return common.QueryDataSuccess,user
}

//PutUserPassword
func (ud UserDb)ModifyPassword(userName,password string)(string,int64)  {
	/*
	1.根据用户名查询passID
	2.修改password表
	*/
	passId := ud.QueryUserPassId(userName)
	userPass := models.Password{PassWord: password}
	fmt.Println(passId)
	fmt.Println(userPass)
	res,err:=tools.DbEngine.Table("password").ID(passId).Update(&userPass)
	if err !=nil{
		errMsg:=common.ResponseFailErr(err)
		return errMsg,res
	}
	return common.UpDataSuccess,res

}

//GetAllOrganization 查询所有的组织架构
func (ud UserDb)GetAllOrganization()(string,bool,[]models.Organization)  {
	var organizations []models.Organization
	err :=tools.DbEngine.Table("organization").Distinct("organizationId","topLayer","twoLayer","threeLayer").Find(&organizations)
	if err !=nil{
		return common.ResponseFailErr(err),false,organizations
	}
	return common.QueryDataSuccess,true,organizations
}

//QueryAllUser 查询所有用户信息
func (ud UserDb)QueryUserOrganization(organizationId int64)(string,bool,[]models.User)  {
	var users []models.User
	err :=tools.DbEngine.Table("user").Where("organizationId = ?",organizationId).Find(&users)
	if err !=nil{
		return common.ResponseFailErr(err),false,users
	}
	return common.QueryDataSuccess,true,users
}