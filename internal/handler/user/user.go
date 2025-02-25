package user

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	dtov1 "github.com/zhuguangfeng123/go-chat/internal/dto/v1"
	iJwt "github.com/zhuguangfeng123/go-chat/internal/handler/jwt"
	"github.com/zhuguangfeng123/go-chat/internal/service/code"
	"github.com/zhuguangfeng123/go-chat/internal/service/sms"
	"github.com/zhuguangfeng123/go-chat/internal/service/user"
	"github.com/zhuguangfeng123/go-chat/pkg/ginx"
)

type UserHandler struct {
	phoneExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	iJwt.JwtHandler
	userSvc user.UserService
	smsSvc  sms.Service
	codeSvc code.CodeService
}

func NewUserHandler(jwtHandler iJwt.JwtHandler, userSvc user.UserService, smsSvc sms.Service, codeSvc code.CodeService) *UserHandler {
	const (
		phoneRegexPattern    = `^1[3-9]\d{9}$`
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,32}$`
	)

	return &UserHandler{
		phoneExp:    regexp.MustCompile(phoneRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		JwtHandler:  jwtHandler,
		userSvc:     userSvc,
		smsSvc:      smsSvc,
		codeSvc:     codeSvc,
	}
}

func (hdl *UserHandler) RegisterRouter(router *gin.Engine) {
	userG := router.Group("/user")
	userG.POST("/pwd-login", ginx.WrapBody[dtov1.UserPwdLoginReq](hdl.PwdLogin))
	userG.POST("/send-login-sms", ginx.WrapBody[dtov1.SendLoginSmsReq](hdl.SendSmsLoginCode))
	userG.POST("/sms-login", ginx.WrapBody[dtov1.UserSmsLoginReq](hdl.SmsLogin))
	userG.POST("/signup", ginx.WrapBody[dtov1.UserSignupReq](hdl.UserSignup))

	userG.GET("/profile", ginx.WrapClaims[iJwt.UserClaims](hdl.UserProfile))
	userG.GET("/get/user/info", hdl.GetUserInfo)
}

func (hdl *UserHandler) GetUser(ctx *gin.Context) iJwt.UserClaims {
	u, _ := ctx.Get("user")
	return u.(iJwt.UserClaims)
}
