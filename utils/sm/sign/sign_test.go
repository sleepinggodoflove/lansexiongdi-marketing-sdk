package sign

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sm2SecretKey(t *testing.T) {
	t.Run("Test_sm2SecretKey", func(t *testing.T) {
		priK, pubK := sm2SecretKey()
		fmt.Printf("%s%s\n", priK, pubK)
	})
}

// Test_getPukByPriK 根据私钥获取公钥
func Test_getPukByPriK(t *testing.T) {
	t.Run("Test_getPukByPriK", func(t *testing.T) {
		prk := "P1cF/bB94wYojLEOS2AAA1VUETtU69l4OMZ5GSkM7FY="
		puk := getPukByPrK(prk)
		fmt.Printf("prk: %s \n", prk)
		fmt.Printf("puk: %s \n", puk)
	})
}

func Test_sign(t *testing.T) {
	var signStr = "123456"
	var prk = "zJRUcwPpKFf4nWiN9wqSO9gpGFx5BP4WviqnPsrhkpc="
	var puk = "BKbxGVVlJGWK/ScU0ebKSe4Jr4LvcBGgvt/HHBk+ODVCYnJYvvmX8cDNpf3TVYuRdz/RUH6UDgcoVpz02jXNfrM="

	t.Run("Test_sign_verify", func(t *testing.T) {
		signRes := sign(signStr, prk)
		if assert.Empty(t, signRes) {
			t.Errorf("Test_sign 签名生成-失败")
			return
		}
		fmt.Printf("Test_sign=%s\n", signRes)
		b := verify(signStr, signRes, puk)
		if assert.True(t, b) {
			fmt.Printf("Test_sign 验签-成功: %t\n", b)
		} else {
			t.Errorf("Test_sign 验签-失败, err: %t\n", b)
		}
	})
}
