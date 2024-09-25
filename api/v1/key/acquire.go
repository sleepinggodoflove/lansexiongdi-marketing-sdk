package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

const method = "/openapi/v1/key/acquire"

func (a *Acquire) Handle(request *AcquireRequest) (*Reply, error) {
	a.Request(method, nil)
	return nil, nil
}
