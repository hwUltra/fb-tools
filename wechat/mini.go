package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/auth"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

type MiniTool struct {
	Mini *miniprogram.MiniProgram
}

func NewMini(conf WxMiniConf) *MiniTool {
	wc := wechat.NewWechat()
	return &MiniTool{
		Mini: wc.GetMiniProgram(wxMiniConfig(conf)),
	}

}

// 内部
func wxMiniConfig(conf WxMiniConf) *miniConfig.Config {
	memory := cache.NewMemory()
	return &miniConfig.Config{
		AppID:     conf.AppId,
		AppSecret: conf.Secret,
		Cache:     memory,
	}
}

// Code2Session  code 获取 openid，SessionKey
func (m *MiniTool) Code2Session(code string) (auth.ResCode2Session, error) {
	code2Session, err := m.Mini.GetAuth().Code2Session(code)
	return code2Session, err
}

// Decrypt  消息解密
func (m *MiniTool) Decrypt(sessionKey string, encryptedData string, iv string) (plainData *encryptor.PlainData, err error) {
	return m.Mini.GetEncryptor().Decrypt(sessionKey, encryptedData, iv)
}

// Send 发送订阅消息
func (m *MiniTool) Send(msg *subscribe.Message) error {
	return m.Mini.GetSubscribe().Send(msg)
}

//wechat.Send(&subscribe.Message{
//	ToUser:           "",
//	TemplateID:       "",
//	Page:             "",
//	Data:             nil,
//	MiniprogramState: "",
//	Lang:             "",
//})
