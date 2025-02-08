package key

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	appId      = "KY7gREqQ96Phxhcvj8es7B6e"
	privateKey = "MIIEowIBAAKCAQEA7a2I4l8OOdW4weVFvj4u/mBqP3aZhJ0mOTKl4MCW4Pf6gNAlZa5dZYOS/BocmG872+pd10BiI73qiAWsuVaPwCL0A37lQbCXlG0fDAfCLogXuF1qVNRZgkYKrx/5Gppo2PNed7E5YyCUkMUKVPbuwuZteMZJH8d1o6Uojbb/xJQvAGOlx5Y04VZWp/6p2GjhW0srwgbpVegMyyn2Qblx1Lo+Uq5zG8um7FTpbtb/L/itpBFEDSZGIIKDfn4FPyt+jQ0SW5TDYQClSvWHK4V3RkWOVkD1nHeBpZyp7JNehK+7kBfO6G4NJabkyoWqFEiZcTy38ZWQdqJ9N4LZuY37NwIDAQABAoIBAGs2u6e5z1YBda1pehN+Q36WCXeFTW0H4qUslq0S0zy6P/L5cdUzWYggWR6FvN56Vts2Foyxy1NqKTCgtrCIPqIiYkZtaIdAXLAkpTutCEgrNeABq6SGgbYFWG51Es6QVrl+1t9RP5zaponDiIyZM00R2tH/SB8gv41JREjhAvEuNIwPyaoVVt+U/kAdhJgiMKsDpoGaMfsJk76sORu6qQqBkBN8cglN94xC0QtROytW3EY8SnZmgGZHcY3YTXM74CWM8yBg7rNuKv0982f9hKvUDHKMFYly1PzYiSgplkT7RYCMjo2FFf1lt7k7N61+4nalS/EM6324m2poisTRFAkCgYEA+3oXvtwrTim0kXG7V7w5PS8u5dU0sAAH4ACSyzy8nEdKEk4ipoaTGm6km8ko0O+9E70SwZQgK/eAnmsYf5WtMzvItweKUUVsBCm01qSHlu5vzGO3H1ndi7hg+tH9VrOQH3+odQJP9FqC4BkncMszHM4nLglWSTixTTvGIovQLy0CgYEA8fPnu0tWqhfQV6svaA5kt4h6cL52ARKlubRuYkI4hGuikKYpd2A3WuVtD1LkuPQSjwID9730HAqLc7ZMwONjQ8NANi9ZoJR6A+Vzba9zDPQmSc80Ax7Kkjc03D1Y7yiP6P8bWnhCCbRcMy+dcobvBZc2zaWzSNjZwPOSV9xaMnMCgYAclz354hg+U7mGy7JsACdV0HZ5hOrvk6FRk18dIjOjZOuD90QzQJua5rdqSs2MK6WIh/eI8KlTtlj2KeDoKIE/kO15+a59HPJx6rf3q08LFuK5DyEzvEjW6MiF27f80n9xRVdGrlOeyWeVyOZWCZQvEzUbI86eloZ57HDTXqf1pQKBgF6T6xeJgZ0Hpgc/AU75oWEk1kfQC6yrr2CCKUv7esA4mtlUOo1RbRH48MK2snWh4sdIEGj9NbjoXk6jCim0OQ85+ZW0uKJOp8tyG8baeGyt23GqrzgxBxpUvjMBQAxsnKSFZBnfPGEywX+4syEbob9btq54gTaOncAQ9jmmBxQFAoGBAIpPbq2lYwOhgoUJ2BR34xjpmNOiOAF5AVLPGTH44a+iGMJ4tbF9AvfL4xsCWK9zMi3ExaKVN0lNn0cWx2lpXxwO6B+l4L//eczmHx4h1eLJd6ZWyTj7lq+RBOOUgHLKEssZfJ11RYTZjSD7s75JZteM2OFw7BVRRNgw387A3mj6"
	publicKey  = "MIIBCgKCAQEAla28cJMhC/Bum9n3Ukymhj9t4Lf4reNMNXOMQJBnQMbTgYFTP71hUDvwTmk+uDliItnqA0AT31brINJlt22lPaJRvrt7uIKLH605BY3Uk/I0jkJnWGhNPH53uAFN43ra1PBNY4oU2FBjYGySSzDUaSssHuuyor+4R5m969FTN5HRJpBBXJckpbnQGeFe7c/WyXXPk3a+J1I6Wu+oeK2+mrUU27e7GpTTy8vRlkTFt8IUJOSd4gK/xat22N17/Kv6U7L4iORecXJMo3H5ObBT6+OV2Qr7duVQLJwAgUdNOQR0DxAeEg0BcM/84bPoeZKf/MmuOws1MG5qkKGc86gpqQIDAQAB"
	key        = "b48a8d47ed8cec04f27133b391da9694"
	baseURL    = "https://market.api.86698.cn"
	signType   = core.SignRSA
)

func newCore() (*core.Core, error) {
	return core.NewCore(&core.Config{
		AppID:      appId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Key:        key,
		SignType:   signType,
		BaseURL:    baseURL,
	})
}

func TestBuildParams(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	req := &OrderRequest{
		OutBizNo:   "001",
		ActivityNo: "Ntest001",
		Number:     1,
		NotifyUrl:  "",
		Extra:      "",
	}
	p, err := c.BuildParams(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", p)
}

func TestOrder(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{c}
	_, r, err := a.Order(context.Background(), &OrderRequest{
		OutBizNo:   "b202412270z8q7r1f704",
		ActivityNo: "2024070901134",
		Number:     1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("response=%+v", r)
	if !r.IsSuccess() {
		t.Errorf("获取key失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("data=%+v", data)
}

func TestQuery(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{c}
	_, r, err := a.Query(context.Background(), &QueryRequest{
		OutBizNo: "006",
		TradeNo:  "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("response=%+v", r)
	if !r.IsSuccess() {
		t.Errorf("查询失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("data=%+v", data)
	//t.Log(result.Status.IsNormal())
	//t.Log(result.Status.IsUsed())
	//t.Log(result.Status.IsDiscardIng())
	//t.Log(result.Status.IsDiscard())
}

func TestDiscard(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{c}
	_, r, err := a.Discard(context.Background(), &DiscardRequest{
		OutBizNo: "20241211002",
		TradeNo:  "",
		Reason:   "正常作废",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("response=%+v", r)
	if !r.IsSuccess() {
		t.Errorf("作废收单失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("data=%+v", data)
	//assert.Equal(t, r.Data.Status, DiscardIng)
}

func TestNotify(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{c}
	r, err := a.Notify(context.Background(), &Notify{
		AppId:     "",
		SignType:  "",
		Timestamp: "",
		Sign:      "",
		Data: NotifyData{
			NotifyId:       "",
			OutBizNo:       "",
			TradeNo:        "",
			Key:            "",
			UsableNum:      0,
			UsageNum:       0,
			Status:         0,
			Url:            "",
			ValidBeginTime: "",
			ValidEndTime:   "",
			UsageTime:      "",
			DiscardTime:    "",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	assert.Equal(t, r.Status, Discard)
	//t.Log(r.Data.Status.IsNormal())
	//t.Log(r.Data.Status.IsUsed())
	//t.Log(r.Data.Status.IsDiscardIng())
	//t.Log(r.Data.Status.IsDiscard())
}

func TestCallback(t *testing.T) {
	data := NotifyData{
		NotifyId:       "7278418772598218752",
		OutBizNo:       "006",
		TradeNo:        "727291384764309505",
		Key:            "dpK5yorx2M2g2e0W",
		UsableNum:      1,
		UsageNum:       1,
		Status:         3,
		Url:            "https://market.86698.cn/dpK5yorx2M2g2e0W",
		ValidBeginTime: "2024-12-27 22:37:41",
		ValidEndTime:   "2024-12-27 18:08:25",
		UsageTime:      "2024-12-27 22:37:41",
		DiscardTime:    "",
	}
	n := &Notify{
		AppId:     "KY7gREqQ96Phxhcvj8es7B6e",
		SignType:  "RSA",
		Timestamp: time.Now().Format(time.DateTime),
		Sign:      "",
		Data:      data,
	}
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	signStr, err := c.CryptographySuite.Signer.Sign(n.SignString())
	if err != nil {
		t.Error(err)
		return
	}
	n.Sign = signStr
	a := &Key{c}
	r, err := a.Notify(context.Background(), n)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.Status.IsNormal())
}

func TestResponse(t *testing.T) {
	jsonBytes := []byte(`{"code":200,"data":{},"message":"成功"}`)
	resp, err := response(jsonBytes)
	if err != nil {
		t.Error(err)
		return
	}
	result, err := resp.ConvertData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", resp)
	t.Logf("%s", string(resp.Data))
	t.Logf("%+v", result)

	jsonBytes2 := []byte(`{"code":200,"message":"成功","data":{"out_biz_no":"123","trade_no":"456"}}`)
	resp2, err := response(jsonBytes2)
	if err != nil {
		t.Error(err)
		return
	}
	result2, err := resp2.ConvertData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", resp2)
	t.Logf("%s", string(resp2.Data))
	t.Logf("%+v", result2)
}
