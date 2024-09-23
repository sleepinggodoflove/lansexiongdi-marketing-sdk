package auth

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/common"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
)

var _ core.Auth = (*RSA)(nil)

type RSA struct {
	PrivateKey        string // 私钥
	MerchantPublicKey string // 商户公钥
}

func (R *RSA) Sign(params *common.Params) (string, error) {
	params.SignString()
	//TODO implement me
	panic("implement me")
}

func (R *RSA) Verify(params *common.Params) bool {
	params.SignString()
	//TODO implement me
	panic("implement me")
}
