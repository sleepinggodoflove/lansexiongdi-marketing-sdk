package core

import (
	"errors"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/consts"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/interface"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/rsa"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/sm"
)

// SignerFactory to create signers and verifiers
type SignerFactory struct{}

// SignerVerifier create signers and verifiers
func (f *SignerFactory) SignerVerifier(signType string, s *Config) (_interface.Signer, _interface.Verifier, _interface.EncodeDecode, error) {
	switch signType {
	case consts.SignRSA:
		prk, err := rsa.PrivateKeyRsa(s.PrivateKey)
		if err != nil {
			return nil, nil, nil, err
		}
		puk, err := rsa.PublicKeyRsa(s.MerchantPublicKey)
		if err != nil {
			return nil, nil, nil, err
		}
		return &RsaSigner{privateKey: prk}, &RsaVerifier{publicKey: puk}, &RsaEncodeDecode{key: s.Key}, nil
	case consts.SignSM:
		prk, err := sm.PrivateKeySM(s.PrivateKey)
		if err != nil {
			return nil, nil, nil, err
		}
		puk, err := sm.PublicKeySM(s.MerchantPublicKey)
		if err != nil {
			return nil, nil, nil, err
		}
		return &SmSigner{privateKey: prk}, &SmVerifier{publicKey: puk}, &SmEncodeDecode{key: s.Key}, nil
	default:
		return nil, nil, nil, errors.New("signType,不支持的类型")
	}
}
