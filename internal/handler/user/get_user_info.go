package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuguangfeng123/go-chat/pkg/ginx/result"
)

// GetUserInfo 获取用户信息
func (hdl *UserHandler) GetUserInfo(ctx *gin.Context) {
	userId := ctx.Query("userId")

	user, err := hdl.userSvc.GetUser(ctx, userId)
	if err != nil {
		result.Failed(ctx, -1, "系统错误")
		return
	}

	result.Success(ctx, user)
}
