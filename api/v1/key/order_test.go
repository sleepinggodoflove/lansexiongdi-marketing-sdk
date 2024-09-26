package key

import (
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/consts"
	core2 "github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"testing"
)

var rsaPrivateKey = "MIIEpQIBAAKCAQEA10X/4v44TU0sPOKhnHLx7Qe0SREVIne9+3DEtTCbzrewXSrlHedMT6RVlwv1i0D6q5zWGV9WXMI4ZrM2EM6wE4bZfjZU5FvzYTQN8fjiBQTggU8BDFpUPGibwt2Fyv6SAnj0AQzD7itvcxydu0tDOgg9Aaa701kxhz9yZukpZGVhAEBussRc41EuEwE2gQJ5prSFijzhqg5awwAJESX9gDHX3DncLO2s3FuaaReJ1U/c/LsJIvRGPckCllSJM1s9WAvjz5yFXZMpvWpmAnmpxd5fNd349Gr/swfM50TN5D3cTJblO771FifNFm6r+4o1nkR/thWX5vMasvKwimGo3QIDAQABAoIBAQCQkaXi3y8YWrdWvCwkUN0/fWkJmLtExn2Dmpu/wsEf9iQurVvo1SheY9JG+fUQa7bsAQuXRntNF/GgpsGsT+HXezwckog4Q7gSk066LZY8IKZUsKXXkeH4H5hbKUFsrcGIf4n+GoCKNglGmPUkjsq68kVmEn8Y1FF6rpU5n2P40xEKAieKxlM2JNwR22DQYRw3iw4PcMAD88nKx9OBUwGig8MQnUka7OCZk9fNLdwBT0VfgCzRdvyBCDieif4vB7TnMmvYlr6wWOMi2Ad9ccY2wTlVOUyHoC6BZ72FgOYyfnAmEZbDChCNTEDNNj0m056slCTMO+nIVUqMip4imgiBAoGBAOWI84+4dRA+gm6xynXsp9TAto43/DmbohHrUWRE2tSGqPevBz2i0c5AOaNUdxtzoEWOj260Zdzf3gRhv/iH3Mgp2+R2+cJ1QYoaX/2auJ50dPJf1SHZrKVIYhYqqlIWc8jQo8XBK/Ys/Lf+N2We2EsMLMtUUMH9OZ+20XtlxQBZAoGBAPAYGVOzJnqFuSqxFzAA7VXP6p9WKxGEzbDUCxFaLj751DnvogI8FczAE6ADBxNYVkeQYtvqzMb9nTollOL6+/T9MUJn3DXTT8/St+REVdWvmO9Az9nGVDkLGjz0CH0iMjN49iSMFsdmo6L2528kUOj20dPh5IkzyWykYqni+fwlAoGBAMa1A51U40rXwpTPp2TlJfnBh4ihIOJCQFDg9YonLYY0uUwKousR7C1wXjVuJtqGA6aTnsoIs/I9f3ctpEIkY9aInksvUFKurbk/0f+7FL5gNOmqWtk+Fv7TJc7oyp/bvgqHzG+jJkqscW9bTVvU4ow9kv3HFU6KyHriioEX/i6pAoGBAJs5yW4e3lrKh/u9ANPNVaRsRzF64V9zMBUKEpnGZy3aEcbfUiwFssZszINgUbvFGgsso22xcXGZ2IQWdhsFz84FwEpBodK+6tPfVXrkX2ZHICZXDcqrehpjPjR4ReC5MiGrK+BXHgcPKe6bmOd3YEQuB1zop/u4mpp98TgLAjptAoGAF1lgcmwnHDduvWErwSmMR3w45FppV0pOdSEmv/6DU6iglzleSuuCm4fhXAEgSahJMKiflNx1ufl84IGir0r/we8GNfs4ceSPlyKDYnaBG19f5jWLLNFTWfwVhZqa85pd9b2dcqVWMazHq8psmyt0opmctRcM9s4lNq2OmvB7vM4="
var publicKeyStr = "MIIBCgKCAQEA10X/4v44TU0sPOKhnHLx7Qe0SREVIne9+3DEtTCbzrewXSrlHedMT6RVlwv1i0D6q5zWGV9WXMI4ZrM2EM6wE4bZfjZU5FvzYTQN8fjiBQTggU8BDFpUPGibwt2Fyv6SAnj0AQzD7itvcxydu0tDOgg9Aaa701kxhz9yZukpZGVhAEBussRc41EuEwE2gQJ5prSFijzhqg5awwAJESX9gDHX3DncLO2s3FuaaReJ1U/c/LsJIvRGPckCllSJM1s9WAvjz5yFXZMpvWpmAnmpxd5fNd349Gr/swfM50TN5D3cTJblO771FifNFm6r+4o1nkR/thWX5vMasvKwimGo3QIDAQAB"

func TestGetParams(t *testing.T) {
	c := core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey, // 客户私钥
		PublicKey:  publicKeyStr,  // 公钥
		Key:        "870abfc720f86ce2c5e4d3345741d48d",
		BaseURL:    "http://127.0.0.1:9000",
	}
	core, err := core2.NewCore(&c, core2.WithSignType(consts.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	orderReq := &OrderRequest{
		OutBizNo:   "out_biz_no",
		ActivityNo: "activity_no",
		Number:     1,
	}
	p, err := core.GetParams(orderReq)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p)
}

func TestOrder(t *testing.T) {
	c := core2.Config{
		AppID:      "123",
		PrivateKey: rsaPrivateKey, // 客户私钥
		PublicKey:  publicKeyStr,  // 公钥
		Key:        "870abfc720f86ce2c5e4d3345741d48d",
		BaseURL:    "http://127.0.0.1:9000",
	}
	core, err := core2.NewCore(&c, core2.WithSignType(consts.SignRSA))
	if err != nil {
		t.Error(err)
		return
	}
	a := &Acquire{core}
	orderReq := &OrderRequest{
		OutBizNo:   "out_biz_no",
		ActivityNo: "activity_no",
		Number:     1,
	}
	r, err := a.Order(orderReq)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(r)
}
