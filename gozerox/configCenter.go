package gozerox

import (
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
)

type ConfigCenter struct {
	Conf CenterConf
	Sub  subscriber.Subscriber
}

func MustNewConfigCenter(EtcdHost []string, EtcdKey string) *ConfigCenter {

	return &ConfigCenter{
		Conf: CenterConf{
			EtcdHost: EtcdHost,
			EtcdKey:  EtcdKey,
		},
		Sub: subscriber.MustNewEtcdSubscriber(subscriber.EtcdConf{
			Hosts: EtcdHost,
			Key:   EtcdKey,
		}),
	}
}

func (c ConfigCenter) GetConf(t configurator.Configurator[any]) any {
	cc := configurator.MustNewConfigCenter[t](configurator.Config{
		Type: "yaml",
	}, c.Sub)

	// 获取配置
	v, err := cc.GetConfig()
	if err != nil {
		panic(err)
	}
	// 监听配置
	cc.AddListener(func() {
		v, err = cc.GetConfig()
		if err != nil {
			panic(err)
		}
		//这个地方要写 触发配置变化后 需要处理的操作
		println("config changed", v)
	})

	return v

}
