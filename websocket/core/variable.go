package core

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

// ClientMoreParams  为客户端成功上线后设置更多的参数
// ws 客户端成功上线以后，可以通过客户端携带的唯一参数，在数据库查询更多的其他关键信息，设置在 *Client 结构体上
// 这样便于在后续获取在线客户端时快速获取其他关键信息，例如：进行消息广播时记录日志可能需要更多字段信息等
type ClientMoreParams struct {
	ClientId string `json:"clientId"` //全局唯一的client_id
	Uid      int64  `json:"uid"`      //用户id
	Role     string `json:"role"`     //角色 - 》 admin 可以全局广播，普通用户只能广播自己
}
