package key

import (
	"context"
	core2 "github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"testing"
)

var (
	// 私钥
	rsaPrivateKey = "MIIEpQIBAAKCAQEA10X/4v44TU0sPOKhnHLx7Qe0SREVIne9+3DEtTCbzrewXSrlHedMT6RVlwv1i0D6q5zWGV9WXMI4ZrM2EM6wE4bZfjZU5FvzYTQN8fjiBQTggU8BDFpUPGibwt2Fyv6SAnj0AQzD7itvcxydu0tDOgg9Aaa701kxhz9yZukpZGVhAEBussRc41EuEwE2gQJ5prSFijzhqg5awwAJESX9gDHX3DncLO2s3FuaaReJ1U/c/LsJIvRGPckCllSJM1s9WAvjz5yFXZMpvWpmAnmpxd5fNd349Gr/swfM50TN5D3cTJblO771FifNFm6r+4o1nkR/thWX5vMasvKwimGo3QIDAQABAoIBAQCQkaXi3y8YWrdWvCwkUN0/fWkJmLtExn2Dmpu/wsEf9iQurVvo1SheY9JG+fUQa7bsAQuXRntNF/GgpsGsT+HXezwckog4Q7gSk066LZY8IKZUsKXXkeH4H5hbKUFsrcGIf4n+GoCKNglGmPUkjsq68kVmEn8Y1FF6rpU5n2P40xEKAieKxlM2JNwR22DQYRw3iw4PcMAD88nKx9OBUwGig8MQnUka7OCZk9fNLdwBT0VfgCzRdvyBCDieif4vB7TnMmvYlr6wWOMi2Ad9ccY2wTlVOUyHoC6BZ72FgOYyfnAmEZbDChCNTEDNNj0m056slCTMO+nIVUqMip4imgiBAoGBAOWI84+4dRA+gm6xynXsp9TAto43/DmbohHrUWRE2tSGqPevBz2i0c5AOaNUdxtzoEWOj260Zdzf3gRhv/iH3Mgp2+R2+cJ1QYoaX/2auJ50dPJf1SHZrKVIYhYqqlIWc8jQo8XBK/Ys/Lf+N2We2EsMLMtUUMH9OZ+20XtlxQBZAoGBAPAYGVOzJnqFuSqxFzAA7VXP6p9WKxGEzbDUCxFaLj751DnvogI8FczAE6ADBxNYVkeQYtvqzMb9nTollOL6+/T9MUJn3DXTT8/St+REVdWvmO9Az9nGVDkLGjz0CH0iMjN49iSMFsdmo6L2528kUOj20dPh5IkzyWykYqni+fwlAoGBAMa1A51U40rXwpTPp2TlJfnBh4ihIOJCQFDg9YonLYY0uUwKousR7C1wXjVuJtqGA6aTnsoIs/I9f3ctpEIkY9aInksvUFKurbk/0f+7FL5gNOmqWtk+Fv7TJc7oyp/bvgqHzG+jJkqscW9bTVvU4ow9kv3HFU6KyHriioEX/i6pAoGBAJs5yW4e3lrKh/u9ANPNVaRsRzF64V9zMBUKEpnGZy3aEcbfUiwFssZszINgUbvFGgsso22xcXGZ2IQWdhsFz84FwEpBodK+6tPfVXrkX2ZHICZXDcqrehpjPjR4ReC5MiGrK+BXHgcPKe6bmOd3YEQuB1zop/u4mpp98TgLAjptAoGAF1lgcmwnHDduvWErwSmMR3w45FppV0pOdSEmv/6DU6iglzleSuuCm4fhXAEgSahJMKiflNx1ufl84IGir0r/we8GNfs4ceSPlyKDYnaBG19f5jWLLNFTWfwVhZqa85pd9b2dcqVWMazHq8psmyt0opmctRcM9s4lNq2OmvB7vM4="
	// 公钥
	publicKeyStr = "MIIBCgKCAQEA10X/4v44TU0sPOKhnHLx7Qe0SREVIne9+3DEtTCbzrewXSrlHedMT6RVlwv1i0D6q5zWGV9WXMI4ZrM2EM6wE4bZfjZU5FvzYTQN8fjiBQTggU8BDFpUPGibwt2Fyv6SAnj0AQzD7itvcxydu0tDOgg9Aaa701kxhz9yZukpZGVhAEBussRc41EuEwE2gQJ5prSFijzhqg5awwAJESX9gDHX3DncLO2s3FuaaReJ1U/c/LsJIvRGPckCllSJM1s9WAvjz5yFXZMpvWpmAnmpxd5fNd349Gr/swfM50TN5D3cTJblO771FifNFm6r+4o1nkR/thWX5vMasvKwimGo3QIDAQAB"
	aesKey       = "870abfc720f86ce2c5e4d3345741d48d"
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
		OutBizNo:   "outBizNo",
		ActivityNo: "activityNo",
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
