package rsa

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Decode 解密函数
func Decode(code, key string) string {
	// Base64解码
	encryptString, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return ""
	}

	// 创建AES解密器
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	// 创建ECB模式解密器
	mode := NewECDecrypted(block)

	// 解密
	decrypted := make([]byte, len(encryptString))
	mode.CryptBlocks(decrypted, encryptString)

	// 去除填充
	padding := decrypted[len(decrypted)-1]
	return string(decrypted[:len(decrypted)-int(padding)])
}

// Encode 加密函数
func Encode(code, key string) string {
	// 创建AES加密器
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	// PKCS7填充
	blockSize := block.BlockSize()
	padding := blockSize - len(code)%blockSize
	padText := string(byte(padding))
	for i := 0; i < padding; i++ {
		code += padText
	}

	// 创建ECB模式加密器
	mode := NewECEncrypted(block)

	// 计算加密后数据的长度
	encrypted := make([]byte, len(code))
	mode.CryptBlocks(encrypted, []byte(code))

	// Base64编码
	return base64.StdEncoding.EncodeToString(encrypted)
}

// ecEncrypted ECB加密器
type ecEncrypted struct {
	b         cipher.Block
	blockSize int
}

func NewECEncrypted(b cipher.Block) cipher.BlockMode {
	return &ecEncrypted{b, b.BlockSize()}
}

func (x *ecEncrypted) BlockSize() int { return x.blockSize }

func (x *ecEncrypted) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("src not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

// ECB解密器
type ecDecrypted struct {
	b         cipher.Block
	blockSize int
}

func NewECDecrypted(b cipher.Block) cipher.BlockMode {
	return &ecDecrypted{b, b.BlockSize()}
}

func (x *ecDecrypted) BlockSize() int { return x.blockSize }

func (x *ecDecrypted) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("src not full blocks")
	}
	if len(dst) < len(src) {
		panic("output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
