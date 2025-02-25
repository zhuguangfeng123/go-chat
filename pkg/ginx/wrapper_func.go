package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func WrapBody[Req any](bizFn func(ctx *gin.Context, req Req)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": -1,
				"message":    "请求参数有误",
			})
			return
		}
		bizFn(ctx, req)
	}
}

func WrapClaims[Claims jwt.Claims](bizFn func(ctx *gin.Context, uc Claims)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(Claims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println(uc)
		bizFn(ctx, uc)
	}
}

func WrapBodyAndClaims[Req any, Claims jwt.Claims](bizFn func(ctx *gin.Context, req Req, uc Claims)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusOK, "请求参数有误")
			return
		}

		val, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(Claims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		bizFn(ctx, req, uc)
	}
}
