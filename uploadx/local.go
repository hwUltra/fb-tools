package uploadx

import (
	"errors"
	"github.com/hwUltra/fb-tools/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

type LocalOSS struct {
	conf LocalConf
}

func NewLocalOSS(conf LocalConf) *LocalOSS {
	return &LocalOSS{
		conf: conf,
	}
}

func (m *LocalOSS) UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadInfo, error) {

	ext := filepath.Ext(fileHeader.Filename)
	name := strings.TrimSuffix(fileHeader.Filename, ext)
	name = utils.MD5V(name)
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	err := os.MkdirAll(m.conf.StorePath, os.ModePerm)
	if err != nil {
		return nil, err
	}
	p := m.conf.StorePath + "/" + filename
	filePath := m.conf.Path + "/" + filename

	out, err := os.Create(p)
	if err != nil {
		return nil, err
	}
	defer out.Close() // 创建文件 defer 关闭

	_, err = io.Copy(out, file) // 传输（拷贝）文件
	if err != nil {
		return nil, err
	}
	return &UploadInfo{
		Path: filePath,
		Name: filename,
		Size: fileHeader.Size,
		Ext:  ext,
	}, nil
}

func (m *LocalOSS) DeleteFile(key string) error {
	// 检查 key 是否为空
	if key == "" {
		return errors.New("key不能为空")
	}

	// 验证 key 是否包含非法字符或尝试访问存储路径之外的文件
	if strings.Contains(key, "..") || strings.ContainsAny(key, `\/:*?"<>|`) {
		return errors.New("非法的key")
	}
	p := filepath.Join(m.conf.StorePath, key)
	// 检查文件是否存在
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return errors.New("文件不存在")
	}
	// 使用文件锁防止并发删除
	mu.Lock()
	defer mu.Unlock()
	err := os.Remove(p)
	if err != nil {
		return errors.New("文件删除失败: " + err.Error())
	}

	return nil
}

// 返回值说明：
//	7z、exe、doc 类型会返回 application/octet-stream  未知的文件类型
//	jpg	=>	image/jpeg
//	png	=>	image/png
//	ico	=>	image/x-icon
//	bmp	=>	image/bmp
//  xlsx、docx 、zip	=>	application/zip
//  tar.gz	=>	application/x-gzip
//  txt、json、log等文本文件	=>	text/plain; charset=utf-8   备注：就算txt是gbk、ansi编码，也会识别为utf-8
//
//// GetFilesMimeByFileName 通过文件名获取文件mime信息
//func GetFilesMimeByFileName(filepath string) string {
//	f, err := os.Open(filepath)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer func(f *os.File) {
//		err := f.Close()
//		if err != nil {
//
//		}
//	}(f)
//
//	// 只需要前 32 个字节就可以了
//	buffer := make([]byte, 32)
//	if _, err := f.Read(buffer); err != nil {
//		return ""
//	}
//
//	return http.DetectContentType(buffer)
//}
//
//// GetFilesMimeByFp 通过文件指针获取文件mime信息
//func GetFilesMimeByFp(fp multipart.File) string {
//
//	buffer := make([]byte, 32)
//	if _, err := fp.Read(buffer); err != nil {
//
//		return ""
//	}
//
//	return http.DetectContentType(buffer)
//}
