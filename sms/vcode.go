package sms

import (
	"bytes"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type VCode struct {
	Config      VCodeConf
	RedisClient *redis.Redis
	AliSms      *AliSms
}

func NewVCode(config VCodeConf, redisClient *redis.Redis) *VCode {
	aliSms := NewAliSms(config.AliConf)
	return &VCode{config, redisClient, aliSms}
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

	key := v.getKey(template, mobile)
	isExists, _ := v.RedisClient.Exists(key)
	if isExists {
		return errors.New("验证码已发送，请勿重复请求")
	}
	//缓存
	code := RandCode(v.Config.Length)
	err := v.RedisClient.Setex(key, code, v.Config.Life)
	if err != nil {
		return err
	}
	//发送短信
	aliSms := AliSms{}
	return aliSms.SendCode(template, mobile, code)
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

	vCode, err := v.RedisClient.Get(key)
	if err != nil {
		return errors.New("验证码有误")
	}
	if vCode != code {
		return errors.New("验证码有误")
	}
	_, _ = v.RedisClient.Del(key)
	return nil
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
