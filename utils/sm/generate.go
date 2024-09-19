package sm

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func GenerateKey() (pri, puk string, err error) {
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		return pri, puk, fmt.Errorf("秘钥生成失败:%v", err)
	}

	priKeyBytes, err := hex.DecodeString(x509.WritePrivateKeyToHex(privateKey))
	pri = base64.StdEncoding.EncodeToString(priKeyBytes)

	publicKeyBytes, err := hex.DecodeString(x509.WritePublicKeyToHex(&privateKey.PublicKey))
	puk = base64.StdEncoding.EncodeToString(publicKeyBytes)

	return
}

func GetPukByPrK(prk string) (string, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(prk)
	if err != nil {
		return "", fmt.Errorf("私钥base64解码失败:%v", err)
	}

	privateKey, err := x509.ReadPrivateKeyFromHex(hex.EncodeToString(privateKeyBytes))
	if err != nil {
		return "", fmt.Errorf("私钥hex解码失败:%v", err)
	}

	// 根据私钥得到公钥
	publicKeyBytes, err := hex.DecodeString(x509.WritePublicKeyToHex(&privateKey.PublicKey))
	if err != nil {
		return "", fmt.Errorf("公钥hex解码失败:%v", err)
	}

	return base64.StdEncoding.EncodeToString(publicKeyBytes), nil
}
