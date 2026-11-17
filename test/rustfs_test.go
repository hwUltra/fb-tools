package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hwUltra/fb-tools/storage"
	"github.com/hwUltra/fb-tools/utils"
)

func TestFs(t *testing.T) {
	//BOHb5yqElmzK44z8ecLk
	//lkO5lmnxRWfriTVBaXvyyOPuSsnU54GEErbPA4F4
	storageCfg := &storage.Config{
		Type:             "s3",
		Endpoint:         "http://127.0.0.1:9000",
		AccessKey:        "rustfs-admin",
		SecretKey:        "rustfs-admin",
		Bucket:           "public",
		Domain:           "https://file.imguo.com",
		PathNoBucketName: true,
		BasePath:         "",
	}
	if err := storage.InitDefaultStorage(storageCfg); err != nil {
		log.Fatalf("文件存储初始化失败: %v", err)
	}
	fmt.Println("文件存储初始化成功")
	file, err := os.Open("test.png")
	if err != nil {
		fmt.Println("打开失败:", err)
		return
	}
	fmt.Println("file = ", file, file.Name())
	filename := utils.GenerateFileName(file.Name())
	fmt.Println("filename = ", filename)
	path := utils.GeneratePath("uploads", filename)
	fmt.Println("path = ", path)
	contentType := utils.GetContentType(file.Name())
	fmt.Println("contentType = ", contentType)
	url, err := storage.DefaultStorage.Upload(path, file, contentType)
	if err != nil {
		fmt.Println("打开失败:", err)
	}
	fmt.Println("url = ", url)

}
