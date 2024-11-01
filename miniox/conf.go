package miniox

type MinioConf struct {
	MinIOAccessKeyID     string `json:"minIOAccessKeyID"`     //admin
	MinIOAccessSecretKey string `json:"minIOAccessSecretKey"` //MinIOAccessSecretKey
	MinIOEndpoint        string `json:"minIOEndpoint"`        //localhost:9000
	MinIOBucketLocation  string `json:"minIOBucketLocation"`  //cn-north-1
	MinIOSSLBool         bool   `json:"minIOSSLBool"`
	MinIOBucket          string `json:"minIOBucket"` //mymusic
	MinIOFile            string `json:"MinIOFile"`
	MinIOBasePath        string `json:"minIoBasePath"`
	//MinIOBucketPolicy    string `json:"minIOBucketPolicy"`
}

type UploadInfo struct {
	Hash string `json:"hash"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
	Size int64  `json:"size"`
	Path string `json:"path"`
}

//var MinIOAccessKeyID = "minio123"
//var MinIOAccessSecretKey = "minio123"
//var MinIOEndpoint = "192.168.204.130:9000"
//var MinIOBucket = "wttest"
//var MinIOBucketLocation = "beijing"
//var MinIOSSLBool = false
//var MINIO_BASE_PATH = "breakpoint"
//// BucketPolicy 设置存储桶权限
//var MinIOBucketPolicy = "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:GetBucketLocation\",\"s3:ListBucket\",\"s3:ListBucketMultipartUploads\"],\"Resource\":[\"arn:aws:s3:::%s\"]},{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:AbortMultipartUpload\",\"s3:DeleteObject\",\"s3:GetObject\",\"s3:ListMultipartUploadParts\",\"s3:PutObject\"],\"Resource\":[\"arn:aws:s3:::%s/*\"]}]}"
