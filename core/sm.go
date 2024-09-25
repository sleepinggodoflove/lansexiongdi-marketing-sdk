package core

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/sm"
	"github.com/tjfoc/gmsm/sm2"
)

// SmSigner for SM signing (国密)
type SmSigner struct {
	privateKey *sm2.PrivateKey
}

// SmVerifier for SM verification (国密)
type SmVerifier struct {
	publicKey *sm2.PublicKey
}

// SmEncodeDecode .
type SmEncodeDecode struct {
	key string
}

func (s *SmSigner) Sign(data string) (string, error) {
	return sm.Sign(data, s.privateKey)
}

func (s *SmVerifier) Verify(data, signature string) bool {
	b, err := sm.Verify(data, signature, s.publicKey)
	if err != nil {
		return false
	}
	return b
}

func (s *SmEncodeDecode) Encode(data string) (string, error) {
	return sm.Encode([]byte(s.key), data)
}

func (s *SmEncodeDecode) Decode(data string) (string, error) {
	return sm.Decode(s.key, data)
}
