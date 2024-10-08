package key

import "encoding/json"

type Status uint8

const (
	Normal Status = iota + 1
	DiscardIng
	Used
	Discard
)

var statusMap = map[Status]string{
	Normal:     "正常",
	DiscardIng: "作废中",
	Used:       "已核销",
	Discard:    "已作废",
}

func (s Status) Value() uint8 {
	return uint8(s)
}

func (s Status) GetText() string {
	tex, ok := statusMap[s]
	if !ok {
		return ""
	}
	return tex
}

func (s Status) IsNormal() bool {
	return s == Normal
}

func (s Status) IsUsed() bool {
	return s == Used
}

func (s Status) IsDiscard() bool {
	return s == Discard
}

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Data    *Reply `json:"data,omitempty"`
}

type Reply struct {
	OutBizNo       string `json:"out_biz_no"`
	TradeNo        string `json:"trade_no"`
	Key            string `json:"key"`
	Status         Status `json:"status"`
	Url            string `json:"url"`
	ValidBeginTime string `json:"valid_begin_time,omitempty"`
	ValidEndTime   string `json:"valid_end_time,omitempty"`
	UsageTime      string `json:"usage_time,omitempty"`
	DiscardTime    string `json:"discard_time,omitempty"`
}

func (a *Response) Response(b []byte) (*Response, error) {
	if err := json.Unmarshal(b, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Response) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(b)
}

func (a *Response) IsSuccess() bool {
	return a.Code == 200
}

func response(b []byte) (*Response, error) {
	var resp *Response
	return resp.Response(b)
}
