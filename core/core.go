package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"subCode"`
	SubMsg     string `json:"subMsg"`
	Ciphertext string `json:"ciphertext"`
}

type Config struct {
	AppID             string
	PrivateKey        string
	MerchantPublicKey string
	Key               string
	BaseURL           string
}

// Core structure
type Core struct {
	config       *Config
	httpClient   *http.Client
	signType     string
	Signer       Signer
	Verifier     Verifier
	EncodeDecode EncodeDecode
}

type Option func(*Core)

func WithSignType(signType string) Option {
	return func(s *Core) {
		s.signType = signType
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
		signType: SignRSA,
		config:   s,
	}
	for _, f := range o {
		f(core)
	}
	factory := &SignerFactory{}
	signer, verifier, encodeDecode, err := factory.SignerVerifier(core.signType, s)
	if err != nil {
		return nil, err
	}
	core.Signer = signer
	core.Verifier = verifier
	core.EncodeDecode = encodeDecode
	return core, nil
}

func (c *Core) GetCiphertext(request Request) (string, error) {
	plaintext, err := request.String()
	if err != nil {
		return "", err
	}
	ciphertext, err := c.EncodeDecode.Encode(plaintext)
	if err != nil {
		return "", err
	}
	return ciphertext, nil
}

// Request sends the request and verifies the response signature
func (c *Core) Request(method string, request Request) (*http.Response, error) {
	ciphertext, err := c.GetCiphertext(request)
	if err != nil {
		return nil, err
	}
	timestamps := time.Now().Format(time.RFC3339)
	dataToSign := c.config.AppID + timestamps + ciphertext

	signature, err := c.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}
	reqData := map[string]string{
		"app_id":     c.config.AppID,
		"sign_type":  c.signType,
		"timestamp":  timestamps,
		"ciphertext": ciphertext,
		"sign":       signature,
	}
	reqBody, _ := json.Marshal(reqData)

	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}
	resp, err := c.httpClient.Post(c.config.BaseURL+method, ApplicationJSON, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
