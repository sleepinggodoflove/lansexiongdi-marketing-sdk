package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

const Method = "/openapi/v1/key/acquire"

func (a *Acquire) Handle(request *AcquireRequest) (*AcquireReply, error) {
	a.Request(Method, "")
	a.Verifier.Verify("", "")
	return nil, nil
}
