package key

import (
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/common"
)

var _ common.Request = (*Acquire)(nil)

type AcquireRequest struct {
	OutBizNo   string `json:"out_biz_no"`
	ActivityNo string `json:"activity_no"`
	Number     string `json:"number"`
}

func (a *Acquire) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Response struct {
	OutBizNo       string `json:"out_biz_no"`
	TradeNo        string `json:"trade_no"`
	Key            string `json:"key"`
	Status         string `json:"status"`
	Url            string `json:"url"`
	ValidBeginTime string `json:"valid_begin_time"`
	ValidEndTime   string `json:"valid_end_time"`
}

type AcquireReply struct {
	Code     string    `json:"code"`
	Msg      string    `json:"msg"`
	SubCode  string    `json:"subCode"`
	SubMsg   string    `json:"subMsg"`
	Response *Response `json:"response"`
	Sign     string    `json:"sign"`
}

type DiscardRequest struct {
	OutBizNo string `json:"out_biz_no"`
	TradeNo  string `json:"trade_no"`
	Key      string `json:"key"`
}

type DiscardReply struct {
	Code     string    `json:"code"`
	Msg      string    `json:"msg"`
	SubCode  string    `json:"subCode"`
	SubMsg   string    `json:"subMsg"`
	Response *Response `json:"response"`
	Sign     string    `json:"sign"`
}
