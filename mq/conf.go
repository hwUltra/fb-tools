package mq

type ClientConfig struct {
	Addr     string //redis服务器地址，默认localhost:6379
	Password string //密码
	DB       int    //db
	Retry    int    //最大重试次数
	Queue    string //加入的队列
	Group    string //加入的任务组
}
