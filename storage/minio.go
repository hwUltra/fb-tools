package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioOSS struct {
	conf        MinIoConf
	minioClient *minio.Client
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

func NewMinioOSS(conf MinIoConf) *MinioOSS {
	minioClient, _ := minio.New(
		conf.MinIOEndpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(conf.MinIOAccessKeyID, conf.MinIOAccessSecretKey, ""),
			Secure: conf.MinIOSSLBool,
		})
	return &MinioOSS{
		conf:        conf,
		minioClient: minioClient,
	}
}

func (m *MinioOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {

	ext := path.Ext(fileHeader.Filename)
	objectName := fmt.Sprintf("%02d/%02d/%02d/",
		time.Now().Year(), time.Now().Month(), time.Now().Day()) + uuid.New().String() + ext

	_, err := m.minioClient.PutObject(context.Background(), m.conf.MinIOBucket, objectName, file, fileHeader.Size,
		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", m.conf.MinIOBasePath, objectName), nil
}

func (m *MinioOSS) DeleteFile(key string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := m.minioClient.RemoveObject(ctx, m.conf.MinIOBucket, key, minio.RemoveObjectOptions{})
	return err
}
