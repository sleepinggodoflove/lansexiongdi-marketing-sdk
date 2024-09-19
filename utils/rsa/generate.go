package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"os"
)

type Generate struct {
	publicKeyBytes  []byte
	privateKeyBytes []byte
}

func NewGenerateKey() *Generate {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	return &Generate{
		publicKeyBytes:  x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
		privateKeyBytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
}

func (g *Generate) SavePem(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}

	privateKeyFile, err := os.Create(path + "/private.pem")
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()
	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: g.privateKeyBytes})
	if err != nil {
		return err
	}

	publicKeyFile, err := os.Create(path + "/public.pem")
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()
	return pem.Encode(publicKeyFile, &pem.Block{Type: "RSA PUBLIC KEY", Bytes: g.publicKeyBytes})
}

func (g *Generate) GetKey() (privateKeyStr, publicKeyStr string) {
	privateKeyStr = base64.StdEncoding.EncodeToString(g.privateKeyBytes)
	publicKeyStr = base64.StdEncoding.EncodeToString(g.publicKeyBytes)
	return
}
