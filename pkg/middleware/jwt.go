package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"tiktok/cmd/api/global"
	"tiktok/pkg/utils"
)

type OnlineUser struct {
	Id            int64  `json:"id" `           // 用户id
	LoginTime     int64  `json:"loginTime"`     // 登录时间
	LoginLocation string `json:"loginLocation"` // 归属地
	Ip            string `json:"ip"`            // ip地址
	Token         string `json:"key"`           // token
}
type Response struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"user"`
}

const (
	ERROR   = 400
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
func FailWithMessage(message string, c echo.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

// JWTAuth 基于JWT的认证中间件
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var token string
		token = c.QueryParam("token")
		fmt.Println(c)
		if token == "" {
			global.LOG.Error("token错误!")
			FailWithMessage("token错误", c)
			return next(c)
		}
		// 获取用户信息
		online := new(OnlineUser)
		j := new(utils.JWT)
		claims, _ := j.ParseToken(token, global.Config)
		redis := utils.Redis(global.Config, global.LOG)
		res, err := redis.Get(context.Background(), strconv.FormatInt(claims.UserID, 10)).Result()
		err = json.Unmarshal([]byte(res), &online)
		if err != nil {
			global.LOG.Error("token反序列化失败")
			FailWithMessage("token反序列化失败", c)
			return next(c)
		}
		if err != nil {
			global.LOG.Error("登录已过期，请重新登录!")
			FailWithMessage("登录已过期，请重新登录", c)
			return next(c)
		}

		if claims.UserID != online.Id || claims.Issuer != "XDream" {
			global.LOG.Error("token错误!")
			FailWithMessage("token错误", c)
		}
		return next(c)
	}
}
