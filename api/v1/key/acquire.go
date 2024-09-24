package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

const method = "/openapi/v1/key/acquire"

func (a *Acquire) Handle(request *AcquireRequest) (*AcquireReply, error) {
	a.Request(method, "")
	return nil, nil
}
