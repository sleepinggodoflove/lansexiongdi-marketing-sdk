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
	OutBizNo string `json:"out_biz_no,omitempty" validate:"omitempty,alphanum,min=2,max=32"` // out_biz_no/trade_no二选一
	TradeNo  string `json:"trade_no,omitempty" validate:"omitempty,alphanum,min=2,max=32"`   // out_biz_no/trade_no二选一 若不为空，则优先使用
}

func (a *QueryRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (q *QueryRequest) Validate() error {
	if q.OutBizNo == "" && q.TradeNo == "" {
		return fmt.Errorf("out_biz_no/trade_no 二选一")
	}
	if err := validator.New().Struct(q); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}

type DiscardRequest struct {
	OutBizNo string `json:"out_biz_no,omitempty" validate:"omitempty,alphanum,min=2,max=32"` // out_biz_no/trade_no二选一
	TradeNo  string `json:"trade_no,omitempty" validate:"omitempty,alphanum,min=2,max=32"`   // out_biz_no/trade_no二选一 若不为空，则优先使用
	Reason   string `json:"reason,omitempty" validate:"omitempty,min=1,max=50"`              // 可为空
}

func (a *DiscardRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (d *DiscardRequest) Validate() error {
	if d.OutBizNo == "" && d.TradeNo == "" {
		return fmt.Errorf("out_biz_no/trade_no 二选一")
	}
	if err := validator.New().Struct(d); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}
	return nil
}
