package mq

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

type TaskClient struct {
	conf   ClientConfig
	client *asynq.Client
}

func NewTaskClient(conf ClientConfig) *TaskClient {
	return &TaskClient{
		conf: conf,
		client: asynq.NewClient(asynq.RedisClientOpt{
			Addr:     conf.Addr,
			Password: conf.Password,
			DB:       conf.DB,
		}),
	}
}

func (t *TaskClient) SetTimeOut(timeout time.Duration) asynq.Option {
	return asynq.Timeout(timeout)
}

func (t *TaskClient) SetProcessAt(timeAt time.Time) asynq.Option {
	return asynq.ProcessAt(timeAt)
}

func (t *TaskClient) SetProcessIn(timeIn time.Duration) asynq.Option {
	return asynq.ProcessIn(timeIn)
}

func (t *TaskClient) getOpts() []asynq.Option {
	opts := make([]asynq.Option, 0)
	if t.conf.Queue != "" {
		opts = append(opts, asynq.Queue(t.conf.Queue))
	}
	if t.conf.Retry != 0 {
		opts = append(opts, asynq.MaxRetry(t.conf.Retry))
	}
	if t.conf.Group != "" {
		opts = append(opts, asynq.Group(t.conf.Group))
	}
	return opts
}

// Timeout   time.Duration //指定任务可以运行多长时间
// Deadline  time.Time     //任务的截止时间
// UniqueTTL time.Duration //在多长时间内保证当前任务唯一
// ProcessAt time.Time     //在什么时候处理任务
// ProcessIn time.Duration //等待多久后处理任务
// Retention time.Duration //任务成功处理后保留多长时间

func (t *TaskClient) Dispatch(typeName string, v interface{}, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	opts = append(opts, t.getOpts()...)
	task := asynq.NewTask(typeName, payload, opts...)
	return t.client.Enqueue(task)
}

func (t *TaskClient) Close() {
	_ = t.client.Close()
}
