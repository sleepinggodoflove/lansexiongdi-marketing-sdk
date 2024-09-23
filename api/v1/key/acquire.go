package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/api"
)

type Acquire api.Service

func (a *Acquire) Handle(request *AcquireRequest) (*AcquireReply, error) {
	a.Request("", "")
	return nil, nil
}
