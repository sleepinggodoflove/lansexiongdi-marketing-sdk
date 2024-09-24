package core

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/err"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/rsa"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/sm"
)

// SignerFactory to create signers and verifiers
type SignerFactory struct{}

// SignerVerifier create signers and verifiers
func (f *SignerFactory) SignerVerifier(signType string, s *Config) (Signer, Verifier, error) {
	switch signType {
	case SignRSA:
		prk, err := rsa.PrivateKeyRsa(s.PrivateKey)
		if err != nil {
			return nil, nil, err
		}
		puk, err := rsa.PublicKeyRsa(s.MerchantPublicKey)
		if err != nil {
			return nil, nil, err
		}
		return &RsaSigner{privateKey: prk}, &RsaVerifier{publicKey: puk}, nil
	case SignSM:
		prk, err := sm.PrivateKeySM(s.PrivateKey)
		if err != nil {
			return nil, nil, err
		}
		puk, err := sm.PublicKeySM(s.MerchantPublicKey)
		if err != nil {
			return nil, nil, err
		}
		return &SmSigner{privateKey: prk}, &SmVerifier{publicKey: puk}, nil
	default:
		return nil, nil, err.ErrUnsupportedSignType
	}
}