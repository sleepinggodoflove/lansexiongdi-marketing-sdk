package sm4

import (
	"crypto/cipher"
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm4"
	"log"
)

func encrypt(content string, encryptKey string) string {
	// 传输过来的数据base64 encode 的，因为传输过程中，
	// 特殊字符串可能会丢失，比如：json字符串的引号就会丢失
	// 所以需要将传输过来的数据进行 base64 decode
	contentBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal(err)
	}

	d, err := base64.StdEncoding.DecodeString(encryptKey)
	if err != nil {
		log.Fatal(err)
	}

	// 创建密码器
	cipherBlock, err := sm4.NewCipher(d)
	if err != nil {
		log.Fatal(err)
	}

	// 设置加密模式和初始向量（IV）
	blockSize := cipherBlock.BlockSize()
	iv := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		iv[i] = 0
	}
	blockMode := cipher.NewCBCEncrypter(cipherBlock, iv)

	padding := blockSize - len(contentBytes)%blockSize
	for i := 0; i < padding; i++ {
		contentBytes = append(contentBytes, byte(padding))
	}
	cipherText := make([]byte, len(contentBytes))
	blockMode.CryptBlocks(cipherText, contentBytes)

	// 对加密结果使用base64进行编码
	return base64.StdEncoding.EncodeToString(cipherText)
}
