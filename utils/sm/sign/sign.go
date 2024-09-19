package sign

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"log"
)

func sign(content, privateKey string) string {
	contentBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		log.Fatal(err)
	}

	d, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	pri, err := x509.ReadPrivateKeyFromHex(hex.EncodeToString(d))
	if err != nil {
		log.Fatal(err)
	}

	r, s, err := sm2.Sm2Sign(pri, contentBytes, nil, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	rs := rsAsn1ToPlainByteArray(r, s)

	return base64.StdEncoding.EncodeToString(rs)
}
