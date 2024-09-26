package core

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/consts"
	"testing"
)

func TestNewCore(t *testing.T) {
	c := Config{
		AppID:             "123456",
		PrivateKey:        "zJRUcwPpKFf4nWiN9wqSO9gpGFx5BP4WviqnPsrhkpc=",
		MerchantPublicKey: "BKbxGVVlJGWK/ScU0ebKSe4Jr4LvcBGgvt/HHBk+ODVCYnJYvvmX8cDNpf3TVYuRdz/RUH6UDgcoVpz02jXNfrM=",
		Key:               "t+VxHnp+K9huhtNT84Pk7A==",
		BaseURL:           "http://127.0.0.1:9000",
	}
	core, err := NewCore(&c, WithSignType(consts.SignSM))
	if err != nil {
		t.Error(err)
		return
	}
	data := "123456{}测试"
	signature := "MEUCIHfYOk7yrhdqWgahCW4cKYLjyfxmiKyKR1IWRYxnayx7AiEAmuSgsY1BKytMSbcV/wlaEPEeBuBdyqlEVT6sHOesj6Q="
	b := core.Verifier.Verify(data, signature)
	if !b {
		t.Error("验签失败")
	}
}
