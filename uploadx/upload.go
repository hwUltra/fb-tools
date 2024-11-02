package uploadx

import "mime/multipart"

type OSS interface {
	UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error)
	DeleteFile(key string) error
}

func NewOss(conf OSSConf) OSS {
	switch conf.Type {
	case LocalType:
		return NewLocalOSS(conf.LocalConf)
	case MinioType:
		return NewMinioOSS(conf.MinIoConf)
	case AliYunType:
		return NewAliOSS(conf.AliYunConf)
	case TencentType:
		return NewTencentOSS(conf.TencentConf)
	case QiNiuType:
		return NewQiNiuOSS(conf.QiNiuConf)
	case AwsType:
		return NewAwsOSS(conf.AwsConf)
	case HuaWeiType:
		return NewHwOSS(conf.HuaWeiConf)
	default:
		return NewMinioOSS(conf.MinIoConf)
	}
}
