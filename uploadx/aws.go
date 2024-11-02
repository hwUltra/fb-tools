package uploadx

import "mime/multipart"

type AwsOSS struct {
	Conf AwsConf
}

func NewAwsOSS(conf AwsConf) *AwsOSS {
	return &AwsOSS{
		Conf: conf,
	}
}

func (*AwsOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error) {

	return nil, nil
}

func (*AwsOSS) DeleteFile(key string) error {
	return nil
}
