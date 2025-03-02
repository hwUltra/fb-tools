package deepseek

import (
	"github.com/hwUltra/fb-tools/http_client"
	"time"
)

type DeepSeekApi struct {
	ApiKey     string
	httpClient *http_client.Client
}

type opFunc func(*DeepSeekApi)

func NewDeepSeekApi(opFuncs ...opFunc) *DeepSeekApi {
	d := &DeepSeekApi{
		httpClient: http_client.NewClient(10 * time.Second),
	}
	for _, op := range opFuncs {
		op(d)
	}

	return d
}

func WithApiKey(apiKey string) opFunc {
	return func(d *DeepSeekApi) {
		d.ApiKey = apiKey
	}
}

func WithHttpClient(httpClient *http_client.Client) opFunc {
	return func(d *DeepSeekApi) {
		d.httpClient = httpClient
	}
}
