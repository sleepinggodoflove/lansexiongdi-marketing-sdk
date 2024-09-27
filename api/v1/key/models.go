package key

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
)

var _ core.Request = (*OrderRequest)(nil)
var _ core.Request = (*DiscardRequest)(nil)
var _ core.Request = (*QueryRequest)(nil)

type OrderRequest struct {
	OutBizNo   string `validate:"required,min=2,max=30" json:"out_biz_no"`
	ActivityNo string `validate:"required,min=2,max=30" json:"activity_no"`
	Number     int32  `validate:"required,min=1" json:"number"`
}

func (a *OrderRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *OrderRequest) Validate() error {
	if err := validator.New().Struct(c); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

type QueryRequest struct {
	OutBizNo string `json:"out_biz_no"` // out_biz_no/trade_no二选一
	TradeNo  string `json:"trade_no"`   // // out_biz_no/trade_no二选一 若不为空，则优先使用
}

func (a *QueryRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *QueryRequest) Validate() error {
	if err := validator.New().Struct(c); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

type DiscardRequest struct {
	OutRequestNo string `validate:"required,min=2,max=30" json:"out_request_no"`
	OutBizNo     string `json:"out_biz_no,omitempty"` // out_biz_no/trade_no二选一
	TradeNo      string `json:"trade_no,omitempty"`   // out_biz_no/trade_no二选一 若不为空，则优先使用
}

func (a *DiscardRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *DiscardRequest) Validate() error {
	if err := validator.New().Struct(c); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	if c.OutBizNo == "" && c.TradeNo == "" {
		return fmt.Errorf("out_biz_no/trade_no 二选一")
	}
	return nil
}
