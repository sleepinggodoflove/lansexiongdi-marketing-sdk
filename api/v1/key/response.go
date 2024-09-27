package key

import "encoding/json"

type OrderStatus uint8

const (
	OrderNormal OrderStatus = iota + 1
	OrderUsed
	OrderDiscard
)

var OrderStatusMap = map[OrderStatus]string{
	OrderNormal:  "正常",
	OrderUsed:    "已核销",
	OrderDiscard: "已作废",
}

func (s OrderStatus) Value() uint8 {
	return uint8(s)
}

func (s OrderStatus) GetText() string {
	tex, ok := OrderStatusMap[s]
	if !ok {
		return ""
	}
	return tex
}

func (s OrderStatus) IsNormal() bool {
	return s == OrderNormal
}

func (s OrderStatus) IsUsed() bool {
	return s == OrderUsed
}

func (s OrderStatus) IsDiscard() bool {
	return s == OrderDiscard
}

type Response struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
	Data    *Reply `json:"data,omitempty"`
}

type Reply struct {
	OutBizNo       string      `json:"out_biz_no"`
	TradeNo        string      `json:"trade_no"`
	Key            string      `json:"key"`
	Status         OrderStatus `json:"status"`
	Url            string      `json:"url"`
	ValidBeginTime string      `json:"valid_begin_time"`
	ValidEndTime   string      `json:"valid_end_time"`
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

type Notify struct {
	AppId     string `json:"app_id"`
	SignType  string `json:"sign_type"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Data      *Reply `json:"data"`
}

func (a *Notify) String() string {
	b, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(b)
}

func (a *Notify) SignStr() string {
	b, err := json.Marshal(a.Data)
	if err != nil {
		return ""
	}
	return a.AppId + a.Timestamp + string(b)
}
