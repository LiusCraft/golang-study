package util

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

func Base64DesDecrypt(inputStr, desKey string) (string, error) {
	//<1> 进行base64转码解密
	// 解码一个base64编码可能返回一个错误，如果你不知道输入是否是正确的base64编码，你需要检测一些解码错误
	b64origData, err := base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		return "", err
	}

	//<2> 进行DES解密，解析出原始数据
	origData, err2 := DesDecrypt(b64origData, []byte(desKey))
	if err2 != nil {
		return "", err2
	}

	return string(origData), nil
}

// des解密
func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	//origData := make([]byte, len(crypted))
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)

	origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
