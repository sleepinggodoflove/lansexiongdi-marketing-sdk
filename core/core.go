package core

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/interfaces"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Params request params
type Params struct {
	AppId      string `json:"app_id"`
	SignType   string `json:"sign_type"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
	Ciphertext string `json:"ciphertext"`
}

// Config merchant app config
type Config struct {
	AppID      string `validate:"required"`
	PrivateKey string `validate:"required"`
	PublicKey  string `validate:"required"`
	Key        string `validate:"required"`
	BaseURL    string `validate:"required"`
}

// Validate config
func (c *Config) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}

// Core structure
type Core struct {
	config       *Config
	httpClient   *http.Client
	signType     string
	Signer       interfaces.Signer
	Verifier     interfaces.Verifier
	EncodeDecode interfaces.EncodeDecode
}

type Option func(*Core)

// WithSignType sets the sign type
func WithSignType(signType string) Option {
	return func(s *Core) {
		s.signType = signType
	}
}

// WithHttpClient sets the http client
func WithHttpClient(client *http.Client) Option {
	return func(s *Core) {
		s.httpClient = client
	}
}

// NewCore creates a new Core instance
func NewCore(s *Config, o ...Option) (*Core, error) {
	if err := s.Validate(); err != nil {
		return nil, err
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

// GetCiphertext gets the ciphertext
func (c *Core) GetCiphertext(request interfaces.Request) (string, error) {
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

// GetParams gets the params
func (c *Core) GetParams(request interfaces.Request) (*Params, error) {
	ciphertext, err := c.GetCiphertext(request)
	if err != nil {
		return nil, err
	}
	timestamps := time.Now().Format("2006-01-02 15:04:05")
	dataToSign := c.config.AppID + timestamps + ciphertext

	signature, err := c.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}
	return &Params{
		AppId:      c.config.AppID,
		SignType:   c.signType,
		Timestamp:  timestamps,
		Sign:       signature,
		Ciphertext: ciphertext,
	}, nil
}

// Verify verifies the params
func (c *Core) Verify(params *Params) bool {
	dataToSign := c.config.AppID + params.Timestamp + params.Ciphertext
	b := c.Verifier.Verify(dataToSign, params.Sign)
	if b {
		return true
	}
	return false
}

// Request sends the request and Analysis the response
func (c *Core) Request(method string, request interfaces.Request) ([]byte, error) {
	reqBody, err := c.GetParams(request)
	if err != nil {
		return nil, err
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}
	resp, err := c.httpClient.Post(c.config.BaseURL+method, ApplicationJSON, strings.NewReader(string(reqBodyBytes)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
