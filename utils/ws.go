package utils

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

func WsSendMsg(host string, path string, query string, protocol []string, data []byte) error {
	var addr = flag.String("addr", host, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: path, RawQuery: query}

	requestHeader := http.Header{
		"Sec-WebSocket-Protocol": protocol,
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), requestHeader)
	if err != nil {
		return err
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)
	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}
