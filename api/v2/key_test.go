package v2

import (
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"testing"
)

var (
	appId      = ""
	privateKey = ""
	publicKey  = ""
	key        = ""
	baseURL    = ""
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
}

func TestNotify(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}

	n := &Notify{
		AppId:     "",
		SignType:  "",
		Timestamp: "",
		Sign:      "",
		Data: NotifyData{
			Event:            "",
			NotifyId:         "",
			OutBizNo:         "",
			TradeNo:          "",
			ActivityNo:       "",
			Number:           0,
			Status:           0,
			KeyMapCiphertext: "",
		},
	}

	str, err := n.SignString()
	if err != nil {
		t.Error(err)
		return
	}

	sign, err := c.CryptographySuite.Signer.Sign(str)
	if err != nil {
		t.Error(err)
		return
	}

	n.Sign = sign

	b := c.CryptographySuite.Verifier.Verify(str, sign)
	if !b {
		t.Error("验签失败")
		return
	}

	keyMapDecode, err := c.CryptographySuite.Cipher.Decode(n.Data.KeyMapCiphertext)
	if err != nil {
		t.Error(err)
		return
	}

	keyMap := make([]*KeyInfo, 0, n.Data.Number)
	if err = json.Unmarshal([]byte(keyMapDecode), &keyMap); err != nil {
		t.Error(err)
		return
	}
}
