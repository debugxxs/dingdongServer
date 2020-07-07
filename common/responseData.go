package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	QueryDataSuccess	string = "数据查询成功"
	QueryDataFail		string	=	"数据查询失败"
	InsertDataSuccess	string =	"数据添加成功"
	InsertDataFail		string =	"数据添加失败"
	UserIsExits			string =	"用户已存在"
	UserIsNotExits		string =	"用户不存在"
	IdNotExits			string =	"Id不存在"
	DelDataSuccess		string =	"删除数据成功"
	DelDataFail			string =	"删除数据失败"
	RoleNameNotExits	string	=	"用户角色不存在"
	OrganizationNotExits string =	"组织架构不存在"
	UpDataSuccess		string =	"数据更新成功"
	UpDataFail			string = 	"数据更新失败"
	RoleNameExits		string = 	"用户角色已存在"
	ParamsParseFail		string	=	"参数解析失败"
	ModifyPassErr		string	=	"修改密码错误，请传入正在登陆的用户名来修改密码"
)

func ResponseSuccessData(msg string,data interface{},c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":msg,
		"data":data,
	})
}

func ResponseSuccessMsg(msg string,c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"msg":msg,
	})
}

func ResponseDataFail(msg string,c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":http.StatusBadRequest,
		"msg":msg,
	})
}

func ResponseFailErr(err error)string {
	return fmt.Sprintln("请求数据出错",err)
}