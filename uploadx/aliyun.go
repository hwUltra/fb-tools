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

// UploadFile todo
func (*AliOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error) {

	return nil, nil
}

// DeleteFile todo
func (*AliOSS) DeleteFile(key string) error {
	return nil
}
