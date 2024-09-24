package sm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sign(t *testing.T) {
	data := "123456{}测试"
	prkStr := "zJRUcwPpKFf4nWiN9wqSO9gpGFx5BP4WviqnPsrhkpc="
	prk, err := PrivateKeySM(prkStr)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := Sign(data, prk)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature=%s\n", signature)
}

func Test_Verify(t *testing.T) {
	data := "123456{}测试"
	prkStr := "zJRUcwPpKFf4nWiN9wqSO9gpGFx5BP4WviqnPsrhkpc="
	pukStr := "BKbxGVVlJGWK/ScU0ebKSe4Jr4LvcBGgvt/HHBk+ODVCYnJYvvmX8cDNpf3TVYuRdz/RUH6UDgcoVpz02jXNfrM="
	prk, err := PrivateKeySM(prkStr)
	if err != nil {
		t.Fatal(err)
	}
	puk, err := PublicKeySM(pukStr)
	if err != nil {
		t.Fatal(err)
	}
	signature, err := Sign(data, prk)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("signature=%s\n", signature)

	b, err := Verify(data, signature, puk)
	if err != nil {
		t.Fatal(err)
	}
	if assert.True(t, b) {
		t.Logf("Test_sign 验签-成功 %t\n", b)
	} else {
		t.Errorf("Test_sign 验签-失败 %t\n", b)
	}
}
