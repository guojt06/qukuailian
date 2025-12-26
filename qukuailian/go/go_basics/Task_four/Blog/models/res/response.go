package res

import (
	"modulename/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = 1
)

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, ctx *gin.Context) {
	Result(Success, data, msg, ctx)
}
func OKWithData(data any, ctx *gin.Context) {
	Result(Success, data, "成功", ctx)
}
func OKWithList(list any, count int64, ctx *gin.Context) {
	OKWithData(ListResponse[any]{
		Count: count,
		List:  list,
	}, ctx)
}

func OKWithMessage(msg string, ctx *gin.Context) {
	Result(Success, map[string]any{}, msg, ctx)
}
func OKWithCode(ctx *gin.Context) {
	Result(Success, map[string]any{}, "成功", ctx)
}

func Fail(data any, msg string, ctx *gin.Context) {
	Result(Error, data, msg, ctx)
}
func FailWithMessage(msg string, ctx *gin.Context) {
	Result(Error, map[string]any{}, msg, ctx)
}
func FailWithError(err error, obj any, ctx *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, ctx)
}

//func FailWithCode(code ErrorCode, ctx *gin.Context) {
//	msg, ok := ErrorMap[code]
//	if ok {
//		Result(int(code), map[string]any{}, msg, ctx)
//		return
//	}
//	Result(Error, map[string]any{}, "未知错误", ctx)
//}
