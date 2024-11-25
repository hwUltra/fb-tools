package uploadx

type OssType int

const (
	LocalType OssType = iota
	MinioType
	AliYunType
	TencentType
	QiNiuType
	AwsType
	HuaWeiType
)

type OSSConf struct {
	Type        OssType     `json:"type"`
	LocalConf   LocalConf   `json:"localConf,omitempty,optional"`
	MinIoConf   MinIoConf   `json:"minioConf,omitempty,optional"`
	AliYunConf  AliYunConf  `json:"aliyunConf,omitempty,optional"`
	TencentConf TencentConf `json:"tencentConf,omitempty,optional"`
	QiNiuConf   QiNiuConf   `json:"qiNiuConf,omitempty,optional"`
	AwsConf     AwsConf     `json:"awsConf,omitempty,optional"`
	HuaWeiConf  HuaWeiConf  `json:"huaWeiConf,omitempty,optional"`
}

type LocalConf struct {
	Path      string `json:"path"`       // 本地文件访问路径
	StorePath string `json:"store-path"` // 本地文件存储路径
}

type MinIoConf struct {
	MinIOAccessKeyID     string `json:"minIOAccessKeyID"`     //admin
	MinIOAccessSecretKey string `json:"minIOAccessSecretKey"` //MinIOAccessSecretKey
	MinIOEndpoint        string `json:"minIOEndpoint"`        //localhost:9000
	MinIOBucketLocation  string `json:"minIOBucketLocation"`  //cn-north-1
	MinIOSSLBool         bool   `json:"minIOSSLBool"`
	MinIOBucket          string `json:"minIOBucket"` //mymusic
	MinIOBasePath        string `json:"minIoBasePath"`
}

type AliYunConf struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret"`
	BucketName      string `json:"bucket-name"`
	BucketUrl       string `json:"bucket-url"`
	BasePath        string `json:"base-path"`
}

type TencentConf struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	SecretID   string `json:"secret-id"`
	SecretKey  string `json:"secret-key"`
	BaseURL    string `json:"base-url"`
	PathPrefix string `json:"path-prefix"`
}

type QiNiuConf struct {
	Zone          string `json:"zone"`            // 存储区域
	Bucket        string `json:"bucket"`          // 空间名称
	ImgPath       string `json:"img-path"`        // CDN加速域名
	AccessKey     string `json:"access-key"`      // 秘钥AK
	SecretKey     string `json:"secret-key"`      // 秘钥SK
	UseHTTPS      bool   `json:"use-https"`       // 是否使用https
	UseCdnDomains bool   `json:"use-cdn-domains"` // 上传是否使用CDN上传加速
}

type AwsConf struct {
	Bucket           string `json:"bucket"`
	Region           string `json:"region"`
	Endpoint         string `json:"endpoint"`
	SecretID         string `json:"secret-id"`
	SecretKey        string `json:"secret-key"`
	BaseURL          string `json:"base-url"`
	PathPrefix       string `json:"path-prefix"`
	S3ForcePathStyle bool   `json:"s3-force-path-style"`
	DisableSSL       bool   `json:"disable-ssl"`
}

type HuaWeiConf struct {
	Path      string `json:"path" yaml:"path"`
	Bucket    string `json:"bucket" yaml:"bucket"`
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	AccessKey string `json:"access-key" yaml:"access-key"`
	SecretKey string `json:"secret-key" yaml:"secret-key"`
}

type CloudflareR2Conf struct {
	Bucket          string `json:"bucket"`
	BaseURL         string `json:"base-url"`
	Path            string `json:"path"`
	AccountID       string `json:"account-id"`
	AccessKeyID     string `json:"access-key-id"`
	SecretAccessKey string `json:"secret-access-key"`
}

type UploadInfo struct {
	Path string `json:"path"`
	Hash string `json:"hash,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
}
