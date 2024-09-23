package auth

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/common"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
)

var _ core.Auth = (*SM)(nil)

type SM struct {
}

func (S SM) Sign(params *common.Params) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (S SM) Verify(params *common.Params) bool {
	//TODO implement me
	panic("implement me")
}
