package core

import "github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/common"

type Auth interface {
	Sign(params *common.Params) (string, error)
	Verify(params *common.Params) bool
}
