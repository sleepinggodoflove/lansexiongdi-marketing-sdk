package core

import (
	"testing"
)

func TestRSASignVerify(t *testing.T) {
	c := Config{
		AppID:      "",
		PrivateKey: "",
		PublicKey:  "",
		Key:        "",
		BaseURL:    "http://127.0.0.1:9000",
	}
	core, err := NewCore(&c)
	if err != nil {
		t.Error(err)
		return
	}
	signStr := "123456{}测试"
	signature, err := core.Signer.Sign(signStr)
	if err != nil {
		t.Error(err)
		return
	}
	b := core.Verifier.Verify(signStr, signature)
	if !b {
		t.Error("验签失败")
	}
}

func TestSMSignVerify(t *testing.T) {
	c := Config{
		AppID:      "123456",
		PrivateKey: "zJRUcwPpKFf4nWiN9wqSO9gpGFx5BP4WviqnPsrhkpc=",
		PublicKey:  "BKbxGVVlJGWK/ScU0ebKSe4Jr4LvcBGgvt/HHBk+ODVCYnJYvvmX8cDNpf3TVYuRdz/RUH6UDgcoVpz02jXNfrM=",
		Key:        "t+VxHnp+K9huhtNT84Pk7A==",
		BaseURL:    "http://127.0.0.1:9000",
	}
	core, err := NewCore(&c, WithSignType(SignSM))
	if err != nil {
		t.Error(err)
		return
	}
	signStr := "123456{}测试"
	signature, err := core.Signer.Sign(signStr)
	if err != nil {
		t.Error(err)
		return
	}
	b := core.Verifier.Verify(signStr, signature)
	if !b {
		t.Error("验签失败")
	}
}
