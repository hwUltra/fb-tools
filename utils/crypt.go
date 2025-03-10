package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func MD5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Md5 先base64，然后MD5
func Base64Md5(params string) string {
	return MD5V(base64.StdEncoding.EncodeToString([]byte(params)))
}
