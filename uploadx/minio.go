package uploadx

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
	"path"
	"time"
)

type MinioOSS struct {
	conf        MinIoConf
	minioClient *minio.Client
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

func (m *MinioOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error) {

	ext := path.Ext(fileHeader.Filename)
	objectName := fmt.Sprintf("%02d/%02d/%02d/",
		time.Now().Year(), time.Now().Month(), time.Now().Day()) + uuid.New().String() + ext

	info, err := m.minioClient.PutObject(context.Background(), m.conf.MinIOBucket, objectName, file, fileHeader.Size,
		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
	if err != nil {
		return nil, err
	}

	return &UploadInfo{
		Path: objectName,
		Size: info.Size,
		Ext:  ext,
		Hash: info.ETag,
	}, nil
}

func (m *MinioOSS) DeleteFile(key string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := m.minioClient.RemoveObject(ctx, m.conf.MinIOBucket, key, minio.RemoveObjectOptions{})
	return err
}
