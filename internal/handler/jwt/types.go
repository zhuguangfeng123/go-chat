package jwt

import "github.com/gin-gonic/gin"

type JwtHandler interface {
	//清除token
	ClearToken(ctx *gin.Context) error
	ExtractToken(ctx *gin.Context) string
	SetLoginToken(ctx *gin.Context, uid string) error
	SetJwtToken(ctx *gin.Context, uid string, ssid string) error
	CheckSession(ctx *gin.Context, ssid string) error
}
