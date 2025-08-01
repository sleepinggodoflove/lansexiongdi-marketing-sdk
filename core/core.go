package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/utils"
	"io"
	"net/http"
	"time"
)

// defaultHeader
var defaultHeader = http.Header{
	"Content-Type": []string{"application/json"},
}

// Params request params
type Params struct {
	// AppId      app id
	AppId string `json:"app_id"`
	// SignType   sign type
	SignType SignType `json:"sign_type"`
	// Timestamp 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Timestamp string `json:"timestamp"`
	// Sign
	Sign string `json:"sign"`
	// Ciphertext
	Ciphertext string `json:"ciphertext"`
}

// Core structure
type Core struct {
	// Headers
	Headers http.Header
	// HttpClient http client
	HttpClient *http.Client
	// Config      config
	Config *Config
	// CryptographySuite
	CryptographySuite *CryptographySuite
}

type Option func(*Core)

// WithHeaders sets the http request headers
func WithHeaders(headers http.Header) Option {
	return func(s *Core) {
		s.Headers = headers
	}
}

// WithHttpClient sets the http client
func WithHttpClient(client *http.Client) Option {
	return func(s *Core) {
		s.HttpClient = client
	}
}

// NewCore creates a new Core instance
func NewCore(c *Config, o ...Option) (*Core, error) {

	if err := c.Validate(); err != nil {
		return nil, err
	}

	core := &Core{
		Headers:    defaultHeader,
		HttpClient: http.DefaultClient,
		Config:     c,
	}
	for _, f := range o {
		f(core)
	}

	core.Headers.Set("Version", "1.0")
	core.Headers.Set("Appid", core.Config.AppID)
	core.Headers.Set("Sign-Type", string(core.Config.SignType))

	crs, err := c.CryptographySuite()
	if err != nil {
		return nil, err
	}

	core.CryptographySuite = crs

	return core, nil
}

// GetBizSignStr gets the biz sign str Go 版本是 1.15 或更高版本
func (c *Core) GetBizSignStr(request any) (plaintext string, err error) {

	kvs := utils.SortStruct(request)

	kvm := make(map[string]any, len(kvs))
	var order []string

	for _, kv := range kvs {
		kvm[kv.Key] = kv.Value
		order = append(order, kv.Key)
	}

	orderedMap := make(map[string]any)
	for _, key := range order {
		orderedMap[key] = kvm[key]
	}

	// 将 orderedMap 转换成 JSON 字符串，保持顺序
	kvmBytes, err := json.Marshal(orderedMap)
	if err != nil {
		return "", err
	}

	return string(kvmBytes), nil
}

// GetCiphertext gets the ciphertext
func (c *Core) GetCiphertext(request any) (string, error) {

	plaintext, err := c.GetBizSignStr(request)
	if err != nil {
		return "", err
	}

	ciphertext, err := c.CryptographySuite.Cipher.Encode(plaintext)
	if err != nil {
		return "", err
	}

	return ciphertext, nil
}

// BuildParams gets the params
func (c *Core) BuildParams(request Request) (*Params, error) {

	if err := request.Validate(); err != nil {
		return nil, err
	}

	ciphertext, err := c.GetCiphertext(request)
	if err != nil {
		return nil, err
	}

	timestamps := time.Now().Format(time.DateTime)
	dataToSign := c.Config.AppID + timestamps + ciphertext

	signature, err := c.CryptographySuite.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}

	c.Headers.Set("Timestamp", timestamps)
	c.Headers.Set("Sign", signature)

	return &Params{
		AppId:      c.Config.AppID,
		SignType:   c.Config.SignType,
		Timestamp:  timestamps,
		Sign:       signature,
		Ciphertext: ciphertext,
	}, nil
}

// BuildAnyApiParams gets the params
func (c *Core) BuildAnyApiParams(bizContent any) (*Params, error) {

	ciphertext, err := c.GetCiphertext(bizContent)
	if err != nil {
		return nil, err
	}

	timestamps := time.Now().Format(time.DateTime)
	dataToSign := c.Config.AppID + timestamps + ciphertext

	signature, err := c.CryptographySuite.Signer.Sign(dataToSign)
	if err != nil {
		return nil, err
	}

	c.Headers.Set("Timestamp", timestamps)
	c.Headers.Set("Sign", signature)

	return &Params{
		AppId:      c.Config.AppID,
		SignType:   c.Config.SignType,
		Timestamp:  timestamps,
		Sign:       signature,
		Ciphertext: ciphertext,
	}, nil
}

// Verify verifies the params
func (c *Core) Verify(timestamp, ciphertext, sign string) bool {
	dataToSign := c.Config.AppID + timestamp + ciphertext
	return c.CryptographySuite.Verifier.Verify(dataToSign, sign)
}

// GetRequestBody gets the request body
func (c *Core) GetRequestBody(_ context.Context, request Request) ([]byte, error) {

	reqBody, err := c.BuildParams(request)
	if err != nil {
		return nil, err
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	return reqBodyBytes, nil
}

// Post sends the request and Analysis the response
func (c *Core) Post(ctx context.Context, method string, request Request) (*http.Response, []byte, error) {

	reqBodyBytes, err := c.GetRequestBody(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	return c.Request(ctx, http.MethodPost, c.Config.BaseURL+method, reqBodyBytes)
}

// Request sends the request and Analysis the response
func (c *Core) Request(ctx context.Context, method, url string, body []byte) (*http.Response, []byte, error) {

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header = c.Headers

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("sending HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	return resp, bodyBytes, nil
}
