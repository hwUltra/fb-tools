package test

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/hwUltra/fb-tools/mq"
	"testing"
	"time"
)

func Test_Enqueue(t *testing.T) {
	payload := map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"}
	err := mq.DfSend("localhost:16379", "orderTask", payload)
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
}

func Test_Enqueue_delay(t *testing.T) {

	payload := map[string]interface{}{"OrderNo": "9999", "OrderMsg": "success ok"}
	err := mq.DfSendAfter("localhost:16379", "orderTask", payload, 10*time.Second)
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
}

func Test_EnqueueOther(t *testing.T) {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:16379"})
	payload := map[string]interface{}{"OrderNo": "6666", "OrderMsg": "success ok"}
	b, _ := json.Marshal(payload)
	task := asynq.NewTask("orderTask", b)
	// 10秒超时，最多重试3次，20秒后过期
	res, err := client.Enqueue(
		task,
		asynq.MaxRetry(3),
		asynq.Timeout(10*time.Second),
		asynq.Deadline(time.Now().Add(20*time.Second)))
	if err != nil {
		t.Errorf("could not enqueue task: %v", err)
		t.FailNow()
	}
	fmt.Printf("Enqueued Result: %+v\n", res)
}
