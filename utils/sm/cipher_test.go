package sm

import (
	"testing"
)

func TestCipher(t *testing.T) {

}

func TestPlain(t *testing.T) {
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
	cs, err := Cipher(puk, []byte(data))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cs)
	ps, err := Plain(prk, cs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ps)
}
