package sign

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
	"log"
)

func sm2SecretKey() (pri, puk string) {
	priKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal("秘钥生产失败：", err)
	}

	dd := x509.WritePrivateKeyToHex(priKey)

	bs, err := hex.DecodeString(dd)
	pri = fmt.Sprintf("私钥：%s", base64.StdEncoding.EncodeToString(bs))

	pubKey := &priKey.PublicKey
	ddp := x509.WritePublicKeyToHex(pubKey)

	bsp, err := hex.DecodeString(ddp)
	puk = fmt.Sprintf("\n公钥：%s", base64.StdEncoding.EncodeToString(bsp))

	return
}

func getPukByPrK(prk string) string {
	d, err := base64.StdEncoding.DecodeString(prk)
	if err != nil {
		log.Fatal(err)
	}

	pri, err := x509.ReadPrivateKeyFromHex(hex.EncodeToString(d))
	if err != nil {
		log.Fatal(err)
	}

	// 根据私钥生成公钥
	publicKey := &pri.PublicKey
	ddp := x509.WritePublicKeyToHex(publicKey)

	bsp, err := hex.DecodeString(ddp)
	if err != nil {
		log.Fatal("DecodeStrings异常：", err)
	}

	return base64.StdEncoding.EncodeToString(bsp)
}
