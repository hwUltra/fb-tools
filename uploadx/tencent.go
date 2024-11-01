package uploadx

import "mime/multipart"

type TencentOSS struct {
	Conf TencentConf
}

func NewTencentOSS(conf TencentConf) *TencentOSS {
	return &TencentOSS{
		Conf: conf,
	}
}

func (*TencentOSS) UploadFile(file *multipart.FileHeader) (*UploadInfo, error) {

	return nil, nil
}

func (*TencentOSS) DeleteFile(key string) error {
	return nil
}
