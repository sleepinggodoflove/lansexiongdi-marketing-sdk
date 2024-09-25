package core

import (
	"crypto/rsa"
	sdkrsa "github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/rsa"
)

// RsaSigner for RSA signing
type RsaSigner struct {
	privateKey *rsa.PrivateKey
}

// RsaVerifier for RSA verification
type RsaVerifier struct {
	publicKey *rsa.PublicKey
}

// RsaEncodeDecode .
type RsaEncodeDecode struct {
	key string
}

func (r *RsaSigner) Sign(data string) (string, error) {
	return sdkrsa.Sign(data, r.privateKey)
}

func (r *RsaVerifier) Verify(data, signature string) bool {
	b, err := sdkrsa.Verify(data, signature, r.publicKey)
	if err != nil {
		return false
	}
	return b
}

func (r *RsaEncodeDecode) Encode(data string) (string, error) {
	return sdkrsa.Encode(r.key, data), nil
}

func (r *RsaEncodeDecode) Decode(data string) (string, error) {
	return sdkrsa.Decode(r.key, data), nil
}
