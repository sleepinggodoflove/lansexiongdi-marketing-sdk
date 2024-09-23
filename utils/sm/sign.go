package sm

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/x509"
)

func Sign(data, privateKeyStr string) (string, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil {
		return "", fmt.Errorf("privateKey base64 decode failed: %v", err)
	}

	privateKey, err := x509.ReadPrivateKeyFromHex(hex.EncodeToString(privateKeyBytes))
	if err != nil {
		return "", fmt.Errorf("read private key from hex failed: %v", err)
	}

	signatureBytes, err := privateKey.Sign(rand.Reader, []byte(data), nil)

	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

func Verify(data string, signature, publicKeyStr string) (bool, error) {
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("signature base64 decode failed: %v", err)
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		return false, fmt.Errorf("publicKeyStr base64 decode failed: %v", err)
	}

	publicKey, err := x509.ReadPublicKeyFromHex(hex.EncodeToString(pubKeyBytes))
	if err != nil {
		return false, fmt.Errorf("read public key from hex failed: %v", err)
	}

	if !publicKey.Verify([]byte(data), signatureBytes) {
		fmt.Println("Signature verification failed")
		return false, nil
	}

	return true, nil
}
