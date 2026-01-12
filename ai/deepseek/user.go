package deepseek

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hwUltra/fb-tools/ai/deepseek/constant"
	"github.com/hwUltra/fb-tools/ai/deepseek/types"
)

// UserBalance gets user balance ,doc link: https://api-docs.deepseek.com/zh-cn/api/get-user-balance
func (d *DeepSeekApi) UserBalance() (*types.UserBalanceResponse, error) {
	resp, err := d.httpClient.Get(constant.UserBalance, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", d.ApiKey),
	})
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	userBalanceRes := &types.UserBalanceResponse{}
	err = json.Unmarshal(body, userBalanceRes)
	if err != nil {
		return nil, err
	}
	return userBalanceRes, nil
}
