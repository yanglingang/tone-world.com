package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
)

var decrypted string

func init() {
	flag.StringVar(&decrypted, "d", "", "加密过的数据")
	flag.Parse()
}

func main() {
	var data []byte
	var err error
	if decrypted != "" {
		data, err = base64.StdEncoding.DecodeString(decrypted)
		if err != nil {
			panic(err)
		}
	} else {
		data, err = RsaEncrypt([]byte("加密过的数据0"))
		if err != nil {
			panic(err)
		}
		fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	}
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDIZ8N9pCbIakC2IbKtbfA7uocWDYwkyuE/8zWYerzyV3eU0hei
9PuS0Qsf5lt5B7StBEsEeEhys32BX0sM4pq1FcZjcRF+9rcEhd7/tXnN9Wg5EZQK
SQBeOtVovTX2uBsAHdjtFFD3HW5UmfWWuuEl1lWjh06Q+Q6JbWzmvvd4JQIDAQAB
AoGBAMFCT2N6SWw8CuuY05Yrzt/KoTrDFcLlYxMolybUNiH993Ospt6fIXwT24aH
vu6YX8P5v94voK38Kav5GYbzf/wpiAP6liY2fts8dci+zXkAM7NExClJ2pxZPB7x
9JiP/zOttTAcU86wFWDR0iwkLoVSBQf8rqJjo8813JQJ8dKBAkEA/qOjD1UwrhzG
IBPukLlM9JQe5iiwsr259SFjw9pNzJry0fYVL+TZuFt0wc3JtmfjUB9DlZLtp/w6
uAM/jN25awJBAMl57mLtCi6HP1Lt6eo0B10yQd0p1bTVvE0fPz8DvBXKLgm+bcq9
3E9H80Ibn/Ld77HaPTL4TdKqqh3WwNWHKK8CQHFib/Md3eVJjrct6OasfCXT5sZZ
jASrEqiiS4gkJsxampD/YIPJBWFf5+d8OLtuGvvMUA3ENOq+F29kkuGowS0CQEZx
bVoFSuQNwaQ7LugGHPUG12R+dgvuFxJX9IMRyTdNI1+gxz51t4u4umLIydneoynq
Bi/GBV+88BHSvkVqJl0CQC2RSojuNCX9fZemtblOHsl6X1BU1lv9DQvHbu0XTrXK
cVufoj9uTLQACxSbL6qFZyOlqRy6H3BbVeV17T9RVSo=
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDIZ8N9pCbIakC2IbKtbfA7uocW
DYwkyuE/8zWYerzyV3eU0hei9PuS0Qsf5lt5B7StBEsEeEhys32BX0sM4pq1FcZj
cRF+9rcEhd7/tXnN9Wg5EZQKSQBeOtVovTX2uBsAHdjtFFD3HW5UmfWWuuEl1lWj
h06Q+Q6JbWzmvvd4JQIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
