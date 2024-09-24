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

// Core structure
type Core struct {
	Config     *Config
	httpClient *http.Client
	SignType   string
	Signer     Signer
	Verifier   Verifier
}

type Option func(*Core)

func WithSignType(signType string) Option {
	return func(s *Core) {
		s.SignType = signType
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(s *Core) {
		s.httpClient = client
	}
}

// NewCore creates a new Core instance
func NewCore(s *Config, o ...Option) (*Core, error) {
	if s == nil {
		return nil, errors.New("config is nil")
	}
	core := &Core{
		SignType:   SignRSA,
		Config:     s,
		httpClient: &http.Client{},
	}
	for _, f := range o {
		f(core)
	}
	factory := &SignerFactory{}
	signer, verifier, err := factory.SignerVerifier(core.SignType, s)
	if err != nil {
		return nil, err
	}
	core.Signer = signer
	core.Verifier = verifier
	return core, nil
}

// Request sends the request and verifies the response signature
func (c *Core) Request(method, ciphertext string) (*http.Response, error) {
	timestamps := time.Now().Format(time.RFC3339)
	dataToSign := c.Config.AppID + timestamps + ciphertext

	signature, err := c.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}
	reqData := map[string]string{
		"app_id":     c.Config.AppID,
		"sign_type":  c.SignType,
		"timestamp":  timestamps,
		"ciphertext": ciphertext,
		"sign":       signature,
	}
	reqBody, _ := json.Marshal(reqData)

	resp, err := c.httpClient.Post(c.Config.BaseURL+method, ApplicationJSON, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
