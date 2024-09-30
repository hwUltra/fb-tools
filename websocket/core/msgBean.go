package core

import "encoding/json"

// MsgBean 发送
type MsgBean struct {
	Type string `json:"type,default=sendAll"`
	To   string `json:"to,default=0"`
	Data string `json:"data,optional"`
}

func GetMsgBean(data []byte) (*MsgBean, error) {
	msgBean := MsgBean{}
	err := json.Unmarshal(data, &msgBean)
	return &msgBean, err
}
