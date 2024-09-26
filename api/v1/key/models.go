package key

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/interfaces"
)

var _ interfaces.Request = (*OrderRequest)(nil)
var _ interfaces.Request = (*DiscardRequest)(nil)
var _ interfaces.Request = (*QueryRequest)(nil)

var _ interfaces.Validate = (*OrderRequest)(nil)
var _ interfaces.Validate = (*QueryRequest)(nil)
var _ interfaces.Validate = (*DiscardRequest)(nil)

type OrderRequest struct {
	OutBizNo   string `validate:"required" json:"out_biz_no"`
	ActivityNo string `validate:"required" json:"activity_no"`
	Number     int32  `validate:"required" json:"number"`
}

func (a *OrderRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *OrderRequest) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

type QueryRequest struct {
	OutBizNo string `validate:"required" json:"out_biz_no"`
	TradeNo  string `validate:"required" json:"trade_no"`
	Key      string `validate:"required" json:"key"`
}

func (a *QueryRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *QueryRequest) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

type DiscardRequest struct {
	OutBizNo string `validate:"required" json:"out_biz_no"`
	TradeNo  string `validate:"required" json:"trade_no"`
	Key      string `validate:"required" json:"key"`
}

func (a *DiscardRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *DiscardRequest) Validate() error {
	err := validator.New().Struct(c)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}
