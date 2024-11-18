package sms

type VCodeTypeEnum int

const (
	AliYun = iota
	QCloud
)

type VCodeConf struct {
	AliConf   AliConf
	Type      VCodeTypeEnum
	Debug     bool              `json:"debug,default=false"`
	Length    int               `json:"length,default=6"`
	Life      int64             `json:"life,default=300"`
	MagicCode string            `json:"magicCode,omitempty,optional"`
	TestUsers []string          `json:"testUsers,omitempty,optional"`
	Template  map[string]string `json:"template,omitempty,optional"`
}

type AliConf struct {
	RegionId     string
	AccessKeyId  string
	AccessSecret string
	SignName     string
}
