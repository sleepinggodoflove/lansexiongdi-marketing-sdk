package key

import (
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

const method = "/openapi/v1/key/order"

func (a *Acquire) Order(request *OrderRequest) (*Response, error) {
	b, err := a.Request(method, request)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
