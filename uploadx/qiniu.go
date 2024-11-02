package uploadx

import "mime/multipart"

type QiNiuOSS struct {
	Conf QiNiuConf
}

func NewQiNiuOSS(conf QiNiuConf) *QiNiuOSS {
	return &QiNiuOSS{
		Conf: conf,
	}
}

func (*QiNiuOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error) {

	return nil, nil
}

func (*QiNiuOSS) DeleteFile(key string) error {
	return nil
}
