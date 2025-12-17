package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success = 0
	Error   = 1
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 成功的
func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}
func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}
func OKWithnil(c *gin.Context) {
	Result(Success, map[string]any{}, "修改成功！！", c)
}

// 失败的
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithData(data any, c *gin.Context) {
	Result(Error, data, "成功", c)
}
func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func FailWithCode(code int, c *gin.Context) {
	//ok表示是否存在，meg为map中code对应的值
	msg, ok := ErrorMap[code]
	if ok {
		Result(code, map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误！", c)
}
