package wsCore

import (
	"encoding/json"
	"github.com/hwUltra/fb-tools/result"
	"github.com/pkg/errors"
)

const (
	WebsocketHandshakeSuccess      = `{"code":200,"msg":"ws连接成功"}`
	WebsocketHandshakeError        = `{"code":0,"msg":"发送消息格式不正确"}`
	WebsocketServerPingMsg         = "Server->Ping->Client"
	WebsocketHeartbeatFailMaxTimes = 4
	WebsocketWriteReadBufferSize   = 20480
	WebsocketMaxMessageSize        = 65535
	WebsocketPingPeriod            = 20
	WebsocketReadDeadline          = 100
	WebsocketWriteDeadline         = 35
)

type ClientMoreParams struct {
	ClientId string `json:"clientId"` //全局唯一的client_id
	Uid      int64  `json:"uid"`      //用户id
	Role     string `json:"role"`     //角色 - 》 admin 可以全局广播，普通用户只能广播自己
}

// MsgBean 发送
type MsgBean struct {
	Type string `json:"type,default=sendAll"`
	To   string `json:"to,default=0"`
	Data string `json:"data,optional"`
}

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

func ReturnBeanFail(msg string) []byte {
	msgByte, _ := json.Marshal(
		&ReturnBean{
			Code: 10001,
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

func GetMsgBean(data []byte) (*MsgBean, error) {
	msgBean := MsgBean{}
	err := json.Unmarshal(data, &msgBean)
	return &msgBean, err
}
