package key

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
