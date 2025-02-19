package v2

import (
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"testing"
)

var (
	appId      = "GKETQmUiARPL"
	privateKey = "MIIEpQIBAAKCAQEA0jEraBz0E1sWF4NDYSD+REMUDINTdo3P4wnfVPuJ83v8cdqKNQtXQmg887gWUX5cj+4xsCTRBjyeKvKgROKYDrqamreuVM+CIp0SaauE3P9AZxjUW7q8W4+OT3yfEXjMhnrNaiC/TzomeLcLW8MYWthM6EbOn1Ko31DHsQto+bJm5wAp5FBoay9cKPvYT2D74ccbvUdh6LHeeKtL67gH4RHS+P9qP6EF2JtF3fkIsxdjISP5qXFpP09EmZ5Gx9tbXB8N08ni2j4O47SPVW6UnQM6PAtJWjFb9N8TTJHPqptUpdi2uPjEFXDEfSphgd2g0omvNraPFvx7FMHQhd/uzQIDAQABAoIBAQDI0GstFARofayUmDcGk/P5GbEM52cVLBWCTwtM8Ojyc/FSgT1rwkMC0f3xx12jTDt401QrenEtKTrfw2A1j9tAry1IRdbLdllZYoGV7WWJkmBgX0t7u+N7AqMu89wxYBzfGnIoQ9MjCWZ6DD9Q0wrwuBh6DjQX6WwnttCiKEmJzBnfiYAzonzEEcKwEgppBC/e6DP7KIlbbovh7Db5D10PhCllCY1X0bZoJinV36lrlZnsYhf8nYlF5kD36lSzXLWk2mZSu8ZLyminsSbNLAsff0gE+UMts7aI7HbxPn/+GgGbuRYqi7XnEJuS6kmfObZAoIiWWGuNVeRnKhhMamBBAoGBAOf4Rx++Cn4d6J8Gpesajd6lxKStYpH2WNYx+kL02lgxKIX3IlOZfPMHHCytSQomG8nPz8mt4VpplIcIMxQjgCmeAD/2/aoCBLsHsPqZDDV2wIiLSfYe2dvAhKOzS8GzQtSqgIFzAPyWn5IPKeQDRWo2OL/fc2ZZFdGIFR+2SNzdAoGBAOf3Wz50WsHgIgQThQeLgC5sK0YgzmaIyM863viQVJpacTG5IryBhN7Q32fWp4BH/uJR4XjLNW4pv4l3WkbSD6QR8d2UHC3CErRrb1nAloZWKCtQL2kHP9GQsL0i4t44l3paWDiHg0sCrUFpjMv5JjRikMDM/Vvd43MhQbK1w4KxAoGAa+NkqRXJaYjdlYEREDzkeQZeZD0kGEEcZlsOS6/4EYajk1MzCvDbVWkcKIdb7jV1PTLDMMkHg/aRFxCwORCd3j1XXmiw6C5SJu7X2GXwcLlPqPInryoAJ53t0vlkJa4LSkAWzp4/ejtP8i3NTPhcg56+XDdAQ/zxEegZ8wrj7p0CgYEAv5okpwjK8ntk/YL2Dg5fWXkFoY51jmILpIiXJi9pgnmUKCMpGxXMn5NeBVJb1u0hDuXdXL5VKuKXEEAitH7MZpRf+MtfHzi/5IEdX8BhVSMTWPuJvbI8N6jRI9kOFcoe7PR1DHR6sVnLrE03/D+XHAwSTv7Dg79bIXzl3GygynECgYEAyzm/cVK8xF7P2Uxx9NFJjs2EU7763rS2fAPHh6/uXOSfyJUPPbg29g1+SJeUGUIIEJAr3yBt6tyC7MwCGVDyI5pvqEyKmbbdBdOaX+psQnIsHz+JFixpelhcpL8GvhUEVxqfDsfjosMpeK4fi4r5C49wysS9tvefE8HBN8+CMWc="
	publicKey  = "MIIBCgKCAQEA0jEraBz0E1sWF4NDYSD+REMUDINTdo3P4wnfVPuJ83v8cdqKNQtXQmg887gWUX5cj+4xsCTRBjyeKvKgROKYDrqamreuVM+CIp0SaauE3P9AZxjUW7q8W4+OT3yfEXjMhnrNaiC/TzomeLcLW8MYWthM6EbOn1Ko31DHsQto+bJm5wAp5FBoay9cKPvYT2D74ccbvUdh6LHeeKtL67gH4RHS+P9qP6EF2JtF3fkIsxdjISP5qXFpP09EmZ5Gx9tbXB8N08ni2j4O47SPVW6UnQM6PAtJWjFb9N8TTJHPqptUpdi2uPjEFXDEfSphgd2g0omvNraPFvx7FMHQhd/uzQIDAQAB"
	key        = "36bee3e6b2ff8882436f5580f5df52ff"
	baseURL    = "https://gateway.dev.cdlsxd.cn"
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
