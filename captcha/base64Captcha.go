package captchaTool

import (
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"image/color"
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
	RedisConf redis.RedisConf
}

type CaptchaTool struct {
	Conf   CaptchaConf
	Store  base64Captcha.Store
	Driver base64Captcha.Driver
}

func NewCaptchaTool(conf CaptchaConf) *CaptchaTool {
	var store base64Captcha.Store
	if conf.Store == RedisType {
		store = NewRedisStore(conf.RedisConf)
	} else {
		store = base64Captcha.DefaultMemStore
	}

	var driver base64Captcha.Driver
	switch conf.Type {
	case DigType:
		driver = digitConfig()
	case MathType:
		driver = mathConfig()
	case StringType:
		driver = stringConfig()
	case ChineseType:
		driver = chineseConfig()
	case AudioType:
		driver = audioConfig()
	}
	return &CaptchaTool{
		Conf:   conf,
		Store:  store,
		Driver: driver,
	}
}

func (t *CaptchaTool) Make() (id, b64s, answer string, err error) {
	c := base64Captcha.NewCaptcha(t.Driver, t.Store)
	return c.Generate()
}

func (t *CaptchaTool) VerifyCaptcha(id string, VerifyValue string, clear bool) bool {
	return t.Store.Verify(id, VerifyValue, clear)
}

// mathConfig 生成图形化算术验证码配置
func mathConfig() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          60,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return mathType
}

// digitConfig 生成图形化数字验证码配置
func digitConfig() *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   60,
		Width:    240,
		Length:   4,
		MaxSkew:  0.45,
		DotCount: 80,
	}
	return digitType
}

// stringConfig 生成图形化字符串验证码配置
func stringConfig() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          60,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          4,
		Source:          "123456789qwertyuiopasdfghjklzxcvb",
		BgColor: &color.RGBA{
			R: 10,
			G: 20,
			B: 50,
			A: 10,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	return stringType
}

// chineseConfig 生成图形化汉字验证码配置
func chineseConfig() *base64Captcha.DriverChinese {
	chineseType := &base64Captcha.DriverChinese{
		Height:          60,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Source:          "大家啊第三方阿斯顿发的而且我和公司颠覆三观啊啊的发请求而且嘎达",
		Length:          4,
		BgColor: &color.RGBA{
			R: 10,
			G: 20,
			B: 50,
			A: 10,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	return chineseType
}

// autoConfig 生成图形化数字音频验证码配置
func audioConfig() *base64Captcha.DriverAudio {
	chineseType := &base64Captcha.DriverAudio{
		Length:   4,
		Language: "zh",
	}
	return chineseType
}
