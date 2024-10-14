package test

import (
	"fmt"
	"github.com/hwUltra/fb-tools/mqttx"
	"github.com/hwUltra/fb-tools/utils/ghelp"
	"testing"
	"time"
)

func TestMqtt(t *testing.T) {
	mt, err := mqttx.Create(mqttx.Conf{
		Host:     "192.168.3.88",
		Port:     11883,
		Username: "user",
		Password: "123123",
		ClientId: "go_mqtt_client111",
	}, func(msg mqttx.MqtMsg) {
		fmt.Printf("defalut Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	}, func() {
		fmt.Println("on connection")
	}, func(err error) {
		fmt.Println("lost connection err", err)
	})
	if err != nil {
		t.Errorf("mqttx CreateMqttX: %v", err)
	}
	topic := "topic/test"
	mt.Subscribe(mqttx.MqtMsg{Topic: topic, Qos: 1}, func(msg mqttx.MqtMsg) {
		fmt.Printf("TestMQTTX Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	})
	//mt.Sub(mqttx.MqtMsg{Topic: topic, Qos: 1}, nil)

	for i := 0; i < 4; i++ {
		text := fmt.Sprintf("Message %d", i)
		mt.Publish(mqttx.MqtMsg{Topic: topic, Payload: []byte(text)})
		time.Sleep(time.Second)
	}

	mt.Disconnect(0)
}

func TestMqttQuick(t *testing.T) {
	mt, err := mqttx.QuickCreate(mqttx.Conf{
		Host:     "192.168.3.88",
		Port:     11883,
		Username: "user",
		Password: "123123",
		ClientId: "go_mqtt_client111",
	})
	if err != nil {
		t.Errorf("Quick CreateMqttX: %v", err)
	}
	topic := "topic/test"
	mt.Subscribe(mqttx.MqtMsg{Topic: topic, Qos: 1}, func(msg mqttx.MqtMsg) {
		fmt.Printf("Quick Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	})

	for i := 0; i < 4; i++ {
		text := fmt.Sprintf("Quick Message %d", i)
		mt.Publish(mqttx.MqtMsg{Topic: topic, Payload: []byte(text)})
		time.Sleep(time.Second)
	}

	mt.Disconnect(250)
}

func TestMqttPublish(t *testing.T) {
	clientId := "go_mqtt_client222"
	mt, err := mqttx.QuickCreate(mqttx.Conf{
		Host:     "192.168.3.88",
		Port:     11883,
		Username: "user",
		Password: "123123",
		ClientId: clientId,
	})
	if err != nil {
		t.Errorf("Quick CreateMqttX: %v", err)
	}
	topic := "topic/test"
	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("TestMqttPublish msg %d", i)
		mt.Publish(mqttx.MqtMsg{Topic: topic, Payload: []byte(text)})
		time.Sleep(time.Second)
	}
	mt.Disconnect(0)
}

func TestMqttSub(t *testing.T) {
	clientId := "go_mqtt_client_01"
	mt, err := mqttx.Create(mqttx.Conf{
		Host:     "192.168.3.88",
		Port:     11883,
		Username: "user",
		Password: "123123",
		ClientId: clientId,
	}, func(msg mqttx.MqtMsg) {
		fmt.Printf("defalut Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	}, func() {
		fmt.Println("on connection", clientId)
	}, func(err error) {
		fmt.Println("lost connection err", clientId, err)
	})
	if err != nil {
		t.Errorf("mqttx CreateMqttX: %v", err)
	}
	topic := "order/#"
	mt.Subscribe(mqttx.MqtMsg{Topic: topic, Qos: 1}, func(msg mqttx.MqtMsg) {
		fmt.Printf("Quick Sub: %s from topic: %s\n", msg.Payload, msg.Topic)
	})
	//times := 0
	//for {
	//	times++
	//	time.Sleep(3 * time.Second)
	//}
	select {}
}

func TestA(t *testing.T) {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"a", "d"}
	s3 := ghelp.RemoveElements(s1, s2)
	fmt.Println(s3)
}
