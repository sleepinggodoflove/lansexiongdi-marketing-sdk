package sm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateKey(t *testing.T) {
	priK, pubK, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("\n私钥=%s\n公钥=%s\n", priK, pubK)
}

// Test_GetPukByPrK 根据私钥获取公钥
func Test_GetPukByPriK(t *testing.T) {
	prk := "ZvMCTiG67qUyPN65fcgg+EHhy2W/fN+9ixBudcmfbuU="
	puk, err := GetPukByPrK(prk)
	if err != nil {
		t.Fatal(err)
	}
	want := "BL6PopnKr/hhPkxgn700Li1hPGx2/J5y2dQ4BDPLKDXe1sS4JeIG8/W1B8AO7hBzi0bKTArti0E/HJJcR9WcH/I="
	fmt.Printf("prk: %s \n", prk)
	fmt.Printf("puk: %s \n", puk)
	assert.Equal(t, puk, want)
}
