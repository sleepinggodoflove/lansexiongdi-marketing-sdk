package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/common"
)

type Acquire api.Service

func (a *Acquire) Method() string {
	return "/openapi/v1/key/acquire"
}

func (a *Acquire) Method2() string {
	return "openapi.v1.key.acquire"
}

func (a *Acquire) Handle(request *AcquireRequest) (*AcquireReply, error) {
	a.Params = common.NewParams(a.Method(), nil)
	common.NewParams("", nil)
	return nil, nil
}
