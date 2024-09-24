package core

import "testing"

func TestNewCore(t *testing.T) {
	c := Config{
		AppID:             "123456",
		PrivateKey:        "zJRUcwPpKFf4nWiN9wqSO9gpGFx5BP4WviqnPsrhkpc=",
		MerchantPublicKey: "BKbxGVVlJGWK/ScU0ebKSe4Jr4LvcBGgvt/HHBk+ODVCYnJYvvmX8cDNpf3TVYuRdz/RUH6UDgcoVpz02jXNfrM=",
		BaseURL:           "http://127.0.0.1:8007",
	}
	core, err := NewCore(&c, WithSignType(SignSM))
	if err != nil {
		t.Error(err)
		return
	}
	data := "123456{}测试"
	signature := "MEUCIHfYOk7yrhdqWgahCW4cKYLjyfxmiKyKR1IWRYxnayx7AiEAmuSgsY1BKytMSbcV/wlaEPEeBuBdyqlEVT6sHOesj6Q="
	b, err := core.Verifier.Verify(data, signature)
	if err != nil {
		t.Error(err)
		return
	}
	if !b {
		t.Fatal("verify failed")
	}
	t.Log(b)
}
