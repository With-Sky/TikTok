package pack

import (
	"errors"
	"tiktok/kitex_gen/message"
	"tiktok/pkg/errno"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp1(err error, msg string, code int64) *message.BaseResp {
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

func baseResp(err errno.ErrNo) *message.BaseResp {
	return &message.BaseResp{
		StatusCode: int32(err.ErrCode), StatusMessage: err.ErrMsg,
		ServiceTime: time.Now().Unix(),
	}
}
func BuildBaseResp(err error, code int64) *message.BaseResp {
	if err == nil {
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
