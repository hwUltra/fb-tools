package types

type BalanceInfo struct {
	Currency        string `json:"currency"`          // 货币，人民币或美元
	TotalBalance    string `json:"total_balance"`     // 总的可用余额，包括赠金和充值余额
	GrantedBalance  string `json:"granted_balance"`   // 未过期的赠金余额
	ToppedUpBalance string `json:"topped_up_balance"` // 充值余额
}

type UserBalanceResponse struct {
	IsAvailable  bool          `json:"is_available"` // 当前账户是否有余额可供 API 调用
	BalanceInfos []BalanceInfo `json:"balance_infos"`
}
