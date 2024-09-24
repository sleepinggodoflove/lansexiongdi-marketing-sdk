package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	AppID             string
	PrivateKey        string
	MerchantPublicKey string
	BaseURL           string
}

// SDK structure
type SDK struct {
	Config     *Config
	httpClient *http.Client
	SignType   string
	Signer     Signer
	Verifier   Verifier
}

type Option func(*SDK)

func WithSignType(signType string) Option {
	return func(s *SDK) {
		s.SignType = signType
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(s *SDK) {
		s.httpClient = client
	}
}

// NewSDK creates a new SDK instance
func NewSDK(s *Config, o ...Option) (*SDK, error) {
	if s == nil {
		return nil, errors.New("config is nil")
	}
	sdk := &SDK{
		SignType: SignRSA,
		Config:   s,
	}
	for _, f := range o {
		f(sdk)
	}
	factory := &SignerFactory{}
	signer, verifier, err := factory.SignerVerifier(sdk.SignType, s)
	if err != nil {
		return nil, err
	}
	sdk.Signer = signer
	sdk.Verifier = verifier
	return sdk, nil
}

// Request sends the request and verifies the response signature
func (s *SDK) Request(method, ciphertext string) (*http.Response, error) {
	timestamps := time.Now().Format(time.RFC3339)
	dataToSign := s.Config.AppID + timestamps + ciphertext
	signature, err := s.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}
	reqData := map[string]string{
		"app_id":     s.Config.AppID,
		"sign_type":  s.SignType,
		"timestamp":  timestamps,
		"ciphertext": ciphertext,
		"sign":       signature,
	}
	reqBody, _ := json.Marshal(reqData)

	resp, err := http.Post(s.Config.BaseURL+method, ApplicationJSON, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
