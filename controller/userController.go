package controller

import (
	"dindongwork/common"
	"dindongwork/models"
	"dindongwork/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
)

type UserController struct {
	service.UserService
}
var (
	upGrader = &websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
//ModifyUserPass 修改用户密码
func (uc UserController)ModifyUserPass( c *gin.Context)  {
	/*
	1.获取用户登录token，验证是否正在登陆的用户
	2.获取要修改的密码json
	3.将json传入service进行处理
	*/
	userHandler := c.Request.Header.Get("Authorization")
	handlerList := strings.Split(userHandler," ")
	tokenStr := handlerList[1]

	var userPass models.UserPass
	if err := c.ShouldBindJSON(&userPass);err!=nil{
		errMsg:=common.ResponseFailErr(err)
		common.ResponseDataFail(errMsg,c)
	}
	msg,result:=uc.CheckUser(userPass.User.UserName,tokenStr,userPass.Password.PassWord)
	if result{
		common.ResponseDataFail(msg,c)
	}else {
		common.ResponseSuccessMsg(msg,c)
	}
}


//GetUserLIst 获取首页好友列表
func (uc UserController)GetUserLIst(c *gin.Context){
	msg,res,userOrganization:=uc.GetListData()
	if res{
		common.ResponseSuccessData(msg,userOrganization,c)
	}else {
		common.ResponseDataFail(msg,c)
	}

}

func (uc UserController)MsgNotice(c *gin.Context){
	/*
	1.升级协议
	*/
	var (
		upGrader = &websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
	)
	ws,err := upGrader.Upgrade(c.Writer,c.Request,nil)
	ws.Subprotocol()
	if err !=nil{
		errMsg:=common.ResponseFailErr(err)
		common.ResponseDataFail(errMsg,c)
	}
	client:=ws.RemoteAddr()
	fmt.Println(client.String())
	_, msg, err := ws.ReadMessage()
	if err !=nil{
		logger.Fatalln("出错了",err)
	}
	fmt.Println("msg",msg)
	ws.WriteMessage(websocket.TextMessage,[]byte("hello"))
}