package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
)

func PrivateKeyRSA(privateKeyStr string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("解码 Base64 编码的 RSA 私钥字符串: %v", err)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("解析 RSA 私钥: %v", err)
	}
	return privateKey, nil
}

func PrivateKeyPem(privateKeyStr string) (string, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return "", fmt.Errorf("解码 Base64 编码的 RSA 私钥字符串: %v", err)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("解析 RSA 私钥: %v", err)
	}

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	return string(pem.EncodeToMemory(privateKeyPEM)), nil
}

func PublicKeyRSA(publicKeyStr string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return nil, fmt.Errorf("解码 Base64 编码的 RSA 公钥字符串: %v", err)
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("解析 RSA 公钥: %v", err)
	}

	return publicKey, nil
}

func PublicKeyPem(publicKeyStr string) (string, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return "", fmt.Errorf("解码 Base64 编码的 RSA 公钥字符串: %v", err)
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBytes)
	if err != nil {
		log.Fatal("解析 RSA 公钥:", err)
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	return string(pem.EncodeToMemory(publicKeyPEM)), nil
}
