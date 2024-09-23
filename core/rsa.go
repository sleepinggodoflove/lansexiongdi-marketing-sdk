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

func (s *RsaSigner) Sign(data string) (string, error) {
	return sdkrsa.Sign(data, s.privateKey)
}

func (v *RsaVerifier) Verify(data, signature string) (bool, error) {
	return sdkrsa.Verify(data, v.publicKey, signature)
}
