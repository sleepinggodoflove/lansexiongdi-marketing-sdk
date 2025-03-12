package key

import (
	"encoding/json"
)

type Data struct {
	OutBizNo       string  `json:"out_biz_no"`
	TradeNo        string  `json:"trade_no"`
	Key            string  `json:"key"`
	UsableNum      uint32  `json:"usable_num"`
	UsageNum       uint32  `json:"usage_num"`
	Status         Status  `json:"status"`
	Url            string  `json:"url"`
	Amount         float32 `json:"amount,omitempty"`
	PayAmount      float32 `json:"pay_amount,omitempty"`
	PayTime        string  `json:"pay_time,omitempty"`
	ValidBeginTime string  `json:"valid_begin_time,omitempty"`
	ValidEndTime   string  `json:"valid_end_time,omitempty"`
	UsageTime      string  `json:"usage_time,omitempty"`
	DiscardTime    string  `json:"discard_time,omitempty"`
}

type Response struct {
	Code    int32           `json:"code"`
	Message string          `json:"message"`
	Reason  string          `json:"reason,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func response(b []byte) (*Response, error) {
	var resp Response
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (a *Response) ConvertData() (*Data, error) {
	if a.Data == nil || len(a.Data) == 0 {
		return nil, nil
	}

	var reply Data
	if err := json.Unmarshal(a.Data, &reply); err != nil {
		return nil, err
	}

	if reply == (Data{}) {
		return nil, nil
	}

	return &reply, nil
}

func (a *Response) IsSuccess() bool {
	return a.Code == SuccessCode
}
