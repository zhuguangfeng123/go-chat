package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	dtov1 "github.com/zhuguangfeng123/go-chat/internal/dto/v1"
	userSvc "github.com/zhuguangfeng123/go-chat/internal/service/user"
	"github.com/zhuguangfeng123/go-chat/pkg/ginx/result"
)

const ()

// PwdLogin 密码登录
func (hdl *UserHandler) PwdLogin(ctx *gin.Context, req dtov1.UserPwdLoginReq) {
	user, err := hdl.userSvc.UserPwdLogin(ctx, domain.User{
		Phone:    req.Phone,
		Password: req.Password,
	})

	if err != nil {
		if errors.Is(err, userSvc.ErrInvalidPhoneOrPassword) {
			result.Failed(ctx, -1, "密码错误")
			return
		}
		result.Failed(ctx, -1, "系统错误")
		return
	}

	err = hdl.SetLoginToken(ctx, user.UserId)
	if err != nil {
		result.Failed(ctx, -1, "系统错误")
		return
	}
	result.Success(ctx, nil)
}

// SendSmsLoginCode 发送短信登录验证码
func (hdl *UserHandler) SendSmsLoginCode(ctx *gin.Context, req dtov1.SendLoginSmsReq) {
	err := hdl.codeSvc.SendCode(ctx, "login", req.Phone)
	if err != nil {
		fmt.Println(err)
		result.Failed(ctx, -1, "系统错误")
		return
	}
	result.Success(ctx, nil)
}

// SmsLogin 短信登录
func (hdl *UserHandler) SmsLogin(ctx *gin.Context, req dtov1.UserSmsLoginReq) {
	fmt.Println(req)
	ok, err := hdl.codeSvc.VerifyCode(ctx, "login", req.Code, req.Phone)
	if err != nil {
		result.Failed(ctx, -1, err.Error())
		return
	}
	if !ok {
		result.Failed(ctx, -1, "验证码错误")
		return
	}
	user, err := hdl.userSvc.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		result.Failed(ctx, -1, err.Error())
		return
	}
	err = hdl.SetLoginToken(ctx, user.UserId)
	if err != nil {
		result.Failed(ctx, -1, "系统错误")
		return
	}
	result.Success(ctx, nil)
}
