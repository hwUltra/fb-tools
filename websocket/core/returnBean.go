package core

import (
	"encoding/json"
	"github.com/hwUltra/fb-tools/result"
	"github.com/pkg/errors"
)

// ReturnBean 返回
type ReturnBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg,optional"`
	Data interface{} `json:"data,optional"`
}

// NewReturnBean 创建ReturnBean
func NewReturnBean(data interface{}, code uint32, msg string) *ReturnBean {
	return &ReturnBean{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func ReturnBeanSuccess(data interface{}) []byte {
	msgByte, _ := json.Marshal(
		&ReturnBean{
			Code: 200,
			Msg:  "ok",
			Data: data,
		})
	return msgByte
}

func ReturnBeanMsg(msg string) []byte {
	msgByte, _ := json.Marshal(
		&ReturnBean{
			Code: 200,
			Msg:  msg,
		})
	return msgByte
}

func ReturnBeanError(err error) []byte {

	errCode := uint32(10001)
	errMsg := "服务器开小差啦，稍后再来试一试"

	causeErr := errors.Cause(err) // err类型
	var e *result.CodeError
	if errors.As(causeErr, &e) { //自定义错误类型
		//自定义CodeError
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	}

	msg, _ := json.Marshal(&ReturnBean{Code: errCode, Msg: errMsg})
	return msg
}
