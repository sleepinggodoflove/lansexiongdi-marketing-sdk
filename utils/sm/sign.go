package sm

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
)

func Sign(data string, privateKey *sm2.PrivateKey) (string, error) {
	signatureBytes, err := privateKey.Sign(rand.Reader, []byte(data), nil)
	if err != nil {
		return "", fmt.Errorf("sign failed: %v", err)
	}
	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

func Verify(data string, signature string, publicKey *sm2.PublicKey) (bool, error) {
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, fmt.Errorf("signature base64 decode failed: %v", err)
	}
	if !publicKey.Verify([]byte(data), signatureBytes) {
		fmt.Println("Signature verification failed")
		return false, nil
	}
	return true, nil
}
