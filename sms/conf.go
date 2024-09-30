package sms

type VCodeTypeEnum int

const (
	AliYun = iota
	QCloud
)

type VCodeConf struct {
	AliConf   AliConf
	Type      VCodeTypeEnum
	Debug     bool
	Length    int
	Life      int
	MagicCode string
	TestUsers []string
	Template  Template
}

type AliConf struct {
	RegionId     string
	AccessKeyId  string
	AccessSecret string
	SignName     string
}

type Template struct {
	Reg string
}
