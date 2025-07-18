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
var _ core.Request = (*NotifyData)(nil)

type OrderRequest struct {
	OutBizNo   string `validate:"required,alphanum,min=2,max=32" json:"out_biz_no"` // 同一商户应用下不可重复
	ActivityNo string `validate:"required,min=2,max=32" json:"activity_no"`
	Number     int32  `validate:"required,eq=1" json:"number"` // v1只支持1，若要多个，请异步v2接口批量生产
	NotifyUrl  string `json:"notify_url,omitempty"`            // 回调地址,可为空
	Account    string `json:"account,omitempty"`               // 可兑换账号
	Extra      string `json:"extra,omitempty"`                 // 拓展参数，备用
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
	OutBizNo string `json:"out_biz_no,omitempty" validate:"omitempty,alphanum,min=2,max=32"` // out_biz_no/trade_no二选一 同一商户应用下不可重复
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
		return fmt.Errorf("参数错误,out_biz_no/trade_no 二选一")
	}

	if err := validator.New().Struct(q); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}

type DiscardRequest struct {
	OutBizNo string `json:"out_biz_no,omitempty" validate:"omitempty,alphanum,min=2,max=32"` // out_biz_no/trade_no二选一 同一商户应用下不可重复
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

type NotifyData struct {
	NotifyId        string  `json:"notify_id" validate:"required,alphanum,min=2,max=32"`
	OutBizNo        string  `json:"out_biz_no" validate:"required,alphanum,min=2,max=32"`
	TradeNo         string  `json:"trade_no" validate:"required,alphanum,min=2,max=32"`
	Key             string  `json:"key,omitempty"`
	UsableNum       uint32  `json:"usable_num"`
	UsageNum        uint32  `json:"usage_num"`
	Status          Status  `json:"status" validate:"required"`
	Url             string  `json:"url,omitempty"`
	Amount          float32 `json:"amount,omitempty"`
	PayAmount       float32 `json:"pay_amount,omitempty"`
	PayTime         string  `json:"pay_time,omitempty"`
	SettlementPrice float32 `json:"settlement_price,omitempty"`
	ValidBeginTime  string  `json:"valid_begin_time,omitempty"`
	ValidEndTime    string  `json:"valid_end_time,omitempty"`
	UsageTime       string  `json:"usage_time,omitempty"`
	DiscardTime     string  `json:"discard_time,omitempty"`
	Account         string  `json:"account,omitempty"` // 可兑换账号
}
type Notify struct {
	AppId     string     `json:"app_id" validate:"required"`
	SignType  string     `json:"sign_type" validate:"required"`
	Timestamp string     `json:"timestamp" validate:"required"`
	Sign      string     `json:"sign" validate:"required"`
	Data      NotifyData `json:"data"`
}

func (d *NotifyData) Validate() error {

	if err := validator.New().Struct(d); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}

func (a *NotifyData) String() (string, error) {

	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (d *Notify) Validate() error {

	if err := validator.New().Struct(d); err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}

func (a *Notify) String() string {

	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}

	return string(b)
}

func (a *Notify) SignString() string {

	b, err := json.Marshal(a.Data)
	if err != nil {
		return ""
	}

	return a.AppId + a.Timestamp + string(b)
}

type Data struct {
	OutBizNo        string  `json:"out_biz_no"`
	TradeNo         string  `json:"trade_no"`
	Key             string  `json:"key"`
	UsableNum       uint32  `json:"usable_num"`
	UsageNum        uint32  `json:"usage_num"`
	Status          Status  `json:"status"`
	Url             string  `json:"url"`
	Amount          float32 `json:"amount,omitempty"`
	PayAmount       float32 `json:"pay_amount,omitempty"`
	PayTime         string  `json:"pay_time,omitempty"`
	SettlementPrice float32 `json:"settlement_price,omitempty"`
	ValidBeginTime  string  `json:"valid_begin_time,omitempty"`
	ValidEndTime    string  `json:"valid_end_time,omitempty"`
	UsageTime       string  `json:"usage_time,omitempty"`
	DiscardTime     string  `json:"discard_time,omitempty"`
	Account         string  `json:"account,omitempty"` // 可兑换账号
}

func ConvertData(b []byte) (*Data, error) {

	var data *Data
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return data, nil
}
