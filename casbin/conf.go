package casbin

type Conf struct {
	ModelText string `json:"ModelText,optional,env=CASBIN_MODEL_TEXT"`
}
