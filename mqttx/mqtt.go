package mqttx

import (
	"crypto/tls"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/hwUltra/fb-tools/utils"
)

type MqttX struct {
	Conf   Conf
	Client mqtt.Client
	topics []string
}

type MsgHandler func(MqtMsg)
type OnConnectHandler func()
type OnConnectionLostHandler func(error)

func Create(conf Conf,
	defaultPublishHandler MsgHandler,
	onConnectHandler OnConnectHandler,
	onConnectionLostHandler OnConnectionLostHandler,
) (*MqttX, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf.Host, conf.Port))
	opts.SetUsername(conf.Username)
	opts.SetPassword(conf.Password)
	opts.SetClientID(conf.ClientId)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		if defaultPublishHandler != nil {
			defaultPublishHandler(MqtMsg{
				Topic:     message.Topic(),
				Duplicate: message.Duplicate(),
				Qos:       message.Qos(),
				Retained:  message.Retained(),
				MessageID: message.MessageID(),
				Payload:   message.Payload(),
				Ack:       message.Ack,
			})
		}
	})
	opts.OnConnect = func(client mqtt.Client) {
		if onConnectHandler != nil {
			onConnectHandler()
		}
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		if onConnectionLostHandler != nil {
			onConnectionLostHandler(err)
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MqttX{
		Conf:   conf,
		Client: client,
		topics: make([]string, 0),
	}, nil
}

func QuickCreate(conf Conf) (*MqttX, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf.Host, conf.Port))
	opts.SetUsername(conf.Username)
	opts.SetPassword(conf.Password)
	opts.SetClientID(conf.ClientId)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &MqttX{
		Conf:   conf,
		Client: client,
	}, nil
}

func (mx *MqttX) Publish(msg MqtMsg) {
	//client mqtt.Client
	token := mx.Client.Publish(msg.Topic, msg.Qos, msg.Retained, msg.Payload)
	token.Wait()
}

func (mx *MqttX) Subscribe(msg MqtMsg, callback MsgHandler) {
	//这里如果不指定方法，就用的上面的SetDefaultPublishHandler设置的方法
	var token mqtt.Token
	if callback != nil {
		token = mx.Client.Subscribe(msg.Topic, msg.Qos, func(client mqtt.Client, message mqtt.Message) {
			callback(MqtMsg{
				Topic:     message.Topic(),
				Duplicate: message.Duplicate(),
				Qos:       message.Qos(),
				Retained:  message.Retained(),
				MessageID: message.MessageID(),
				Payload:   message.Payload(),
				Ack:       message.Ack,
			})
		})
	} else {
		token = mx.Client.Subscribe(msg.Topic, msg.Qos, nil)
	}
	mx.topics = append(mx.topics, msg.Topic)
	//log.Println(fmt.Sprintf("订阅主题[%s]成功", msg.Topic))
	token.Wait()
}

func (mx *MqttX) Unsubscribe(topics ...string) error {
	if tc := mx.Client.Unsubscribe(topics...); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	mx.topics = utils.RemoveElements(mx.topics, topics)
	return nil
}

func (mx *MqttX) Disconnect(quiesce uint) {
	mx.Client.Disconnect(quiesce)
}

// 新建证书，也可以不用
func newTLSConfig(certFile string, privateKey string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, privateKey)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		ClientAuth:         tls.NoClientCert, //不需要证书
		ClientCAs:          nil,              //不验证证书
		InsecureSkipVerify: true,             //接受服务器提供的任何证书和该证书中的任何主机名
		Certificates:       []tls.Certificate{cert},
	}, nil
}
