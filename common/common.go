package common

import (
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
)

type Request interface {
	String() (string, error)
}

type Params struct {
	AppId      string
	SignType   string
	Method     string
	Timestamp  string
	Ciphertext string
	Sign       string
	Auth       core.Auth
}

func NewParams(appId string, auth core.Auth) *Params {
	return &Params{
		AppId:      appId,
		SignType:   "RSA",
		Method:     "",
		Timestamp:  "",
		Ciphertext: "",
		Sign:       "",
		Auth:       auth,
	}
}

func (p *Params) SignString() string {
	return fmt.Sprintf("app_id=%s&sign_type=%s&method=%s&timestamp=%s&cipher_text=%s",
		p.AppId,
		p.SignType,
		p.Method,
		p.Timestamp,
		p.Ciphertext)
}
