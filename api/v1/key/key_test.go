package key

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"github.com/stretchr/testify/assert"
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

func TestBuildParams(t *testing.T) {
	c, err := newCore()
	if err != nil {
		t.Error(err)
		return
	}
	req := &OrderRequest{
		OutBizNo:   "321312",
		ActivityNo: "lzm",
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
	r, err := a.Order(context.Background(), &OrderRequest{
		OutBizNo:   "20241211002",
		ActivityNo: "lzm",
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
	r, err := a.Query(context.Background(), &QueryRequest{
		OutBizNo: "20241211002",
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
	r, err := a.Discard(context.Background(), &DiscardRequest{
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
