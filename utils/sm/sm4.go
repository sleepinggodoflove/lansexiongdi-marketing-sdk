package sm

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm4"
)

func GenerateSM4Key() (string, error) {
	key := make([]byte, sm4.BlockSize)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func Encode(contentBytes []byte, encryptKey string) (string, error) {
	d, err := base64.StdEncoding.DecodeString(encryptKey)
	if err != nil {
		return "", err
	}
	cipherBlock, err := sm4.NewCipher(d)
	if err != nil {
		return "", err
	}

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

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decode(encrypted, encryptKey string) (string, error) {
	d, err := base64.StdEncoding.DecodeString(encryptKey)
	if err != nil {
		return "", err
	}

	cipherBlock, err := sm4.NewCipher(d)
	if err != nil {
		return "", err
	}

	blockSize := cipherBlock.BlockSize()
	iv := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		iv[i] = 0
	}

	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	plainText := make([]byte, len(cipherText))
	blockMode := cipher.NewCBCDecrypter(cipherBlock, iv)
	blockMode.CryptBlocks(plainText, cipherText)

	plainTextLen := len(plainText)
	padding := int(plainText[plainTextLen-1])
	buff := plainText[:plainTextLen-padding]

	return string(buff), nil
}
