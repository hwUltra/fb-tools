package deepseek

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hwUltra/fb-tools/ai/deepseek/constant"
	"github.com/hwUltra/fb-tools/ai/deepseek/types"
)

// ListModels list models ,doc link : https://api.deepseek.com/models
func (d *DeepSeekApi) ListModels() (*types.ModelListResponse, error) {
	resp, err := d.httpClient.Get(constant.ListModels, map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %v", d.ApiKey),
	})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	models := &types.ModelListResponse{}
	err = json.Unmarshal(body, models)
	if err != nil {
		return nil, err
	}
	return models, nil
}
