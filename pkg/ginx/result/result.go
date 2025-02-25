package result

import "github.com/gin-gonic/gin"

type Result struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(ctx *gin.Context, data any) {
	var r Result
	r.Code = 200
	r.Msg = "SUCCESS"
	r.Data = data

	ctx.JSON(200, r)
	return
}

func Failed(ctx *gin.Context, code int64, msg string) {
	var r Result
	r.Code = code
	r.Msg = msg

	ctx.JSON(-1, r)
	return
}
