package sm

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm2"
)

// Cipher 使用公钥加密数据
func Cipher(publicKey *sm2.PublicKey, plaintext []byte) (string, error) {
	ciphertext, err := sm2.Encrypt(publicKey, plaintext, rand.Reader, sm2.C1C2C3)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Plain 使用私钥解密数据
func Plain(privateKey *sm2.PrivateKey, ciphertext string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	decryptedText, err := privateKey.Decrypt(rand.Reader, ciphertextBytes, sm2.C1C2C3)
	if err != nil {
		return "", err
	}
	return string(decryptedText), nil
}
