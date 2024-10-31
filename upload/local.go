package upload

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

type Local struct {
	StorePath string
	Path      string
}

func (u *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V(name)
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(u.StorePath, os.ModePerm)
	if mkdirErr != nil {
		return "", "", mkdirErr
	}
	// 拼接路径和文件名
	p := u.StorePath + "/" + filename
	filepath := u.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		return "", "", openError
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		return "", "", createErr
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", "", copyErr
	}
	return filepath, filename, nil
}

func (u *Local) DeleteFile(key string) error {
	// 检查 key 是否为空
	if key == "" {
		return errors.New("key不能为空")
	}

	// 验证 key 是否包含非法字符或尝试访问存储路径之外的文件
	if strings.Contains(key, "..") || strings.ContainsAny(key, `\/:*?"<>|`) {
		return errors.New("非法的key")
	}

	p := filepath.Join(u.StorePath, key)

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
