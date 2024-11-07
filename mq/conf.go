package mq

type ClientConfig struct {
	Addr     string `json:"addr"`                        //redis服务器地址，默认localhost:6379
	Password string `json:"password,omitempty,optional"` //密码
	DB       int    `json:"db,omitempty,optional"`       //db
	Retry    int    `json:"retry,omitempty,optional"`    //最大重试次数
	Queue    string `json:"queue,omitempty,optional"`    //加入的队列
	Group    string `json:"group,omitempty,optional"`    //加入的任务组
}
