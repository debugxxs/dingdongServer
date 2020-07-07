package jwt

import (
	"dindongwork/common"
	"dindongwork/models"
	"dindongwork/service"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthJwt struct {
	service.UserService
}
type AuthRoleFunction func(data interface{}, c *gin.Context) bool

var (
	identityKey    string = "UserName"
	GlobalUserName string
)
//AuthMiddlewareFunc 用户jwt定义
func (aj AuthJwt)AuthMiddlewareFunc(AuthRoleFunc AuthRoleFunction)(authmiddleware *jwt.GinJWTMiddleware){
	authmiddleware,err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Minute * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v,ok := data.(*models.User);ok{
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(context *gin.Context) interface{} {
			claims := jwt.ExtractClaims(context)
			return &models.User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context)(interface{},error){
			userPass := models.UserPass{}
			if err := c.ShouldBindJSON(&userPass);err !=nil{
				common.ErrHandler("用户登录参数解析错误",err)
			}
			GlobalUserName = userPass.UserName
			if aj.CheckUserPass(userPass.UserName, userPass.PassWord){
				return &models.User{UserName: userPass.UserName},nil
			}
			return nil,jwt.ErrFailedAuthentication
		},
		Authorizator: AuthRoleFunc,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time){
			userData := aj.CheckUserDataFunc(GlobalUserName)
			//保存token到数据库方便修改密码使用
			aj.SaveUserToken(GlobalUserName,token)
			c.JSON(http.StatusOK,gin.H{
				"code":http.StatusOK,
				"token":token,
				"time":expire,
				"msg":"login success",
				"data":userData,
			})
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,

	})
	common.ErrHandler("初始化jwt失败",err)
	return authmiddleware
}

//AllAuthMiddleware 所用用户登录定义方法
func (aj AuthJwt) AllAuthMiddleware(data interface{}, c *gin.Context) bool {
	return true
}


//NoRouteHandler 没有路由登录返回的内容
func (aj AuthJwt) NoRouteHandler(c *gin.Context) {
	c.JSON(404, gin.H{"code": 404, "message": "访问的路径不存在"})
}

