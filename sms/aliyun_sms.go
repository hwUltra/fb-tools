package sms

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

type AliSms struct {
	Conf   AliConf
	Client *dysmsapi20170525.Client
}

func NewAliSms(conf AliConf) *AliSms {
	config := &openapi.Config{
		AccessKeyId:     tea.String(conf.AccessKeyId),
		AccessKeySecret: tea.String(conf.AccessSecret),
	}
	config.Endpoint = tea.String(conf.RegionId)
	client := &dysmsapi20170525.Client{}
	client, _ = dysmsapi20170525.NewClient(config)
	return &AliSms{
		Conf:   conf,
		Client: client,
	}
}

func (s *AliSms) SendCode(template string, phone string, code string) (_err error) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers: tea.String(phone),
		SignName:     tea.String(code),
		TemplateCode: tea.String(template),
	}
	//request := dysmsapi.CreateSendSmsRequest()
	////request属性设置
	//request.Scheme = "https"
	//request.SignName = s.Conf.SignName
	//request.TemplateCode = template
	//request.PhoneNumbers = phone
	//par, err := json.Marshal(map[string]interface{}{
	//	"code": code,
	//})
	//if err != nil {
	//	return err
	//}
	////设置验证码
	//request.TemplateParam = string(par)

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = s.Client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
