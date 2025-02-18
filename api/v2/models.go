package v2

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type OrderRequest struct {
	OutBizNo   string `validate:"required,min=2,max=32" json:"out_biz_no"` // 同一商户应用下不可重复
	ActivityNo string `validate:"required,min=2,max=32" json:"activity_no"`
	Number     int32  `validate:"required,min=1,max=10000" json:"number"`
	NotifyUrl  string `json:"notify_url,omitempty"` // 回调地址,为空则使用客户应用设置地址
	Extra      string `json:"extra,omitempty"`      // 拓展参数，备用
}

type OrderResponse struct {
	OutBizNo string `json:"out_biz_no"`
	TradeNo  string `json:"trade_no"`
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

type QueryResponse struct {
	OutBizNo   string `json:"out_biz_no"`
	TradeNo    string `json:"trade_no"`
	Status     Status `json:"status"`
	Ciphertext string `json:"ciphertext"`
}

type KeyInfo struct {
	Key            string `json:"key"`
	UsableNum      uint32 `json:"usable_num"`
	UsageNum       uint32 `json:"usage_num"`
	Url            string `json:"url"`
	ValidBeginTime string `json:"valid_begin_time,omitempty"`
	ValidEndTime   string `json:"valid_end_time,omitempty"`
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

type NotifyData struct {
	NotifyId       string `json:"notify_id" validate:"required,alphanum,min=2,max=32"`
	OutBizNo       string `json:"out_biz_no" validate:"required,alphanum,min=2,max=32"`
	TradeNo        string `json:"trade_no" validate:"required,alphanum,min=2,max=32"`
	Key            string `json:"key,omitempty"`
	UsableNum      uint32 `json:"usable_num"`
	UsageNum       uint32 `json:"usage_num"`
	Status         Status `json:"status" validate:"required"`
	Url            string `json:"url,omitempty"`
	ValidBeginTime string `json:"valid_begin_time,omitempty"`
	ValidEndTime   string `json:"valid_end_time,omitempty"`
	UsageTime      string `json:"usage_time,omitempty"`
	DiscardTime    string `json:"discard_time,omitempty"`
}
type Notify struct {
	AppId     string     `json:"app_id" validate:"required"`
	SignType  string     `json:"sign_type" validate:"required"`
	Timestamp string     `json:"timestamp" validate:"required"`
	Sign      string     `json:"sign" validate:"required"`
	Data      NotifyData `json:"data" validate:"required"`
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
