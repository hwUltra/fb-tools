package miniox

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/http"
	"path"
	"time"
)

type MinioX struct {
	conf        MinioConf
	minioClient *minio.Client
}

func NewMinioX(conf MinioConf) *MinioX {
	minioClient, _ := minio.New(
		conf.MinIOEndpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(conf.MinIOAccessKeyID, conf.MinIOAccessSecretKey, ""),
			Secure: conf.MinIOSSLBool,
		})
	return &MinioX{
		conf:        conf,
		minioClient: minioClient,
	}
}

// MinIOUpload 上传到自建的minio中
func (m *MinioX) MinIOUpload(r *http.Request) (*UploadInfo, error) {

	//获取文件信息
	file, fileHeader, err := r.FormFile(m.conf.MinIOFile)
	if err != nil {
		return nil, err
	}

	ext := path.Ext(fileHeader.Filename)
	objectName := fmt.Sprintf("%02d/%02d/%02d/",
		time.Now().Year(), time.Now().Month(), time.Now().Day()) + uuid.New().String() + ext

	info, err := m.minioClient.PutObject(context.Background(), m.conf.MinIOBucket, objectName, file, fileHeader.Size,
		minio.PutObjectOptions{ContentType: "binary/octet-stream"})
	fmt.Println("err03", err)
	return &UploadInfo{
		Path: info.Bucket + "/" + info.Key,
		Name: objectName,
		Size: info.Size,
		Ext:  ext,
		Hash: info.ETag,
	}, err
}

//// MinioInitPart Minio 分片上传初始化
//func (m *MinioX) MinioInitPart(ext string) (string, string, error) {
//	// Instantiate new minio client object.
//	core, err := minio.NewCore(
//		m.Conf.MinIOEndpoint,
//		&minio.Options{
//			Creds:  credentials.NewStaticV4(m.Conf.MinIOAccessKeyID, m.Conf.MinIOAccessSecretKey, ""),
//			Secure: m.Conf.MinIOSSLBool,
//		})
//	if err != nil {
//		return "", "", err
//	}
//	key := "breakpoint/" + UUID() + ext
//	uuid := UUID()
//	bucketName := m.Conf.MinIOBucket
//	objectName := strings.TrimPrefix(path.Join(m.Conf.MinIOBasePath, path.Join(uuid[0:1], uuid[1:2], uuid)), "/")
//
//	// objectContentType := "binary/octet-stream"
//	// metadata := make(map[string]string)
//	// metadata["Content-Type"] = objectContentType
//	// putopts := minio.PutObjectOptions{
//	// 	UserMetadata: metadata,
//	// }
//	uploadID, err := core.NewMultipartUpload(context.Background(), bucketName, objectName, minio.PutObjectOptions{})
//	if err != nil {
//		return "", "", err
//	}
//	return key, uploadID, nil
//	// return core.NewMultipartUpload(bucketName, objectName, minio.PutObjectOptions{})
//}
//
//// MinioPartUpload 分片上传
//func (m *MinioX) MinioPartUpload(r *http.Request) (string, error) {
//	// Instantiate new minio client object.
//	core, err := minio.NewCore(
//		m.Conf.MinIOEndpoint,
//		&minio.Options{
//			Creds:  credentials.NewStaticV4(m.Conf.MinIOAccessKeyID, m.Conf.MinIOAccessSecretKey, ""),
//			Secure: m.Conf.MinIOSSLBool,
//		})
//	if err != nil {
//		return "", err
//	}
//
//	key := r.PostForm.Get("key")
//	UploadID := r.PostForm.Get("upload_id")
//	partNumber, err := strconv.Atoi(r.PostForm.Get("part_number"))
//
//	if err != nil {
//		return "", err
//	}
//
//	f, _, err := r.FormFile("file")
//	if err != nil {
//		return "", nil
//	}
//
//	buf := bytes.NewBuffer(nil)
//	io.Copy(buf, f)
//
//	data := bytes.NewReader(buf.Bytes())
//	dataLen := int64(len(buf.Bytes()))
//	// PutObjectPartOptions contains options for PutObjectPart API
//	// type PutObjectPartOptions struct {
//	// 	Md5Base64, Sha256Hex  string
//	// 	SSE                   encrypt.ServerSide
//	// 	CustomHeader, Trailer http.Header
//	// }
//	// opt可选
//	objectPart, err := core.PutObjectPart(context.Background(), m.Conf.MinIOBucket, key, UploadID, partNumber, data, dataLen, minio.PutObjectPartOptions{})
//	if err != nil {
//		return "", err
//	}
//
//	return objectPart.ETag, nil
//}
//
//type completeMultipartUpload struct {
//	XMLName xml.Name             `xml:"http://s3.amazonaws.com/doc/2006-03-01/ CompleteMultipartUpload" json:"-"`
//	Parts   []minio.CompletePart `xml:"Part"`
//}
//type ComplPart struct {
//	PartNumber int    `json:"partNumber"`
//	ETag       string `json:"eTag"`
//}
//
//type completedParts []minio.CompletePart
//
//func (a completedParts) Len() int           { return len(a) }
//func (a completedParts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a completedParts) Less(i, j int) bool { return a[i].PartNumber < a[j].PartNumber }
//
//type CompleteParts struct {
//	Data []ComplPart `json:"completedParts"`
//}
//
//// MinioPartUploadComplete 分片上传完成
//func (m *MinioX) MinioPartUploadComplete(key string, uploadID string, mo []minio.CompletePart) error {
//	// Instantiate new minio client object.
//	core, err := minio.NewCore(
//		m.Conf.MinIOEndpoint,
//		&minio.Options{
//			Creds:  credentials.NewStaticV4(m.Conf.MinIOAccessKeyID, m.Conf.MinIOAccessSecretKey, ""),
//			Secure: m.Conf.MinIOSSLBool,
//		})
//	if err != nil {
//		return err
//	}
//	uuid := uuid.NewV4().String()
//	bucketName := m.Conf.MinIOBucket
//	objectName := strings.TrimPrefix(path.Join(m.Conf.MinIOBasePath, path.Join(uuid[0:1], uuid[1:2], uuid)), "/")
//	// var client *minio_ext.Client
//	// partInfos, err := client.ListObjectParts(bucketName, objectName, uploadID)
//	// if err != nil {
//	// 	return err
//	// }
//
//	var complMultipartUpload completeMultipartUpload
//	// for _, partInfo := range partInfos {
//	// 	complMultipartUpload.Parts = append(complMultipartUpload.Parts, minio.CompletePart{
//	// 		PartNumber: partInfo.PartNumber,
//	// 		ETag:       partInfo.ETag,
//	// 	})
//
//	// }
//	complMultipartUpload.Parts = append(complMultipartUpload.Parts, mo...)
//
//	objectContentType := "binary/octet-stream"
//	metadata := make(map[string]string)
//	metadata["Content-Type"] = objectContentType
//	putopts := minio.PutObjectOptions{
//		UserMetadata: metadata,
//	}
//	// Sort all completed parts.
//	sort.Sort(completedParts(complMultipartUpload.Parts))
//	_, err = core.CompleteMultipartUpload(context.Background(), bucketName, objectName, uploadID, complMultipartUpload.Parts, putopts)
//	if err != nil {
//		return err
//	}
//
//	return err
//}
//
//func (m *MinioX) genMultiPartSignedUrl(uuid string, uploadId string, partNumber int, partSize int64) (*url.URL, error) {
//	minioClient, err := minio.New(m.Conf.MinIOEndpoint, &minio.Options{
//		Creds: credentials.NewStaticV4(m.Conf.MinIOAccessKeyID, m.Conf.MinIOAccessSecretKey, ""),
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	bucketName := m.Conf.MinIOBucket
//	objectName := strings.TrimPrefix(path.Join(m.Conf.MinIOBasePath, path.Join(uuid[0:1], uuid[1:2], uuid)), "/")
//	method := http.MethodPost
//	expires := time.Hour * 24 * 7
//	reqParams := make(url.Values)
//	return minioClient.Presign(context.Background(), method, bucketName, objectName, expires, reqParams)
//}
