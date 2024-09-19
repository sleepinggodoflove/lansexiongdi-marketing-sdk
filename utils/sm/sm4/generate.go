package sm4

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/tjfoc/gmsm/sm4"
	"log"
)

func generateKey() string {
	key := make([]byte, sm4.BlockSize)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
