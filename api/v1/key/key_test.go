package key

import (
	"context"
	core2 "github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"io/ioutil"
	"strings"
	"testing"
)

var (
	// 私钥
	rsaPrivateKey = "MIIEpAIBAAKCAQEAsDOhgtGNgKMH01Tc31O6xfqLWn0MJPBJtVBiHBI2J2WTQ1cmvrXmkO5ZNQr3fg/nmGdvEaKZScxZkcbEv5LBgjcDIE5a18EJVt5wrYM23m94welCdKoe1bbRf14WKsX+oqj3NQlWv4EQjhQGwwxLGXhtQg7o9j3gF0ybL9CxsYFSzAUgPy9TJvlOz/AeeMGCORC7dgYZZhckMRIql9gJSVJoXnNxF/4XnZ2it1m8+qQEngfxT+3VIObev3gdUx7WXx3Uxzlu7MSGWKN0nlGE4x0s8WhPhCyInYZfMJVaXFOazAoOi8/fHoYzpgGeiQ3d8U9sbWlyoki+krqMj8iFbwIDAQABAoIBAGA2AQXWeIZ5/sblOqlzJbP+x2LEjwIIdqbbWobrZsiCTTPi/ZP67QfMLcep0lPySUpNiDc/6qWCQJI5z6qvbpw0f69/OVk/3WKimTIVSLuScISpYpEjZyzY43HBpSb777tPuZQrIkP1LF34D20nZEZnHHmKfKggRyFRwhcMxEogZDby3p+/UEodFG+oxsIaUmjwEaLZYHfX5JX706yptOKXq2yKmSB82yhnWYXvvJnwXer2lhMIs7gCOMWX6sLTNxPM+m1K+Qoa3L0bEwpfl/zvaWsRLyG8DZp7/rQ/1VXUOmKmRNk3nYGhIHoG/tiwAdxEKNkyy3HHkDYyEdE2gdkCgYEA3bQhFXAKLGBfxNqeETQY2dJ+6fB1OzqC5DQDZ2XfXAccMkcAeqZpm8+2gHQ+oqUb3nCZMcxZY1rQ+nTQOhaBb2FUXoHzlF6gWzFSdZvcV7FTKsIxo/rPX+xCucx23HyRhuFancPqXLjH5whZeM/gba7ne90/kJZs9YqMmiEfLpsCgYEAy3WKggQpovwZxHirqjll6oAP/I16WjsADEu34bBBXO0TwEuHYdCiFekOIUIwNcHQANbYnDvmaF5ZIdoJA0JaNtTlLnVdgH67jU/D+LYFh/y4FezzTuF0dOJT/5W7+Nf/1AMs5NAFHB4HXifsOzc6p6ZLfPKLsy9BctDOiUTUp70CgYAd5rh7krPcjlb8Ttv4yAT+A9TYKnU1OswCiT1YSbYZTGAyK79Vy3H88MPViVgVSZ17n8YZoE0CCHKdBPo8i2KDqiXd5Tr77NNF2V4cvJu4PBNSdgSJ8D7d2hiZxIjXbfmWrngxRQYr4UQc1dRzZd4IDQnte8ah8pfZeglClGsHNwKBgQDBELevap3MAGe/LTbQoAja/kytPd9lsej4wf0ql2Ne98UvSzybkCvOmMmEu9cdm91yfm0rzBd24Fi6K7kzs4oB0AHuFQb2AMyeapuu7aLQCPryK1gyePRWWdKVjJPDrcwgdibqtY6zwLcDHOUox97L6ZxmY443KVd0yTnDvmIBDQKBgQDS2CUsvyGXA6aTrHmanPVAQ20ZxIgXWHsoyM0DPAydtjcHoBR5g7EtsbEleo+UNZ1/YaWLUcJVocde1dfTVBiwc28Ln3WaGIVxmqLr8wPN2KiWxmG3gWNHVH/U40vUec8Pj5uKW0WNVu4pZX3EJIPmRSfnkMWrXmNhJzgrri5rHA=="
	// 公钥
	publicKeyStr = "MIIBCgKCAQEAsDOhgtGNgKMH01Tc31O6xfqLWn0MJPBJtVBiHBI2J2WTQ1cmvrXmkO5ZNQr3fg/nmGdvEaKZScxZkcbEv5LBgjcDIE5a18EJVt5wrYM23m94welCdKoe1bbRf14WKsX+oqj3NQlWv4EQjhQGwwxLGXhtQg7o9j3gF0ybL9CxsYFSzAUgPy9TJvlOz/AeeMGCORC7dgYZZhckMRIql9gJSVJoXnNxF/4XnZ2it1m8+qQEngfxT+3VIObev3gdUx7WXx3Uxzlu7MSGWKN0nlGE4x0s8WhPhCyInYZfMJVaXFOazAoOi8/fHoYzpgGeiQ3d8U9sbWlyoki+krqMj8iFbwIDAQAB"
	aesKey       = "381a72d6311974f32aee169492f550eb"
	baseURL      = "http://127.0.0.1:9000"
)

func TestSignVerify(t *testing.T) {
	c := core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	}
	core, err := core2.NewCore(&c, core2.WithSignType(core2.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	orderReq := &OrderRequest{
		OutBizNo:   "test1",
		ActivityNo: "openapi1",
		Number:     1,
	}
	dataToStr, err := orderReq.String()
	if err != nil {
		t.Error(err)
		return
	}
	signature, err := core.Signer.Sign(dataToStr)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(signature)
	b := core.Verifier.Verify(dataToStr, signature)
	if !b {
		t.Error("签名验证失败")
	}
}

func TestGetParams(t *testing.T) {
	c := core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	}
	core, err := core2.NewCore(&c, core2.WithSignType(core2.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	orderReq := &OrderRequest{
		OutBizNo:   "outBizNo",
		ActivityNo: "activityNo",
		Number:     1,
	}
	p, err := core.GetParams(orderReq)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", p)
}

func TestOrder(t *testing.T) {
	core, err := core2.NewCore(&core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{core}
	r, err := a.Order(context.Background(), &OrderRequest{
		OutBizNo:   "outBizNo",
		ActivityNo: "activityNo",
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
	core, err := core2.NewCore(&core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{core}
	r, err := a.Query(context.Background(), &QueryRequest{
		OutBizNo: "outBizNo",
		TradeNo:  "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.IsSuccess())
}

func TestDiscard(t *testing.T) {
	core, err := core2.NewCore(&core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{core}
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
}

func TestNotify(t *testing.T) {
	core, err := core2.NewCore(&core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    baseURL,
	})
	if err != nil {
		t.Error(err)
		return
	}
	a := &Key{core}
	r, err := a.Notify(context.Background(), &Notify{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.Status.IsNormal())
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
	core, err := core2.NewCore(&core2.Config{
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
	signStr, err := core.Signer.Sign(n.SignString())
	if err != nil {
		t.Error(err)
		return
	}
	n.Sign = signStr
	a := &Key{core}
	r, err := a.Notify(context.Background(), n)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
	t.Log(r.Status.IsNormal())
}

func TestCallBackNotify(t *testing.T) {
	c := core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey,
		PublicKey:  publicKeyStr,
		Key:        aesKey,
		BaseURL:    "http://127.0.0.1:8080/utils/v1/wechat/notify",
	}
	core, err := core2.NewCore(&c, core2.WithSignType(core2.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	resp, err := core.Post(context.Background(), c.BaseURL, []byte(`{}`))
	if err != nil {
		t.Error(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
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
