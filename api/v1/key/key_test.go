package key

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"io"
	"strings"
	"testing"
)

var (
	rsaPrivateKey = "MIIEowIBAAKCAQEA7a2I4l8OOdW4weVFvj4u/mBqP3aZhJ0mOTKl4MCW4Pf6gNAlZa5dZYOS/BocmG872+pd10BiI73qiAWsuVaPwCL0A37lQbCXlG0fDAfCLogXuF1qVNRZgkYKrx/5Gppo2PNed7E5YyCUkMUKVPbuwuZteMZJH8d1o6Uojbb/xJQvAGOlx5Y04VZWp/6p2GjhW0srwgbpVegMyyn2Qblx1Lo+Uq5zG8um7FTpbtb/L/itpBFEDSZGIIKDfn4FPyt+jQ0SW5TDYQClSvWHK4V3RkWOVkD1nHeBpZyp7JNehK+7kBfO6G4NJabkyoWqFEiZcTy38ZWQdqJ9N4LZuY37NwIDAQABAoIBAGs2u6e5z1YBda1pehN+Q36WCXeFTW0H4qUslq0S0zy6P/L5cdUzWYggWR6FvN56Vts2Foyxy1NqKTCgtrCIPqIiYkZtaIdAXLAkpTutCEgrNeABq6SGgbYFWG51Es6QVrl+1t9RP5zaponDiIyZM00R2tH/SB8gv41JREjhAvEuNIwPyaoVVt+U/kAdhJgiMKsDpoGaMfsJk76sORu6qQqBkBN8cglN94xC0QtROytW3EY8SnZmgGZHcY3YTXM74CWM8yBg7rNuKv0982f9hKvUDHKMFYly1PzYiSgplkT7RYCMjo2FFf1lt7k7N61+4nalS/EM6324m2poisTRFAkCgYEA+3oXvtwrTim0kXG7V7w5PS8u5dU0sAAH4ACSyzy8nEdKEk4ipoaTGm6km8ko0O+9E70SwZQgK/eAnmsYf5WtMzvItweKUUVsBCm01qSHlu5vzGO3H1ndi7hg+tH9VrOQH3+odQJP9FqC4BkncMszHM4nLglWSTixTTvGIovQLy0CgYEA8fPnu0tWqhfQV6svaA5kt4h6cL52ARKlubRuYkI4hGuikKYpd2A3WuVtD1LkuPQSjwID9730HAqLc7ZMwONjQ8NANi9ZoJR6A+Vzba9zDPQmSc80Ax7Kkjc03D1Y7yiP6P8bWnhCCbRcMy+dcobvBZc2zaWzSNjZwPOSV9xaMnMCgYAclz354hg+U7mGy7JsACdV0HZ5hOrvk6FRk18dIjOjZOuD90QzQJua5rdqSs2MK6WIh/eI8KlTtlj2KeDoKIE/kO15+a59HPJx6rf3q08LFuK5DyEzvEjW6MiF27f80n9xRVdGrlOeyWeVyOZWCZQvEzUbI86eloZ57HDTXqf1pQKBgF6T6xeJgZ0Hpgc/AU75oWEk1kfQC6yrr2CCKUv7esA4mtlUOo1RbRH48MK2snWh4sdIEGj9NbjoXk6jCim0OQ85+ZW0uKJOp8tyG8baeGyt23GqrzgxBxpUvjMBQAxsnKSFZBnfPGEywX+4syEbob9btq54gTaOncAQ9jmmBxQFAoGBAIpPbq2lYwOhgoUJ2BR34xjpmNOiOAF5AVLPGTH44a+iGMJ4tbF9AvfL4xsCWK9zMi3ExaKVN0lNn0cWx2lpXxwO6B+l4L//eczmHx4h1eLJd6ZWyTj7lq+RBOOUgHLKEssZfJ11RYTZjSD7s75JZteM2OFw7BVRRNgw387A3mj6"
	publicKeyStr  = "MIIBCgKCAQEA43/kRptu3Y5i/LjZhGQMaExG7+VE4MvonWpXELhxFdLAsfLA+e1XcKBzD6uHWaKo+L7CWvSBtj3LXAr++uInDnxAiPbgsnmYe8tEeZzg2IYeYPThLH84XQouuTx1pspqvU2t7ZXQPJulq5LbYJUmRR+V4d8zKhSsctLlg3AhujHnZ4LUWJXjrnt++JYi3hFai+p1knUic5rXh/35HoYbnnljGhPPe2U3xn4TWvVkeAJQpxkdS4s2AuSUg+L0M5kkHgeYAr28YQTwOm8GcCaXuLUDdNdX1/Iud9tgsLzQVP3qVdxKVNp8vRZOnKn6OjwLaTBkePVuMVgmtBTR+Gsm+QIDAQAB"
	aesKey        = "ffb403db11face06d59526ac93f25789"
	//baseURL       = "https://gateway.dev.cdlsxd.cn"
	baseURL = "http://127.0.0.1:9000"
	appID   = "lzm"
)

func TestSignVerify(t *testing.T) {
	c := core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	}
	co, err := core.NewCore(&c, core.WithSignType(core.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	orderReq := &OrderRequest{
		OutBizNo:   "321312",
		ActivityNo: "lzm",
		Number:     1,
	}
	dataToStr, err := orderReq.String()
	if err != nil {
		t.Error(err)
		return
	}
	signature, err := co.Signer.Sign(dataToStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(signature)
	b := co.Verifier.Verify(dataToStr, signature)
	if !b {
		t.Error("签名验证失败")
	}
}

func TestGetParams(t *testing.T) {
	c := core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	}
	co, err := core.NewCore(&c, core.WithSignType(core.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	orderReq := &OrderRequest{
		OutBizNo:   "321312",
		ActivityNo: "lzm",
		Number:     1,
	}
	p, err := co.GetParams(orderReq)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", p)
}

func TestOrder(t *testing.T) {
	co, err := core.NewCore(&core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{co}
	r, err := a.Order(context.Background(), &OrderRequest{
		OutBizNo:   "321312",
		ActivityNo: "lzm",
		Number:     1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.IsSuccess())
}

func TestQuery(t *testing.T) {
	co, err := core.NewCore(&core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{co}
	r, err := a.Query(context.Background(), &QueryRequest{
		OutBizNo: "5555555",
		TradeNo:  "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	if r.IsSuccess() {
		t.Log(r.Data.Status.IsNormal())
		t.Log(r.Data.Status.IsUsed())
		t.Log(r.Data.Status.IsDiscardIng())
		t.Log(r.Data.Status.IsDiscard())
	}
}

func TestDiscard(t *testing.T) {
	co, err := core.NewCore(&core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{co}
	r, err := a.Discard(context.Background(), &DiscardRequest{
		OutBizNo: "outBizNo",
		TradeNo:  "tradeNo",
		Reason:   "正常作废",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	if r.IsSuccess() {
		t.Log("作废收单成功")
	}
	t.Log(r.Data.Status.IsDiscardIng())
}

func TestNotify(t *testing.T) {
	co, err := core.NewCore(&core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{co}
	r, err := a.Notify(context.Background(), &Notify{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.Status.IsUsed())
	t.Log(r.Status.IsDiscard())
}

func TestCallback(t *testing.T) {
	data := &NotifyData{
		NotifyId:       "123456",
		OutBizNo:       "123456",
		TradeNo:        "1234567",
		Key:            "xdwqdsd",
		Status:         1,
		Url:            "http://lsxd/xdwqdsd",
		ValidBeginTime: "2006-01-02 15:04:05",
		ValidEndTime:   "2006-01-02 15:04:07",
	}
	n := &Notify{
		AppId:     "123",
		SignType:  "RSA",
		Timestamp: "2006-01-02 15:04:05",
		Sign:      "",
		Data:      data,
	}
	co, err := core.NewCore(&core.Config{
		AppID:      n.AppId,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	signStr, err := co.Signer.Sign(n.SignString())
	if err != nil {
		t.Error(err)
		return
	}
	n.Sign = signStr
	a := &Key{co}
	r, err := a.Notify(context.Background(), n)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.Status.IsNormal())
}

func TestCallBackNotify(t *testing.T) {
	c := core.Config{
		AppID:      appID,
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    "http://127.0.0.1:8080/utils/v1/wechat/notify",
	}
	co, err := core.NewCore(&c, core.WithSignType(core.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := co.Post(context.Background(), c.BaseURL, []byte(`{}`))
	if err != nil {
		t.Error(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	bodyStr := string(body)
	t.Log(bodyStr)
	if strings.Contains(bodyStr, "ok") {
		t.Log("ok")
	} else {
		t.Error("error")
	}
}
