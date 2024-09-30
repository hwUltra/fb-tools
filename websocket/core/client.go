package core

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub                *Hub            // 负责处理客户端注册、注销、在线管理
	Conn               *websocket.Conn // 一个ws连接
	Send               chan []byte     // 一个ws连接存储自己的消息管道
	PingPeriod         time.Duration
	ReadDeadline       time.Duration
	WriteDeadline      time.Duration
	HeartbeatFailTimes int
	State              uint8 // ws状态，1=ok；0=出错、掉线等
	sync.RWMutex
	ClientMoreParams // 这里追加一个结构体，方便开发者在成功上线后，可以自定义追加更多字段信息
}

// OnOpen 处理握手+协议升级
func (c *Client) OnOpen(hub *Hub, w http.ResponseWriter, r *http.Request) (*Client, bool) {
	// 1.升级连接,从http--->websocket
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("websocket onopen 发生错误", err)
		}
	}()
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  20480,
		WriteBufferSize: 20480,
		Subprotocols:    []string{r.Header.Get("Sec-WebSocket-Protocol")},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 2.将http协议升级到websocket协议.初始化一个有效的websocket长连接客户端
	if wsConn, err := upGrader.Upgrade(w, r, nil); err != nil {
		fmt.Println("websocket Upgrade 协议升级, 发生错误" + err.Error())
		return nil, false
	} else {
		c.Hub = hub
		c.Conn = wsConn
		c.Send = make(chan []byte, 20480)
		c.PingPeriod = time.Second * 20
		c.ReadDeadline = time.Second * 100
		c.WriteDeadline = time.Second * 35

		if err := c.SendMessage(websocket.TextMessage, []byte(WebsocketHandshakeSuccess)); err != nil {
			fmt.Println("websocket  Write Msg(send msg) 失败", err)
		}
		c.Conn.SetReadLimit(65535) // 设置最大读取长度
		c.Hub.Register <- c
		c.State = 1
		return c, true
	}

}

// ReadPump 主要功能主要是实时接收消息
func (c *Client) ReadPump(callbackOnMessage func(messageType int, receivedData []byte), callbackOnError func(err error), callbackOnClose func()) {
	// 回调 onclose 事件
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("websocket ReadPump(实时读取消息)协程出错", err)
		}
		callbackOnClose()
	}()

	// OnMessage事件
	for {
		if c.State == 1 {
			mt, bReceivedData, err := c.Conn.ReadMessage()
			if err == nil {
				//fmt.Println("ReadPump", mt, string(bReceivedData))
				callbackOnMessage(mt, bReceivedData)

			} else {
				// OnError事件读（消息出错)
				callbackOnError(err)
				break
			}
		} else {
			// OnError事件(状态不可用，一般是程序事先检测到双方无法进行通信，进行的回调)
			callbackOnError(errors.New("websocket  state 状态已经不可用(掉线、卡死等愿意，造成双方无法进行数据交互)"))
			break
		}

	}
}

// SendMessage 发送消息，请统一调用本函数进行发送
// 消息发送时增加互斥锁，加强并发情况下程序稳定性
// 提醒：开发者发送消息时，不要调用 c.Conn.WriteMessage(messageType, []byte(message)) 直接发送消息
func (c *Client) SendMessage(messageType int, message []byte) error {
	c.Lock()
	defer func() {
		c.Unlock()
	}()
	if err := c.Conn.SetWriteDeadline(time.Now().Add(c.WriteDeadline)); err != nil {
		fmt.Println("websocket  设置消息写入截止时间出错", err)
		return err
	}
	if err := c.Conn.WriteMessage(messageType, message); err != nil {
		return err
	}
	return nil
}

// Heartbeat  按照websocket标准协议实现隐式心跳,Server端向Client远端发送ping格式数据包,浏览器收到ping标准格式，自动将消息原路返回给服务器
func (c *Client) Heartbeat() {
	//  1. 设置一个时钟，周期性的向client远端发送心跳数据包
	ticker := time.NewTicker(c.PingPeriod)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("websocket BeatHeart心跳协程出错", err)
		}
		ticker.Stop() // 停止该client的心跳检测
	}()
	//2.浏览器收到服务器的ping格式消息，会自动响应pong消息，将服务器消息原路返回过来
	if c.ReadDeadline == 0 {
		_ = c.Conn.SetReadDeadline(time.Time{})
	} else {
		_ = c.Conn.SetReadDeadline(time.Now().Add(c.ReadDeadline))
	}
	c.Conn.SetPongHandler(func(receivedPong string) error {
		if c.ReadDeadline > time.Nanosecond {
			_ = c.Conn.SetReadDeadline(time.Now().Add(c.ReadDeadline))
		} else {
			_ = c.Conn.SetReadDeadline(time.Time{})
		}
		//fmt.Println("浏览器收到ping标准格式，自动将消息原路返回给服务器：", received_pong)  // 接受到的消息叫做pong，实际上就是服务器发送出去的ping数据包
		return nil
	})
	//3.自动心跳数据
	for {
		select {
		case <-ticker.C:
			if c.State == 1 {
				if err := c.SendMessage(websocket.PingMessage, []byte(WebsocketServerPingMsg)); err != nil {
					c.HeartbeatFailTimes++
					if c.HeartbeatFailTimes > HeartbeatFailMaxTimes {
						c.State = 0
						return
					}
				} else {
					if c.HeartbeatFailTimes > 0 {
						c.HeartbeatFailTimes--
					}
				}
			} else {
				return
			}

		}
	}
}

func (c *Client) Close() {
	c.State = 0
	c.Hub.UnRegister <- c
	_ = c.Conn.Close()
}
