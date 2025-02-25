package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	iJwt "github.com/zhuguangfeng123/go-chat/internal/handler/jwt"
	"net/http"
)

// LoginJwtMiddlewareBuilder jwt登录校验
type LoginJwtMiddlewareBuilder struct {
	paths []string
	iJwt.JwtHandler
}

func NewLoginJwtMiddlewareBuilder(jwt iJwt.JwtHandler) *LoginJwtMiddlewareBuilder {
	return &LoginJwtMiddlewareBuilder{
		JwtHandler: jwt,
	}
}

func (l *LoginJwtMiddlewareBuilder) IgnorePaths(path string) *LoginJwtMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJwtMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		tokenStr := l.ExtractToken(ctx)
		var uc iJwt.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return iJwt.JwtKey, nil
		})

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if token == nil || !token.Valid {

			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = l.CheckSession(ctx, uc.Ssid)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user", uc)
	}
}
