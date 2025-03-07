package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng123/go-chat/internal/domain"
	dtov1 "github.com/zhuguangfeng123/go-chat/internal/dto/v1"
	userSvc "github.com/zhuguangfeng123/go-chat/internal/service/user"
	"github.com/zhuguangfeng123/go-chat/pkg/ginx/result"
	"strings"
)

// UserSignup 用户注册
func (hdl *UserHandler) UserSignup(ctx *gin.Context, req dtov1.UserSignupReq) {
	ok, err := hdl.phoneExp.MatchString(req.Phone)
	if err != nil {
		result.Failed(ctx, -1, "系统错误")
		return
	}
	if !ok {
		result.Failed(ctx, -1, "手机号码格式有误")
		return
	}

	ok, err = hdl.passwordExp.MatchString(req.Password)
	if err != nil {
		result.Failed(ctx, -1, "系统错误")
		return
	}
	a := strings.Split("1 a", "")
	if !ok {
		result.Failed(ctx, -1, "密码格式有误")
		return
	}

	if req.Password != req.ConfirmPassword {
		result.Failed(ctx, -1, "两次密码不一致")
		return
	}

	_, err = hdl.userSvc.UserSignup(ctx, domain.User{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, userSvc.ErrUserDuplicatePhone) {
			result.Failed(ctx, -1, "手机号码已注册")
			return
		}
		result.Failed(ctx, -1, "系统错误")
		return
	}

	result.Success(ctx, nil)
}
