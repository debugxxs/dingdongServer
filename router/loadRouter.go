package router

import (
	"dindongwork/controller"
	"dindongwork/middleware/cors"
	"dindongwork/middleware/jwt"
	"github.com/gin-gonic/gin"
)
var (
	myJwt jwt.AuthJwt
	user controller.UserController
)

func LoadRouter(engine *gin.Engine)  {
	allAuthMiddleware := myJwt.AuthMiddlewareFunc(myJwt.AllAuthMiddleware)
	engine.Use(cors.Cors())
	engine.NoRoute(allAuthMiddleware.MiddlewareFunc(),myJwt.NoRouteHandler)
	engine.POST("/login",allAuthMiddleware.LoginHandler)
	userApi := engine.Group("/user")
	{
		userApi.GET("/refresh_token",allAuthMiddleware.RefreshHandler)
	}
	userApi.Use(allAuthMiddleware.MiddlewareFunc())
	{
		userApi.PUT("/userPass",user.ModifyUserPass)
		userApi.GET("/indexUserList",user.GetUserLIst)
		//消息通知api
		userApi.Any("/msgNotice",user.MsgNotice)
	}
}
