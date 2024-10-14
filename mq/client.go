package mq

import (
	"encoding/json"
	"github.com/hibiken/asynq"

	"time"
)

type TaskClient struct {
	client *asynq.Client
}

type TaskOpt struct {
	retry   int           //最大重试次数
	timeout time.Duration //指定任务可以运行多长时间
	delay   int           //等待多久后运行 单位s
	until   time.Time     //到什么时候开始运行
	queue   string
}

type TaskClientOpt func(*TaskOpt)

func SetRetry(retry int) TaskClientOpt {
	return func(t *TaskOpt) {
		t.retry = retry
	}
}

func SetTimeOut(timeout time.Duration) TaskClientOpt {
	return func(t *TaskOpt) {
		t.timeout = timeout
	}
}

func SetDelay(delay int) TaskClientOpt {
	return func(t *TaskOpt) {
		t.delay = delay
	}
}

func SetUntil(until time.Time) TaskClientOpt {
	return func(t *TaskOpt) {
		t.until = until
	}
}

func SetQueue(name string) TaskClientOpt {
	return func(t *TaskOpt) {
		t.queue = name
	}
}

func NewTaskClient(addr, pwd string, db int) *TaskClient {
	return &TaskClient{
		client: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     addr,
			Password: pwd,
			DB:       db,
		}),
	}
}

func (t *TaskClient) Dispatch(name string, payload *[]byte, opts ...TaskClientOpt) (*asynq.TaskInfo, error) {
	opt := &TaskOpt{
		retry: 3,
	}
	for _, f := range opts {
		f(opt)
	}
	now := time.Now()
	if opt.until.Before(now) {
		opt.until = now.Add(time.Duration(opt.delay) * time.Second)
	}
	return t.client.Enqueue(asynq.NewTask(name, *payload),
		asynq.ProcessAt(opt.until),
		asynq.Queue(opt.queue),
		asynq.MaxRetry(opt.retry),
		asynq.Timeout(opt.timeout))
}

func (t *TaskClient) DispatchNow(name string, payload *[]byte, opts ...TaskClientOpt) (*asynq.TaskInfo, error) {
	opt := &TaskOpt{
		retry: 3,
	}
	for _, f := range opts {
		f(opt)
	}
	return t.client.Enqueue(asynq.NewTask(name, *payload),
		asynq.Queue(opt.queue),
		asynq.MaxRetry(opt.retry),
		asynq.Timeout(opt.timeout))
}

func (t *TaskClient) SendJson(name string, data any, opts ...TaskClientOpt) (*asynq.TaskInfo, error) {
	param, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return t.Dispatch(name, &param, opts...)
}

func (t *TaskClient) Close() {
	_ = t.client.Close()
}

// type ClientConfig struct {
// 	Addr      string        //redis服务器地址，默认localhost:6379
// 	Password  string        //密码
// 	DB        int           //db
// 	Retry     int           //最大重试次数
// 	Queue     string        //加入的队列
// 	TaskID    string        //任务ID
// 	Timeout   time.Duration //指定任务可以运行多长时间
// 	Deadline  time.Time     //任务的截止时间
// 	UniqueTTL time.Duration //在多长时间内保证当前任务唯一
// 	ProcessAt time.Time     //在什么时候处理任务
// 	ProcessIn time.Duration //等待多久后处理任务
// 	Retention time.Duration //任务成功处理后保留多长时间
// 	Group     string        //加入的任务组
// }
