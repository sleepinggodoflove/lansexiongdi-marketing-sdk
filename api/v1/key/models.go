package key

import (
	"encoding/json"
)

type AcquireRequest struct {
	OutBizNo   string `json:"out_biz_no"`
	ActivityNo string `json:"activity_no"`
	Number     int32  `json:"number"`
}

func (a *Acquire) String() (string, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type Reply struct {
	OutBizNo       string `json:"out_biz_no"`
	TradeNo        string `json:"trade_no"`
	Key            string `json:"key"`
	Status         string `json:"status"`
	Url            string `json:"url"`
	ValidBeginTime string `json:"valid_begin_time"`
	ValidEndTime   string `json:"valid_end_time"`
}

type DiscardRequest struct {
	OutBizNo string `json:"out_biz_no"`
	TradeNo  string `json:"trade_no"`
	Key      string `json:"key"`
}
