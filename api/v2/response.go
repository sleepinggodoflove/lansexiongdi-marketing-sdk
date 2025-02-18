package v2

import (
	"encoding/json"
)

type Status uint8

const SuccessCode = 200

func (s Status) Value() uint8 {
	return uint8(s)
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

func (a *Response) IsSuccess() bool {
	return a.Code == SuccessCode
}
