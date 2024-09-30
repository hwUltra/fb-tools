package test

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"path/filepath"
	"testing"
	"time"
)

func TestMinioXBuilder(t *testing.T) {
	//minioX := &miniox.MinioX{Conf: miniox.Conf{
	//	MinIOAccessKeyID: ""}}

	//minioX.MinIOUpload()

}

func TestMinioBuilder(t *testing.T) {

	ctx := context.Background()
	endpoint := "192.168.3.88:9000"
	accessKeyID := "c4h0z6z3GxY8UXJkb7Fd"
	secretAccessKey := "2ej0kOfnNNxt5Jlz6pIw98aGuqUOSOg0BGqyV5L2"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called mymusic.
	bucketName := "files"
	location := "cn-north-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the zip file
	filePath := "a.png"
	fileExt := filepath.Ext(filePath)
	objectName := fmt.Sprintf("%02d/%02d/%02d/%s%s",
		time.Now().Year(), time.Now().Month(), time.Now().Day(), uuid.New().String(), fileExt)

	//contentType := "application/zip"
	contentType := "binary/octet-stream"
	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("minioClient.FPutObject", info, err)
	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

}
