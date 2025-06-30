package key

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"io"
	"net/http"
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

func newCore() *Key {
	c, _ := core.NewCore(&core.Config{
		AppID:      appId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Key:        key,
		SignType:   signType,
		BaseURL:    baseURL,
	})
	k := Key{c}

	return &k
}

func TestOrder(t *testing.T) {

	c := newCore()

	_, r, err := c.Order(context.Background(), &OrderRequest{
		OutBizNo:   "0627002001",
		ActivityNo: "0627002",
		Number:     1,
		NotifyUrl:  "",
		Account:    "18479266021",
		Extra:      "",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if !r.IsSuccess() {
		t.Errorf("获取key失败:%s", r.Message)
		return
	}

	t.Logf("data=%s", string(r.Data))
	t.Logf("headers=%+v", c.Headers)
}

func TestQuery(t *testing.T) {

	c := newCore()

	_, r, err := c.Query(context.Background(), &QueryRequest{
		OutBizNo: "",
		TradeNo:  "794149313755095041",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if !r.IsSuccess() {
		t.Errorf("查询失败:%s", r.Message)
		return
	}

	t.Logf("data=%s", string(r.Data))
	//t.Log(result.Status.IsNormal())
}

func TestDiscard(t *testing.T) {

	c := newCore()

	_, r, err := c.Discard(context.Background(), &DiscardRequest{
		OutBizNo: "N123456003",
		TradeNo:  "",
		Reason:   "正常作废",
	})
	if err != nil {
		t.Error(err)
		return
	}

	if !r.IsSuccess() {
		t.Errorf(r.Message)
		return
	}

	t.Logf("respons=%+v", r)
}

func TestResponse(t *testing.T) {

	jsonBytes := []byte(`{"code":200,"data":{},"message":"成功"}`)
	resp, err := core.BuildResponse(jsonBytes)
	if err != nil {
		t.Error(err)
		return
	}

	result, err := ConvertData(resp.Data)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", result)
}

func TestNotify(t *testing.T) {

	data := NotifyData{
		NotifyId:        "7278418772598218752",
		OutBizNo:        "006",
		TradeNo:         "727291384764309505",
		Key:             "dpK5yorx2M2g2e0W",
		UsableNum:       1,
		UsageNum:        1,
		Status:          3,
		Url:             "https://market.86698.cn/dpK5yorx2M2g2e0W",
		Amount:          0,
		PayAmount:       0,
		PayTime:         "",
		SettlementPrice: 0,
		ValidBeginTime:  "2024-12-27 22:37:41",
		ValidEndTime:    "2024-12-27 18:08:25",
		UsageTime:       "2024-12-27 22:37:41",
		DiscardTime:     "",
	}
	n := &Notify{
		AppId:     "lzm",
		SignType:  "RSA",
		Timestamp: "2025-06-01 12:00:00",
		Sign:      "",
		Data:      data,
	}

	c := newCore()

	str := n.SignString()

	signStr, err := c.CryptographySuite.Signer.Sign(str)
	if err != nil {
		t.Error(err)
		return
	}
	n.Sign = signStr

	r, err := c.Notify(context.Background(), n)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("response=%+v", r)
}

func TestCallback(t *testing.T) {

	data := NotifyData{
		NotifyId:        "7278418772598218752",
		OutBizNo:        "006",
		TradeNo:         "727291384764309505",
		Key:             "dpK5yorx2M2g2e0W",
		UsableNum:       1,
		UsageNum:        1,
		Status:          3,
		Url:             "https://market.86698.cn/dpK5yorx2M2g2e0W",
		Amount:          0,
		PayAmount:       0,
		PayTime:         "",
		SettlementPrice: 0,
		ValidBeginTime:  "2024-12-27 22:37:41",
		ValidEndTime:    "2024-12-27 18:08:25",
		UsageTime:       "2024-12-27 22:37:41",
		DiscardTime:     "",
	}
	n := &Notify{
		AppId:     "lzm",
		SignType:  "RSA",
		Timestamp: "2025-06-01 12:00:00",
		Sign:      "",
		Data:      data,
	}

	c := newCore()

	signStr, _ := c.CryptographySuite.Signer.Sign(n.SignString())
	n.Sign = signStr

	body, _ := json.Marshal(n)

	_, err := c.BuildParams(&data)

	x, _ := json.Marshal(c.Headers)
	t.Logf("x=%s", x)

	req := &http.Request{
		Header: c.Headers,
		Body:   io.NopCloser(bytes.NewBuffer(body)),
	}

	r, err := c.CallBack(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("response=%+v", r)
}

func TestCallBack(t *testing.T) {

	reqStr := `{"data": {"url": "https://gateway.dev.cdlsxd.cn/yxh5/dpK5ly6oVVE2AM0W", "status": 2, "account": "18479266021", "trade_no": "794167617429315585", "notify_id": "7345294954732199936", "usage_num": 0, "out_biz_no": "0627002001", "usable_num": 2, "usage_time": "2025-06-30 11:39:46", "valid_end_time": "2026-06-30 23:59:59", "settlement_price": 2, "valid_begin_time": "2025-06-24 11:00:08"}, "sign": "XLwRQ12EBXSGOSVzUMXwjSlKP88P4Odhe6c9MrfaszKLe+3HtPTeB6QWvyAmXGeIvsy02P0YtcOYV4xQHlWo3Uh5FZc6IJU/+KN+xVnn/DlFLpc+DhCKw6o4hYv+eLLyshjFZPZYVUU2I2YmkI1ZlwBaufsB+N9ds8gBz5+hELn17/qcFcbO6pYOd2te7xmJSGKOAMn0q2c2DSvTLvyQXhKUlDZfUZZGBOc1LGChy9CHc7Z/0E8/p2YYTlMPnvk0VHjEjV5sJxDnXwhSZqE7f3mRx0IN3au3VtZnXJsgl/whxdTyab9dYpfIxK75bS0mjncdqxGf1hLdhYJhTx8bog==", "app_id": "lzm", "sign_type": "RSA", "timestamp": "2025-06-30 11:39:46"}`
	headerStr := `{"Sign": ["Boj6IrOOrRATJt0IBE+z5Ie/g4mo3MZk+JpJ4bLYoBbDfMqvgTBhxqiC8CheRm/nEF9iFFJCvq9S0dL25fLexQ1k5AxE3cX1+qR5fCRdaiZvqWG4jaXOjUUW8K7fQ9g5ii6T4b3cWp71FBHiG3ZH5XohM9JuLo3W17MxrizsLLD0euGROAY3bXcakVustto07V3i0g59+ajsCTTdxF/gNcrsO5a3eTJ8CTSDnMgpwqMbU+E9YMX1zGFH/+m/RtL9s8tLRf8j4T/t8g6b94JfvBv+Fu1wV4eMUO7H4Iv0LJ1TL8qMBkWul5BbwGxSdEGQoWU0CIAehYTfR5meKxTOTQ=="], "Appid": ["lzm"], "Version": ["1.0"], "Sign-Type": ["RSA"], "Timestamp": ["2025-06-30 11:39:46"], "Content-Type": ["application/json"]}`

	var headers http.Header
	if err := json.Unmarshal([]byte(headerStr), &headers); err != nil {
		t.Error(err)
		return
	}

	c := newCore()
	request := &http.Request{
		Header: headers,
		Body:   io.NopCloser(bytes.NewBuffer([]byte(reqStr))),
	}

	r, err := c.CallBack(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("response=%+v", r)
}
