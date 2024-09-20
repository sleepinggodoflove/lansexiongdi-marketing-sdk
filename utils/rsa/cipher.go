package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func Cipher(publicKey *rsa.PublicKey, plaintext []byte) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
	if err != nil {
		return "", fmt.Errorf("error encrypting data: %v", err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Plain(privateKey *rsa.PrivateKey, ciphertext string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("error decoding base64: %v", err)
	}
	decryptedData, err := privateKey.Decrypt(nil, ciphertextBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return "", fmt.Errorf("error decrypting data: %v", err)
	}
	return string(decryptedData), nil
}
