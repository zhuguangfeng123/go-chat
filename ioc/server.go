package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	iJwt "github.com/zhuguangfeng123/go-chat/internal/handler/jwt"
	"github.com/zhuguangfeng123/go-chat/internal/handler/user"
	"github.com/zhuguangfeng123/go-chat/internal/middleware"
	"strings"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc,
	userHandler *user.UserHandler,
) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)

	userHandler.RegisterRouter(server)

	return server
}

func InitGinMiddleware(jwt iJwt.JwtHandler) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//配置跨域请求
		cors.New(cors.Config{
			//AllowOrigins: []string{"*"},
			//AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Content-Type", "Authorization"}, //允许携带哪些请求头
			ExposeHeaders:    []string{"x-jwt-token"},
			AllowCredentials: true, //是否允许携带cookie之类的东西
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				return strings.Contains(origin, "xxxxx.com")
			}, //允许哪些域名访问
			MaxAge: 12 * time.Hour, //多长时间内第二次访问
		}),
		middleware.NewLoginJwtMiddlewareBuilder(jwt).
			IgnorePaths("/user/signup").
			IgnorePaths("/user/pwd-login").
			IgnorePaths("/user/send-login-sms").
			IgnorePaths("/user/sms-login").
			Build(),
	}
}
