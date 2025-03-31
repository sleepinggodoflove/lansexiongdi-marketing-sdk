package anyapi

import (
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"testing"
)

var (
	appId      = "OP001"
	privateKey = "MIIEowIBAAKCAQEA7a2I4l8OOdW4weVFvj4u/mBqP3aZhJ0mOTKl4MCW4Pf6gNAlZa5dZYOS/BocmG872+pd10BiI73qiAWsuVaPwCL0A37lQbCXlG0fDAfCLogXuF1qVNRZgkYKrx/5Gppo2PNed7E5YyCUkMUKVPbuwuZteMZJH8d1o6Uojbb/xJQvAGOlx5Y04VZWp/6p2GjhW0srwgbpVegMyyn2Qblx1Lo+Uq5zG8um7FTpbtb/L/itpBFEDSZGIIKDfn4FPyt+jQ0SW5TDYQClSvWHK4V3RkWOVkD1nHeBpZyp7JNehK+7kBfO6G4NJabkyoWqFEiZcTy38ZWQdqJ9N4LZuY37NwIDAQABAoIBAGs2u6e5z1YBda1pehN+Q36WCXeFTW0H4qUslq0S0zy6P/L5cdUzWYggWR6FvN56Vts2Foyxy1NqKTCgtrCIPqIiYkZtaIdAXLAkpTutCEgrNeABq6SGgbYFWG51Es6QVrl+1t9RP5zaponDiIyZM00R2tH/SB8gv41JREjhAvEuNIwPyaoVVt+U/kAdhJgiMKsDpoGaMfsJk76sORu6qQqBkBN8cglN94xC0QtROytW3EY8SnZmgGZHcY3YTXM74CWM8yBg7rNuKv0982f9hKvUDHKMFYly1PzYiSgplkT7RYCMjo2FFf1lt7k7N61+4nalS/EM6324m2poisTRFAkCgYEA+3oXvtwrTim0kXG7V7w5PS8u5dU0sAAH4ACSyzy8nEdKEk4ipoaTGm6km8ko0O+9E70SwZQgK/eAnmsYf5WtMzvItweKUUVsBCm01qSHlu5vzGO3H1ndi7hg+tH9VrOQH3+odQJP9FqC4BkncMszHM4nLglWSTixTTvGIovQLy0CgYEA8fPnu0tWqhfQV6svaA5kt4h6cL52ARKlubRuYkI4hGuikKYpd2A3WuVtD1LkuPQSjwID9730HAqLc7ZMwONjQ8NANi9ZoJR6A+Vzba9zDPQmSc80Ax7Kkjc03D1Y7yiP6P8bWnhCCbRcMy+dcobvBZc2zaWzSNjZwPOSV9xaMnMCgYAclz354hg+U7mGy7JsACdV0HZ5hOrvk6FRk18dIjOjZOuD90QzQJua5rdqSs2MK6WIh/eI8KlTtlj2KeDoKIE/kO15+a59HPJx6rf3q08LFuK5DyEzvEjW6MiF27f80n9xRVdGrlOeyWeVyOZWCZQvEzUbI86eloZ57HDTXqf1pQKBgF6T6xeJgZ0Hpgc/AU75oWEk1kfQC6yrr2CCKUv7esA4mtlUOo1RbRH48MK2snWh4sdIEGj9NbjoXk6jCim0OQ85+ZW0uKJOp8tyG8baeGyt23GqrzgxBxpUvjMBQAxsnKSFZBnfPGEywX+4syEbob9btq54gTaOncAQ9jmmBxQFAoGBAIpPbq2lYwOhgoUJ2BR34xjpmNOiOAF5AVLPGTH44a+iGMJ4tbF9AvfL4xsCWK9zMi3ExaKVN0lNn0cWx2lpXxwO6B+l4L//eczmHx4h1eLJd6ZWyTj7lq+RBOOUgHLKEssZfJ11RYTZjSD7s75JZteM2OFw7BVRRNgw387A3mj6"
	publicKey  = "MIIBCgKCAQEA6a394VZiQqXtpPQiuFizYpP9mvdThf6lUMeT7dxsrOw1RvU/3gBRObMbL39noo2ExrO97r0csO5sJ+4eSpiinGjnIoonbIl9mEn4vOrJymgUqYJh+SBw5Pge5bvitWCR94+kOh8zJITCcqYYqZz07sDRXGivtH2UuX7axmCgONpXii+vM9GkdNly3XZ+ipOhgvfXV/aSqn1z8c0cYp4QDbT0AI+tooAUTZ0RjDW07FHhTCH4IPJ+zzk3RbDHZ4r+bAlmJSOrwhCvogYTQAccMUpkJx8zONQiC2ZO/N2aHWLPyLu+RkREVBEwxOqWJ9oLdE0IAvEr5HEAPblQQv+TVwIDAQAB"
	key        = "d3626406d44ab4458a4028573f0d408f"
	baseURL    = "http://127.0.0.1:9000"
	signType   = core.SignRSA
)

func newCore() *core.Core {

	c, err := core.NewCore(&core.Config{
		AppID:      appId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Key:        key,
		SignType:   signType,
		BaseURL:    baseURL,
	})

	if err != nil {
		panic(err)
	}

	return c
}

func Test_AnyApi(t *testing.T) {

	c := newCore()

	bizContent := struct {
		Source       string `json:"source"`         // 来源
		AppId        string `json:"app_id"`         // 应用Id
		MchPublicKey string `json:"mch_public_key"` // 客户公钥
		NotifyUrl    string `json:"notify_url"`     // 事件通知地址,可为空
	}{
		Source:       "来源",
		AppId:        "OP002",
		MchPublicKey: "MIIBCgKCAQEA6a394VZiQqXtpPQiuFizYpP9mvdThf6lUMeT7dxsrOw1RvU/3gBRObMbL39noo2ExrO97r0csO5sJ+4eSpiinGjnIoonbIl9mEn4vOrJymgUqYJh+SBw5Pge5bvitWCR94+kOh8zJITCcqYYqZz07sDRXGivtH2UuX7axmCgONpXii+vM9GkdNly3XZ+ipOhgvfXV/aSqn1z8c0cYp4QDbT0AI+tooAUTZ0RjDW07FHhTCH4IPJ+zzk3RbDHZ4r+bAlmJSOrwhCvogYTQAccMUpkJx8zONQiC2ZO/N2aHWLPyLu+RkREVBEwxOqWJ9oLdE0IAvEr5HEAPblQQv+TVwIDAQAB",
		NotifyUrl:    "https://utils.85938.cn/utils/v1/wechat/notify",
	}

	a := &AnyApi{c}

	method := "/openapi/v1/merchant/appSet"

	_, r, err := a.AnyApi(context.Background(), method, bizContent)
	if err != nil {
		t.Error(err)
		return
	}

	if !r.IsSuccess() {
		t.Error(r.Message)
		return
	}

	t.Logf("data=%s", string(r.Data))

	var bizDataContent = struct {
		Ciphertext string `json:"ciphertext,omitempty"`
	}{}

	_ = json.Unmarshal(r.Data, &bizDataContent)

	bizJsonContent, _ := c.CryptographySuite.Cipher.Decode(bizDataContent.Ciphertext)

	t.Logf("bizJsonContent=%s", bizJsonContent)
}
