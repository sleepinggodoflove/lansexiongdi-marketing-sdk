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

func (s *SmSigner) Sign(data string) (string, error) {
	return sm.Sign(data, s.privateKey)
}

func (v *SmVerifier) Verify(data, signature string) (bool, error) {
	return sm.Verify(data, signature, v.publicKey)
}
