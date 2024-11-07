package captchax

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type CaptchaType int
type StoreType int

/*
*
dight 数字验证码
audio 语音验证码
string 字符验证码
math 数学验证码(加减乘除)
chinese中文验证码-有bug
*/
const (
	DigType     CaptchaType = iota //数字验证码
	MathType                       //数学验证码(加减乘除)
	StringType                     //字符验证码
	ChineseType                    //中文验证码
	AudioType                      //语音验证码
)

const (
	MemType StoreType = iota
	RedisType
)

type CaptchaConf struct {
	Type      CaptchaType
	Store     StoreType
	CacheConf cache.CacheConf
}

type CaptchaTool struct {
	Conf   CaptchaConf
	Store  base64Captcha.Store
	Driver base64Captcha.Driver
}
