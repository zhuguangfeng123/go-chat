package handler

import "github.com/gin-gonic/gin"

type Router interface {
	RegisterRouter(router *gin.Engine)
}
