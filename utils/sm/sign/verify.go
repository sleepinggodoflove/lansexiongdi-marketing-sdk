package sign

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"log"
	"math/big"
)

func verify(content string, signature string, publicKeyStr string) bool {
	decodeSig, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.Fatal("signature base64 encode err:", err)
	}

	contentBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal("content base64 decode err:", err)
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyStr)
	if err != nil {
		log.Fatal("publicKeyStr base64 decode err:", err)
	}

	pubKey, err := x509.ReadPublicKeyFromHex(hex.EncodeToString(pubKeyBytes))
	if err != nil {
		log.Fatal("pubKeyBytes x509 ReadPublicKeyFromHex err:", err)
	}

	rAndS := rsPlainByteArrayToAsn1(decodeSig)
	r := new(big.Int).SetBytes(rAndS[:RsLen])
	s := new(big.Int).SetBytes(rAndS[RsLen:])

	return sm2.Sm2Verify(pubKey, contentBytes, nil, r, s)
}
