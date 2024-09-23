package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// SignerFactory to create signers and verifiers
type SignerFactory struct{}

func (f *SignerFactory) SignerVerifier(signType string, s *SecretKey) (Signer, Verifier, error) {
	switch signType {
	case MethodRSA:
		return nil, nil, nil
	case MethodSM:
		return &SmSigner{privateKey: s.PrivateKey}, &SmVerifier{publicKey: s.MerchantPublicKey}, nil
	default:
		return nil, nil, errors.New("unsupported SignType")
	}
}

type SecretKey struct {
	PrivateKey        string
	MerchantPublicKey string
}

// SDK structure
type SDK struct {
	AppID     string
	SignType  string
	Timestamp string
	Signer    Signer
	Verifier  Verifier
}

// NewSDK creates a new SDK instance
func NewSDK(s *SecretKey, appID, signType string) (*SDK, error) {
	factory := &SignerFactory{}
	signer, verifier, err := factory.SignerVerifier(signType, s)
	if err != nil {
		return nil, err
	}
	return &SDK{
		AppID:     appID,
		SignType:  signType,
		Timestamp: time.Now().Format(time.RFC3339),
		Signer:    signer,
		Verifier:  verifier,
	}, nil
}

// Request sends the request and verifies the response signature
func (s *SDK) Request(uri, ciphertext string) (*http.Response, error) {
	dataToSign := s.AppID + s.Timestamp + ciphertext
	signature, err := s.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}
	reqData := map[string]string{
		"AppId":      s.AppID,
		"SignType":   s.SignType,
		"Timestamp":  s.Timestamp,
		"Ciphertext": ciphertext,
		"Sign":       signature,
	}
	reqBody, _ := json.Marshal(reqData)

	resp, err := http.Post(uri, "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Assume response contains the signature and data
	var responseData struct {
		Data      string `json:"data"`
		Signature string `json:"sign"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}
	// Verify the response signature
	isValid, err := s.Verifier.Verify(responseData.Data, responseData.Signature)
	if err != nil || !isValid {
		return nil, errors.New("signature verification failed")
	}

	return resp, nil
}
