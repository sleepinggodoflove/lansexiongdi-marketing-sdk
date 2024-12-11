package core

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/rsa"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils/sm"
)

type SignType string

const (
	SignRSA SignType = "RSA"
	SignSM  SignType = "SM"
)

// Config merchant app Config
type Config struct {
	AppID      string   `validate:"required"`
	PrivateKey string   `validate:"required"`
	PublicKey  string   `validate:"required"`
	Key        string   `validate:"required"`
	SignType   SignType `validate:"required"`
	BaseURL    string   `validate:"required,url"`
}

func (c *Config) Validate() error {
	if err := validator.New().Struct(c); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}

func (s SignType) IsRSA() bool {
	return s == SignRSA
}

func (s SignType) IsSM() bool {
	return s == SignSM
}

func (c *Config) CryptographySuite() (*CryptographySuite, error) {
	if c.SignType.IsRSA() {
		return c.CryptographySuiteRSA()
	}

	if c.SignType.IsSM() {
		return c.CryptographySuiteSM()
	}

	return nil, fmt.Errorf("[%s] invalid sign type", c.SignType)
}

func (c *Config) CryptographySuiteRSA() (*CryptographySuite, error) {
	prk, err := rsa.PrivateKeyRSA(c.PrivateKey)
	if err != nil {
		return nil, err
	}

	puk, err := rsa.PublicKeyRSA(c.PublicKey)
	if err != nil {
		return nil, err
	}

	return &CryptographySuite{
		Signer:   &RsaSigner{privateKey: prk},
		Verifier: &RsaVerifier{publicKey: puk},
		Cipher:   &RsaEncodeDecode{key: c.Key},
	}, nil
}

func (c *Config) CryptographySuiteSM() (*CryptographySuite, error) {
	prk, err := sm.PrivateKeySM(c.PrivateKey)
	if err != nil {
		return nil, err
	}

	puk, err := sm.PublicKeySM(c.PublicKey)
	if err != nil {
		return nil, err
	}

	return &CryptographySuite{
		Signer:   &SmSigner{privateKey: prk},
		Verifier: &SmVerifier{publicKey: puk},
		Cipher:   &SmEncodeDecode{key: c.Key},
	}, nil
}
