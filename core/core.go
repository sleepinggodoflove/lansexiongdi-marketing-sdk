package core

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/consts"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/interfaces"
	"net/http"
	"strings"
	"time"
)

type Params struct {
	AppId      string `json:"app_id"`
	SignType   string `json:"sign_type"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
	Ciphertext string `json:"ciphertext"`
}

type Response struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"subCode,omitempty"`
	SubMsg     string `json:"subMsg,omitempty"`
	Ciphertext string `json:"ciphertext,omitempty"`
}

type Config struct {
	AppID             string `validate:"required"`
	PrivateKey        string `validate:"required"`
	MerchantPublicKey string `validate:"required"`
	Key               string `validate:"required"`
	BaseURL           string `validate:"required"`
}

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
	if err := s.Validate(); err != nil {
		return nil, err
	}
	core := &Core{
		signType: consts.SignRSA,
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

func (c *Core) GetParams(request interfaces.Request) (string, error) {
	ciphertext, err := c.GetCiphertext(request)
	if err != nil {
		return "", err
	}
	timestamps := time.Now().Format(time.RFC3339)
	dataToSign := c.config.AppID + timestamps + ciphertext

	signature, err := c.Signer.Sign(dataToSign)
	if err != nil {
		return "", err
	}
	p := Params{
		AppId:      c.config.AppID,
		SignType:   c.signType,
		Timestamp:  timestamps,
		Sign:       signature,
		Ciphertext: ciphertext,
	}
	reqBody, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(reqBody), nil
}

// Request sends the request and Analysis the response
func (c *Core) Request(method string, request interfaces.Request) (*Response, error) {
	reqBody, err := c.GetParams(request)
	if err != nil {
		return nil, err
	}
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}
	resp, err := c.httpClient.Post(c.config.BaseURL+method, consts.ApplicationJSON, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
