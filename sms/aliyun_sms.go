package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliSms struct {
	Conf   AliConf
	Client *dysmsapi.Client
}

func NewAliSms(conf AliConf) *AliSms {
	client, _ := dysmsapi.NewClientWithAccessKey(conf.RegionId, conf.AccessKeyId, conf.AccessSecret)
	return &AliSms{
		Conf:   conf,
		Client: client,
	}
}

func (s *AliSms) SendCode(template string, phone string, code string) error {
	request := dysmsapi.CreateSendSmsRequest()
	//request属性设置
	request.Scheme = "https"
	request.SignName = s.Conf.SignName
	request.TemplateCode = template
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	if err != nil {
		return err
	}
	//设置验证码
	request.TemplateParam = string(par)
	response, err := s.Client.SendSms(request)
	if err != nil {
		return err
	}
	fmt.Println("response", response)
	//检查返回结果，并判断发生状态
	//{"RequestId":"D07C8355-1EC9-57FB-9A11-19861E18ECFB","Message":"OK","BizId":"215801540511695335^0","Code":"OK"}
	if response.Code != "OK" {
		return errors.New(response.Message)
	}
	return nil
}
