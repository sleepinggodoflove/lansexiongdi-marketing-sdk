package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

const method = "/openapi/v1/key/order"

func (a *Acquire) Order(request *OrderRequest) (*Reply, error) {
	a.Request(method, nil)
	return nil, nil
}
