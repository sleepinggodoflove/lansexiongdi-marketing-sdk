package core

import (
	"errors"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/rsa"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/sm"
)

// SignerFactory to create signers and verifiers
type SignerFactory struct{}

// SignerVerifier create signers and verifiers
func (f *SignerFactory) SignerVerifier(signType string, s *Config) (Signer, Verifier, EncodeDecode, error) {
	switch signType {
	case SignRSA:
		prk, err := rsa.PrivateKeyRsa(s.PrivateKey)
		if err != nil {
			return nil, nil, nil, err
		}
		puk, err := rsa.PublicKeyRsa(s.PublicKey)
		if err != nil {
			return nil, nil, nil, err
		}
		return &RsaSigner{privateKey: prk}, &RsaVerifier{publicKey: puk}, &RsaEncodeDecode{key: s.Key}, nil
	case SignSM:
		prk, err := sm.PrivateKeySM(s.PrivateKey)
		if err != nil {
			return nil, nil, nil, err
		}
		puk, err := sm.PublicKeySM(s.PublicKey)
		if err != nil {
			return nil, nil, nil, err
		}
		return &SmSigner{privateKey: prk}, &SmVerifier{publicKey: puk}, &SmEncodeDecode{key: s.Key}, nil
	default:
		return nil, nil, nil, errors.New("signType,不支持的类型")
	}
}
