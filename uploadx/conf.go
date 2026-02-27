package uploadx

import (
	"io"
	"time"
)

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

// Storage 对象存储接口（支持多种云服务商）
type Storage interface {
	// Upload 上传文件
	// path: 存储路径（如：uploads/2024/01/01/file.jpg）
	// file: 文件内容
	// contentType: 文件MIME类型
	Upload(path string, file io.Reader, contentType string) (string, error)

	// UploadWithOptions 带选项的上传
	UploadWithOptions(path string, file io.Reader, opts *UploadOptions) (string, error)

	// Delete 删除文件
	Delete(path string) error

	// GetURL 获取文件访问URL
	// path: 存储路径
	// expires: 过期时间（0表示永久，>0表示生成临时签名URL）
	GetURL(path string, expires time.Duration) (string, error)

	// Exists 检查文件是否存在
	Exists(path string) (bool, error)

	// GetInfo 获取文件信息
	GetInfo(path string) (*FileInfo, error)
}

// UploadOptions 上传选项
type UploadOptions struct {
	ContentType        string            // MIME类型
	CacheControl       string            // 缓存控制
	ContentDisposition string            // 内容处置
	Metadata           map[string]string // 自定义元数据
	ACL                string            // 访问控制（public-read, private等）
}

// FileInfo 文件信息
type FileInfo struct {
	Path         string    // 文件路径
	Size         int64     // 文件大小（字节）
	ContentType  string    // MIME类型
	LastModified time.Time // 最后修改时间
	ETag         string    // ETag
	URL          string    // 访问URL
}

// Config 存储配置
type Config struct {
	Type      string `mapstructure:"type"`      // 存储类型：local, aliyun
	Endpoint  string `mapstructure:"endpoint"`  // 端点地址
	Bucket    string `mapstructure:"bucket"`    // 存储桶名称
	AccessKey string `mapstructure:"accessKey"` // 访问密钥ID
	SecretKey string `mapstructure:"secretKey"` // 访问密钥Secret
	Region    string `mapstructure:"region"`    // 区域
	Domain    string `mapstructure:"domain"`    // 自定义域名（CDN）
	IsPrivate bool   `mapstructure:"isPrivate"` // 是否私有（影响URL生成）
	BasePath  string `mapstructure:"basePath"`  // 基础路径
}

// StorageType 存储类型常量
const (
	TypeLocal  = "local"
	TypeAliyun = "aliyun"
)
