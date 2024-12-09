package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strings"
	"time"
)

// Params request params
type Params struct {
	AppId      string `json:"app_id"`
	SignType   string `json:"sign_type"`
	Timestamp  string `json:"timestamp"` // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Sign       string `json:"sign"`
	Ciphertext string `json:"ciphertext"`
}

// Config merchant app Config
type Config struct {
	AppID      string `validate:"required"`
	PrivateKey string `validate:"required"`
	PublicKey  string `validate:"required"`
	Key        string `validate:"required"`
	BaseURL    string `validate:"required,url"`
}

// Validate Config
func (c *Config) Validate() error {
	if err := validator.New().Struct(c); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}

// Core structure
type Core struct {
	HttpClient *http.Client
	Config     *Config

	EncodeDecode EncodeDecode

	SignType string
	Signer   Signer
	Verifier Verifier
}

type Option func(*Core)

// WithSignType sets the sign type
func WithSignType(signType string) Option {
	return func(s *Core) {
		s.SignType = signType
	}
}

// WithHttpClient sets the http client
func WithHttpClient(client *http.Client) Option {
	return func(s *Core) {
		s.HttpClient = client
	}
}

// NewCore creates a new Core instance
func NewCore(conf *Config, o ...Option) (*Core, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	core := &Core{
		HttpClient: http.DefaultClient,
		Config:     conf,
		SignType:   SignRSA,
	}
	for _, f := range o {
		f(core)
	}
	factory := &SignerFactory{}
	signer, verifier, encodeDecode, err := factory.SignerVerifier(core.SignType, conf)
	if err != nil {
		return nil, err
	}
	core.Signer = signer
	core.Verifier = verifier
	core.EncodeDecode = encodeDecode
	return core, nil
}

// GetCiphertext gets the ciphertext
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

// GetParams gets the params
func (c *Core) GetParams(request Request) (*Params, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	ciphertext, err := c.GetCiphertext(request)
	if err != nil {
		return nil, err
	}
	timestamps := time.Now().Format(time.DateTime)
	dataToSign := c.Config.AppID + timestamps + ciphertext

	signature, err := c.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}
	return &Params{
		AppId:      c.Config.AppID,
		SignType:   c.SignType,
		Timestamp:  timestamps,
		Sign:       signature,
		Ciphertext: ciphertext,
	}, nil
}

// Verify verifies the params
func (c *Core) Verify(params *Params) bool {
	dataToSign := c.Config.AppID + params.Timestamp + params.Ciphertext
	return c.Verifier.Verify(dataToSign, params.Sign)
}

// Request sends the request and Analysis the response
func (c *Core) Request(ctx context.Context, method string, request Request) ([]byte, error) {
	reqBody, err := c.GetParams(request)
	if err != nil {
		return nil, err
	}
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.Post(ctx, c.Config.BaseURL+method, reqBodyBytes)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// Post sends the request and Analysis the response
func (c *Core) Post(_ context.Context, url string, reqBodyBytes []byte) (*http.Response, error) {
	return c.HttpClient.Post(url, ApplicationJSON, strings.NewReader(string(reqBodyBytes)))
}
