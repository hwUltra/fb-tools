package uploadx

import "mime/multipart"

type AliOSS struct {
	Conf AliYunConf
}

func NewAliOSS(conf AliYunConf) *AliOSS {
	return &AliOSS{
		Conf: conf,
	}
}

func (*AliOSS) UploadFile(file *multipart.FileHeader) (*UploadInfo, error) {

	return nil, nil
}

func (*AliOSS) DeleteFile(key string) error {
	return nil
}
