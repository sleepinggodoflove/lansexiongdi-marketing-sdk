package core

import (
	"encoding/json"
)

const SuccessCode = 200

type Response struct {
	Code    int32           `json:"code"`
	Message string          `json:"message"`
	Reason  string          `json:"reason,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func BuildResponse(b []byte) (*Response, error) {

	var resp Response
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (a *Response) IsSuccess() bool {
	return a.Code == SuccessCode
}
