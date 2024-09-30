package test

import (
	"fmt"
	"github.com/hwUltra/fb-tools/captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

func TestCaptcha(t *testing.T) {
	var ct = captchaTool.NewCaptchaTool(captchaTool.CaptchaConf{
		Type:      captchaTool.MathType,
		Store:     captchaTool.RedisType,
		RedisConf: redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"},
	})
	if id, b64s, answer, err := ct.Make(); err != nil {
		t.Errorf("TestCaptcha: %v", err)
		t.FailNow()
	} else {
		fmt.Printf(" id = %s \n b64s = %s \n answer = %s \n", id, b64s, answer)
	}

}

func TestVerify(t *testing.T) {
	var ct = captchaTool.NewCaptchaTool(captchaTool.CaptchaConf{
		Type:      captchaTool.MathType,
		Store:     captchaTool.RedisType,
		RedisConf: redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"},
	})
	b := ct.VerifyCaptcha("Td39eOouS9FA9s9x0D9B", "q597", true)
	fmt.Println("TestVerifyCommon = ", b)
}
