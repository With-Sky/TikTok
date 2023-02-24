package controllor

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"user"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c echo.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c echo.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c echo.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c echo.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c echo.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c echo.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c echo.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c echo.Context) {
	Result(ERROR, data, message, c)
}
