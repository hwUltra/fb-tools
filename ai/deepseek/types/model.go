package types

type ModelListResponse struct {
	Object string  `json:"object"` // [list]
	Data   []Model `json:"data"`
}

type Model struct {
	ID      string `json:"id"`       // 模型的标识符
	Object  string `json:"object"`   // 对象的类型，其值为 model。
	OwnedBy string `json:"owned_by"` // 拥有该模型的组织。
}
