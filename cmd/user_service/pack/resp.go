package pack

import (
	"errors"
	"fmt"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp1(err error, msg string, code int64) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	return baseResp(errno.ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	})
}

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode: int32(err.ErrCode), StatusMessage: err.ErrMsg,
		ServiceTime: time.Now().Unix(),
	}
}
func BuildBaseResp(err error, code int64) *user.BaseResp {
	if err == nil {
		fmt.Println("++++++++++++")
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	return baseResp(errno.ErrNo{
		ErrCode: code,
		ErrMsg:  err.Error(),
	})
}
