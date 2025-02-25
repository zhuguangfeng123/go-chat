package user

import (
	"github.com/gin-gonic/gin"
	iJwt "github.com/zhuguangfeng123/go-chat/internal/handler/jwt"
	"github.com/zhuguangfeng123/go-chat/pkg/ginx/result"
)

// UserProfile 获取用户自己的信息
func (hdl *UserHandler) UserProfile(ctx *gin.Context, uc iJwt.UserClaims) {
	user, err := hdl.userSvc.GetUser(ctx, uc.Uid)
	if err != nil {
		result.Failed(ctx, -1, "系统错误")
		return
	}

	result.Success(ctx, user)
}
