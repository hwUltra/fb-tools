package upload

type MinIoConf struct {
	MinIOAccessKeyID     string `json:"minIOAccessKeyID"`     //admin
	MinIOAccessSecretKey string `json:"minIOAccessSecretKey"` //MinIOAccessSecretKey
	MinIOEndpoint        string `json:"minIOEndpoint"`        //localhost:9000
	MinIOBucketLocation  string `json:"minIOBucketLocation"`  //cn-north-1
	MinIOSSLBool         bool   `json:"minIOSSLBool"`
	MinIOBucket          string `json:"minIOBucket"` //mymusic
	MinIOBasePath        string `json:"minIoBasePath"`
}

//type UploadInfo struct {
//	Hash string `json:"hash"`
//	Name string `json:"name"`
//	Ext  string `json:"ext"`
//	Size int64  `json:"size"`
//	Path string `json:"path"`
//}
