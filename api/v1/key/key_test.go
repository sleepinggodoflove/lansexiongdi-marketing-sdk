package key

import (
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestBuildParams(t *testing.T) {
	strBytes := []byte(`{"out_biz_no":"","activity_no":"","number":1}`)
	var r *OrderRequest
	if err := json.Unmarshal(strBytes, &r); err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", r)

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
