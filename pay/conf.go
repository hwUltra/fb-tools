package pay

// WxConf 微信支付配置
type WxConf struct {
	AppMiniId string `json:"appMiniId,optional"` //小程序 APPID
	//AppId      string `json:"appId,optional"`      //公众号 APPID
	//APPId      string `json:"AppId,optional"`      //APP   APPID
	MchId      string `json:"mchId,optional"`      //微信商户id
	SerialNo   string `json:"serialNo,optional"`   //商户证书的证书序列号
	ApiV3Key   string `json:"apiV3Key,optional"`   //apiV3Key，商户平台获取
	PrivateKey string `json:"privateKey,optional"` //privateKey：私钥 apiclient_key.pem 读取后的内容
	AppSecret  string `json:"appSecret,optional"`
	NotifyUrl  string `json:"notifyUrl,optional"` //支付通知回调服务端地址
}

// AliConf 微信支付配置
type AliConf struct {
	AppId      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	IsProd     bool   `json:"is_prod"`
	NotifyUrl  string `json:"notifyUrl"`
}
