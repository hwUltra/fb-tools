package deepseek

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hwUltra/fb-tools/ai/deepseek/constant"
	"github.com/hwUltra/fb-tools/ai/deepseek/types"
)

// ChatCompletions chat completions api not stream
func (d *DeepSeekApi) ChatCompletions(req types.ChatRequest) (*http.Response, error) {
	res, err := d.httpClient.Post(constant.ChatCompletions, map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %v", d.ApiKey),
	}, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ChatNoStream chat completions api not stream
func (d *DeepSeekApi) ChatNoStream(req types.ChatRequest) (*types.ChatResponse, error) {
	req.Stream = false
	res, err := d.ChatCompletions(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	chatRes := &types.ChatResponse{}
	err = json.Unmarshal(body, chatRes)
	if err != nil {
		return nil, err
	}
	return chatRes, nil
}

// ChatStream chat completions api stream
func (d *DeepSeekApi) ChatStream(req types.ChatRequest) (chan types.ChatResponse, error) {
	req.Stream = true
	res, err := d.ChatCompletions(req)
	if err != nil {
		return nil, err
	}
	chatChan := make(chan types.ChatResponse, 1)
	go func() {
		scanner := bufio.NewScanner(res.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data: ") {
				line = strings.TrimPrefix(line, "data: ")
				if line == "[DONE]" {
					close(chatChan)
					return
				}
				chatRes := types.ChatResponse{}
				if err := json.Unmarshal([]byte(line), &chatRes); err != nil {
					fmt.Println(err.Error())
					continue
				}
				chatChan <- chatRes
			}
		}
	}()
	return chatChan, nil
}
