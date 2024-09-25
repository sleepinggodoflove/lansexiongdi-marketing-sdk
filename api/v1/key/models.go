package key

import (
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
)

var _ core.Request = (*AcquireRequest)(nil)
var _ core.Request = (*DiscardRequest)(nil)
var _ core.Request = (*DQueryRequest)(nil)

type Reply struct {
	OutBizNo       string `json:"out_biz_no"`
	TradeNo        string `json:"trade_no"`
	Key            string `json:"key,omitempty"`
	Status         string `json:"status"`
	Url            string `json:"url,omitempty"`
	ValidBeginTime string `json:"valid_begin_time,omitempty"`
	ValidEndTime   string `json:"valid_end_time,omitempty"`
}

type AcquireRequest struct {
	OutBizNo   string `json:"out_biz_no"`
	ActivityNo string `json:"activity_no"`
	Number     int32  `json:"number"`
}

func (a *AcquireRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type DQueryRequest struct {
	OutBizNo string `json:"out_biz_no"`
	TradeNo  string `json:"trade_no"`
	Key      string `json:"key"`
}

func (a *DQueryRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type DiscardRequest struct {
	OutBizNo string `json:"out_biz_no"`
	TradeNo  string `json:"trade_no"`
	Key      string `json:"key"`
}

func (a *DiscardRequest) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
