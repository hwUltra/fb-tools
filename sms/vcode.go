package sms

import (
	"bytes"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type VCode struct {
	Config VCodeConf
	Cache  cache.Cache
	AliSms *AliSms
}

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("VCodeCache")
	cacheErr      = errors.New("not found")
)

func NewVCode(config VCodeConf, cacheConf cache.CacheConf) *VCode {
	aliSms := NewAliSms(config.AliConf)
	cacheClient := cache.New(cacheConf, singleFlights, stats, cacheErr)
	return &VCode{config, cacheClient, aliSms}
}

// Send 发送
func (v *VCode) Send(template string, mobile string) error {
	//debug 状态不发送不校验
	if v.Config.Debug {
		return nil
	}
	//testUsers 不发送不校验
	for _, item := range v.Config.TestUsers {
		if mobile == item {
			return nil
		}
	}
	//查看是否发送
	key := v.getKey(template, mobile)
	res := ""
	if err := v.Cache.Get(key, &res); err == nil {
		return errors.New("验证码已发送，请勿重复请求")
	}
	//生成缓存
	code := RandCode(v.Config.Length)
	//发送短信
	aliSms := AliSms{}
	if err := aliSms.SendCode(template, mobile, code); err != nil {
		return err
	}
	if err := v.Cache.SetWithExpire(key, code, v.Config.Life*time.Second); err != nil {
		return err
	}
	return nil

}

// Check 验证
func (v *VCode) Check(template string, mobile string, code string) error {
	//debug 状态不发送不校验
	if v.Config.Debug {
		return nil
	}
	//testUsers 不发送不校验
	for _, item := range v.Config.TestUsers {
		if mobile == item {
			return nil
		}
	}
	// 魔法密码直接放过
	if v.Config.MagicCode == code {
		return nil
	}
	// 正常校验
	key := v.getKey(template, mobile)

	vCode := ""
	if err := v.Cache.Get(key, &vCode); err != nil {
		return errors.New("验证码已过期")
	}
	if vCode != code {
		return errors.New("验证码有误")
	}
	return v.Cache.Del(key)
}

func (v *VCode) getKey(template string, mobile string) string {
	return template + "_validate_code_" + mobile
}

// RandCode 生成随机数
func RandCode(length int) string {
	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(math.Pow10(length)))
	s := bytes.Buffer{}
	s.WriteString(strconv.Itoa(int(randNum)))
	return s.String()
}
