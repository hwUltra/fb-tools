package test

import (
	"encoding/json"
	wsTools "github.com/hwUltra/fb-tools/websocket"
	"testing"
)

func Test_Ws_SendMsg(t *testing.T) {
	host := "192.168.3.88:7300"
	path := "ws"
	query := "uid=88"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTkxODc3MTUsImlhdCI6MTY5OTEwMTMxNSwiand0VXNlcklkIjoxfQ.HF3NikInPM_12lD4mMaIX8jTAgHUp8S75GtJ9umstAA"

	dataByte, _ := json.Marshal(map[string]interface{}{
		"type": 0,
		"to":   "all",
		"msg":  "success",
	})
	err := wsTools.SendMsg(host, path, query, token, dataByte)
	if err != nil {
		t.Errorf("wsTools SendMsg write: %v", err)
	}
}

//fmt.Println("r.URL.RawQuery: ", r.URL.RawQuery)
//fmt.Println("r.URL.Path: ", r.URL.Path)
//fmt.Println("r.URL.Host: ", r.URL.Host)
//fmt.Println("r.URL.Fragment: ", r.URL.Fragment)
//fmt.Println("r.URL.Opaque: ", r.URL.Opaque)
//fmt.Println("r.URL.RawPath: ", r.URL.RawPath)
//fmt.Println("r.URL.Scheme: ", r.URL.Scheme)
//fmt.Println()
//fmt.Println("r.URL.RequestURI(): ", r.URL.RequestURI())
//fmt.Println("r.URL.Hostname(): ", r.URL.Hostname())
//fmt.Println("r.URL.Port(): ", r.URL.Port())
//fmt.Println("r.URL.String(): ", r.URL.String())
//fmt.Println("r.URL.EscapedPath(): ", r.URL.EscapedPath())
//fmt.Println()
//fmt.Println("r.Host: ", r.Host)
//fmt.Println("r.Method: ", r.Method)
//fmt.Println("r.UserAgent(): ", r.UserAgent())
//fmt.Println("r.RequestURI: ", r.RequestURI)
//fmt.Println("r.RemoteAddr: ", r.RemoteAddr)
//fmt.Println("r.FormValue(start): ", r.FormValue("start"))
//fmt.Println("r.FormValue(end): ", r.FormValue("end"))
//fmt.Println("r.FormValue(m): ", r.FormValue("m"))
//fmt.Println("r.FormValue(ms): ", r.FormValue("ms"))
//fmt.Println()
//fmt.Println("r.URL.Query().Get(start): ", r.URL.Query().Get("start"))
//fmt.Println("r.URL.Query().Get(end):", r.URL.Query().Get("end"))
//fmt.Println("r.URL.Query().Get(m):", r.URL.Query().Get("m"))
//fmt.Println("r.r.URL.Query()", r.URL.Query())
//fmt.Println("r.r.URL.Query()[“”\"m\"]", r.URL.Query()["m"])
//fmt.Println("r.URL.Query().Get(ms):", r.URL.Query().Get("ms"))
//
//fmt.Println("r.Header.Get(Content-Type):", r.Header.Get("Content-Type"))
