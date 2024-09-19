package sm4

import (
	"crypto/cipher"
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm4"
	"log"
)

func decrypt(encrypted string, encryptKey string) string {
	// 解码密钥
	d, err := base64.StdEncoding.DecodeString(encryptKey)
	if err != nil {
		log.Fatal(err)
	}

	// 创建密码器
	cipherBlock, err := sm4.NewCipher(d)
	if err != nil {
		log.Fatal(err)
	}

	// 设置解密模式和初始向量（IV）
	blockSize := cipherBlock.BlockSize()
	iv := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		iv[i] = 0
	}

	// 解密数据
	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		log.Fatal(err)
	}

	plainText := make([]byte, len(cipherText))
	blockMode := cipher.NewCBCDecrypter(cipherBlock, iv)
	blockMode.CryptBlocks(plainText, cipherText)

	// 去掉填充，并返回原始明文
	plainTextLen := len(plainText)
	padding := int(plainText[plainTextLen-1])
	buff := plainText[:plainTextLen-padding]

	return string(buff)
}
