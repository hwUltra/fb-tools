package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// RsaGenKey 生成私钥/公钥，格式是PKCS#1
func RsaGenKey(bits int) (privateKey, publicKey []byte, err error) {
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return privateKey, publicKey, err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(rsaPrivateKey)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&rsaPrivateKey.PublicKey)

	if err != nil {
		return privateKey, publicKey, err
	}

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	publicBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return pem.EncodeToMemory(privateBlock), pem.EncodeToMemory(publicBlock), nil

}

// RsaSign 私钥签名
func RsaSign(data, privateKey string, hash crypto.Hash) (string, error) {
	privateKeyBlock, _ := pem.Decode([]byte(privateKey))
	if privateKeyBlock == nil {
		return "", errors.New("pem.Decode private key error！")
	}
	pk, err := toPrivateKey(privateKeyBlock)
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, pk, hash, hashed([]byte(data), hash))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// RsaVerifySign 公钥验签
func RsaVerifySign(data, sig, publicKey string, hash crypto.Hash) error {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return errors.New("decode public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	s, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), hash, hashed([]byte(data), hash), s)
}

// RsaEncrypt 公钥加密
func RsaEncrypt(data, publicKey string) (string, error) {
	publicKeyBlock, _ := pem.Decode([]byte(publicKey))
	if publicKeyBlock == nil {
		return data, errors.New("pem.Decode public key error")
	}
	keyInit, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return data, err
	}
	key := keyInit.(*rsa.PublicKey)

	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(data))
	if err != nil {
		return data, err
	}
	return base64.URLEncoding.EncodeToString(encryptBytes), nil
}

// RsaDecrypt 私钥解密
func RsaDecrypt(data, privateKey string) (string, error) {
	privateKeyBlock, _ := pem.Decode([]byte(privateKey))
	if privateKeyBlock == nil {
		return "", errors.New("pem.Decode private key error")
	}
	dataBytes, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	pk, err := toPrivateKey(privateKeyBlock)
	if err != nil {
		return "", err
	}
	encryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, pk, dataBytes)
	if err != nil {
		return "", err
	}
	return string(encryptBytes), nil
}

// 根据hash类型返回hash值
func hashed(data []byte, hash crypto.Hash) []byte {
	if hash == crypto.SHA256 {
		h := sha256.Sum256(data)
		return h[:]
	} else {
		h := sha1.Sum(data)
		return h[:]
	}
}

// 根据证书类型生成私钥对象
func toPrivateKey(privateKeyBlock *pem.Block) (*rsa.PrivateKey, error) {
	if privateKeyBlock.Type == "RSA PRIVATE KEY" {
		//PKCS#1
		return x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	} else if privateKeyBlock.Type == "PRIVATE KEY" {
		//PKCS#8
		key, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
		if err != nil {
			return nil, err
		}
		return key.(*rsa.PrivateKey), nil
	} else {
		return nil, errors.New("unknown pem private key type")
	}
}
