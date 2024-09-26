package key

import (
	"encoding/json"
	"fmt"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

const method = "/openapi/v1/key/order"

func (a *Acquire) Order(request *OrderRequest) (*Reply, error) {
	b, err := a.Request(method, request)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	var response Response
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	if response.Code != 200 {
		fmt.Println(response.Message)
		return nil, nil
	}
	return response.Data, nil
}
