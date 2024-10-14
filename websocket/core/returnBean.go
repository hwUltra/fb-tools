package core

import (
	"encoding/json"
	"github.com/hwUltra/fb-tools/xerr"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// ReturnBean 返回
type ReturnBean struct {
	Code uint32         `json:"code"`
	Msg  string         `json:"msg,optional"`
	Data ReturnDataBean `json:"data,optional"`
}

type ReturnDataBean struct {
	Type     string `json:"type,default=sendAll"`
	FromType string `json:"fromType,default=0,optional"` //user,admin,group
	From     int64  `json:"from,default=0,optional"`     //谁发的
	Msg      string `json:"msg"`
}

// NewReturnBean 创建ReturnBean
func NewReturnBean(data ReturnDataBean, code uint32, msg string) *ReturnBean {
	return &ReturnBean{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func ReturnBeanSuccess(data ReturnDataBean) []byte {
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

	errCode := xerr.SERVER_COMMON_ERROR
	errMsg := "服务器开小差啦，稍后再来试一试"

	causeErr := errors.Cause(err)                // err类型
	if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
		//自定义CodeError
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	} else {
		if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := uint32(gStatus.Code())
			if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
				errCode = grpcCode
				errMsg = gStatus.Message()
			}
		}
	}

	msg, _ := json.Marshal(&ReturnBean{Code: errCode, Msg: errMsg})
	return msg
}
