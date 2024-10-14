package mq

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"log"
)

type Jobs map[string]func(context.Context, *asynq.Task) error

type Taskser struct {
	redisAddr   string
	password    string
	db          int
	queues      map[string]int
	concurrency int
	jobMap      *Jobs

	server *asynq.Server
}

type Taskopt func(*Taskser)

func SetTaskAddr(addr string) Taskopt {
	return func(t *Taskser) {
		t.redisAddr = addr
	}
}

func SetTaskPwd(pwd string) Taskopt {
	return func(t *Taskser) {
		t.password = pwd
	}
}

func SetTaskDB(db int) Taskopt {
	return func(t *Taskser) {
		t.db = db
	}
}

func SetTaskJobs(jobs *Jobs) Taskopt {
	return func(t *Taskser) {
		t.jobMap = jobs
	}
}

func SetConcurrency(num int) Taskopt {
	return func(t *Taskser) {
		t.concurrency = num
	}
}

func SetTaskQueues(name string, value int) Taskopt {
	return func(t *Taskser) {
		t.queues[name] = value
	}
}

func NewTaskService(opts ...Taskopt) *Taskser {
	t := &Taskser{
		redisAddr:   "localhost:6379",
		concurrency: 10,
		queues:      map[string]int{},
	}
	for _, f := range opts {
		f(t)
	}
	return t
}

func NewTaskSerByConfig(conf *redis.RedisConf, opts ...Taskopt) *Taskser {
	t := &Taskser{
		redisAddr:   conf.Host,
		password:    conf.Pass,
		concurrency: 10,
	}
	for _, f := range opts {
		f(t)
	}
	return t
}

func (t *Taskser) Start() {
	r := asynq.RedisClientOpt{
		Addr:     t.redisAddr,
		Password: t.password,
		DB:       t.db,
	}
	t.server = asynq.NewServer(r, asynq.Config{
		Concurrency: t.concurrency,
		Queues:      t.queues,
	})

	mux := asynq.NewServeMux()
	for name, job := range *t.jobMap {
		mux.HandleFunc(name, job)
	}

	if err := t.server.Run(mux); err != nil {
		log.Fatal(err)
	}

}

func (t *Taskser) Stop() {
	if t.server != nil {
		t.server.Stop()
	}
}

// type ServerConfig struct {
// 	Addr                     string                 //redis服务器地址，默认localhost:6379
// 	Password                 string                 //密码
// 	DB                       int                    //db
// 	Concurrency              int                    //最大并发数
// 	BaseContext              func() context.Context //基本上下文
// 	RetryDelayFunc           asynq.RetryDelayFunc   //计算失败任务的重试延迟时间
// 	IsFailure                func(error) bool       //确定处理程序返回的错误是否为故障
// 	Queues                   map[string]int         //队列
// 	StrictPriority           bool                   //是否严格按优先级来处理队列中的任务
// 	ErrorHandler             asynq.ErrorHandler     //当任务处理函数返回非nil错误时会调用此函数
// 	Logger                   asynq.Logger           //日志
// 	LogLevel                 asynq.LogLevel         //日志级别
// 	ShutdownTimeout          time.Duration          //服务停止时等待多少时间
// 	HealthCheckFunc          func(error)            //当ping redis服务器发生错误时调用
// 	HealthCheckInterval      time.Duration          //检查redis服务器的间隔时间，默认15秒
// 	DelayedTaskCheckInterval time.Duration          //指定对“计划”和“重试”任务运行的检查之间的间隔 ，默认5秒
// 	GroupGracePeriod         time.Duration          //指定服务器等待任务组中的任务时，等待多少时间，默认1分钟
// 	GroupMaxDelay            time.Duration          //等待任务组中的任务最大时间
// 	GroupMaxSize             int                    //一个任务组中的最大任务数
// 	GroupAggregator          asynq.GroupAggregator  //指定怎样将多个任务组合成一个任务
// }
