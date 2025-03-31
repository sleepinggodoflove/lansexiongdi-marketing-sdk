package anyapi

import (
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"testing"
)

var (
	appId      = "OP001"
	privateKey = ""
	publicKey  = ""
	key        = ""
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
		MchPublicKey: "",
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
