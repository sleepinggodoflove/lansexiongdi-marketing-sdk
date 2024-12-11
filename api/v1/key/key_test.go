package key

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"io"
	"strings"
	"testing"
)

var (
	rsaPrivateKey = ""
	publicKeyStr  = ""
	aesKey        = ""
	baseURL       = "http://127.0.0.1:9000"
	appID         = "lzm"
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
		OutBizNo: "testOrder123456",
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
		BaseURL:    "http://127.0.0.1/notify",
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
