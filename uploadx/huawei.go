package uploadx

import "mime/multipart"

type HwOSS struct {
	Conf HuaWeiConf
}

func NewHwOSS(conf HuaWeiConf) *HwOSS {
	return &HwOSS{
		Conf: conf,
	}
}

func (*HwOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error) {

	return nil, nil
}

func (*HwOSS) DeleteFile(key string) error {
	return nil
}
