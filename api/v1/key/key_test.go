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

//var (
//	appId      = ""
//	privateKey = ""
//	publicKey  = ""
//	key        = ""
//	baseURL    = ""
//	signType   = core.SignRSA
//)

var (
	appId      = "lzm"
	privateKey = "MIIEowIBAAKCAQEA7a2I4l8OOdW4weVFvj4u/mBqP3aZhJ0mOTKl4MCW4Pf6gNAlZa5dZYOS/BocmG872+pd10BiI73qiAWsuVaPwCL0A37lQbCXlG0fDAfCLogXuF1qVNRZgkYKrx/5Gppo2PNed7E5YyCUkMUKVPbuwuZteMZJH8d1o6Uojbb/xJQvAGOlx5Y04VZWp/6p2GjhW0srwgbpVegMyyn2Qblx1Lo+Uq5zG8um7FTpbtb/L/itpBFEDSZGIIKDfn4FPyt+jQ0SW5TDYQClSvWHK4V3RkWOVkD1nHeBpZyp7JNehK+7kBfO6G4NJabkyoWqFEiZcTy38ZWQdqJ9N4LZuY37NwIDAQABAoIBAGs2u6e5z1YBda1pehN+Q36WCXeFTW0H4qUslq0S0zy6P/L5cdUzWYggWR6FvN56Vts2Foyxy1NqKTCgtrCIPqIiYkZtaIdAXLAkpTutCEgrNeABq6SGgbYFWG51Es6QVrl+1t9RP5zaponDiIyZM00R2tH/SB8gv41JREjhAvEuNIwPyaoVVt+U/kAdhJgiMKsDpoGaMfsJk76sORu6qQqBkBN8cglN94xC0QtROytW3EY8SnZmgGZHcY3YTXM74CWM8yBg7rNuKv0982f9hKvUDHKMFYly1PzYiSgplkT7RYCMjo2FFf1lt7k7N61+4nalS/EM6324m2poisTRFAkCgYEA+3oXvtwrTim0kXG7V7w5PS8u5dU0sAAH4ACSyzy8nEdKEk4ipoaTGm6km8ko0O+9E70SwZQgK/eAnmsYf5WtMzvItweKUUVsBCm01qSHlu5vzGO3H1ndi7hg+tH9VrOQH3+odQJP9FqC4BkncMszHM4nLglWSTixTTvGIovQLy0CgYEA8fPnu0tWqhfQV6svaA5kt4h6cL52ARKlubRuYkI4hGuikKYpd2A3WuVtD1LkuPQSjwID9730HAqLc7ZMwONjQ8NANi9ZoJR6A+Vzba9zDPQmSc80Ax7Kkjc03D1Y7yiP6P8bWnhCCbRcMy+dcobvBZc2zaWzSNjZwPOSV9xaMnMCgYAclz354hg+U7mGy7JsACdV0HZ5hOrvk6FRk18dIjOjZOuD90QzQJua5rdqSs2MK6WIh/eI8KlTtlj2KeDoKIE/kO15+a59HPJx6rf3q08LFuK5DyEzvEjW6MiF27f80n9xRVdGrlOeyWeVyOZWCZQvEzUbI86eloZ57HDTXqf1pQKBgF6T6xeJgZ0Hpgc/AU75oWEk1kfQC6yrr2CCKUv7esA4mtlUOo1RbRH48MK2snWh4sdIEGj9NbjoXk6jCim0OQ85+ZW0uKJOp8tyG8baeGyt23GqrzgxBxpUvjMBQAxsnKSFZBnfPGEywX+4syEbob9btq54gTaOncAQ9jmmBxQFAoGBAIpPbq2lYwOhgoUJ2BR34xjpmNOiOAF5AVLPGTH44a+iGMJ4tbF9AvfL4xsCWK9zMi3ExaKVN0lNn0cWx2lpXxwO6B+l4L//eczmHx4h1eLJd6ZWyTj7lq+RBOOUgHLKEssZfJ11RYTZjSD7s75JZteM2OFw7BVRRNgw387A3mj6"
	publicKey  = "MIIBCgKCAQEA0w4XGS1eEO9gAtWoB0E1vi1QH3xZAiHnkzZMhZRJOKeZhNUb9nmPzrGtCFD1c+to9/hxKnWZnRi1dklRGI4uXaB7PKuDhifHarBTTPBzW/8m+YqKEwjT2XWYvnG1Zeek4a45xze5cHhLA7Ow1Lwgy0u1rhalvz8GbCa9A7ZHKvZtIJJzfPSIV6gZIz5b7+v7rXZzMNNxvC7m+cwtvvERPjhJoj3O7ithcgdiT3JkZd1fZxkA6HCJx1I+TElt4qA9WnV+rqQwjka1gxBO497c0MUq+4Tx+lLGKpb61RPja4+9wiLFvEiS80WyZYkptWlA0Z5mhsxURs/OqaMyVXzbhwIDAQAB"
	key        = "5f42e758a38cc003c5da7cee814ddfd5"
	baseURL    = "https://gateway.dev.cdlsxd.cn"
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

func TestBuildParams(t *testing.T) {

	strBytes := []byte(`{"out_biz_no":"","activity_no":"","number":1}`)

	var r *OrderRequest
	if err := json.Unmarshal(strBytes, &r); err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", r)

	c := newCore()

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

	c := newCore()

	a := []AccountReq{
		{
			Account: "18666666661",
			Type:    AccountTypeMobile,
			Verify:  true,
		},
		{
			Account: "18666666662",
			Type:    AccountTypeMobile,
			Verify:  true,
		},
	}
	var _, r, err = c.Order(context.Background(), &OrderRequest{
		OutBizNo:   "N123456002",
		ActivityNo: "N123456",
		Number:     1,
		NotifyUrl:  "",
		Accounts:   a,
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
}

func TestQuery(t *testing.T) {

	c := newCore()

	_, r, err := c.Query(context.Background(), &QueryRequest{
		OutBizNo: "N123456001",
		TradeNo:  "",
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
		OutBizNo: "N123456001",
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

	t.Logf("data=%s", string(r.Data))
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

func TestRequestNotify(t *testing.T) {

	var req *Notify

	reqStr := `{"data": {"key": "3GZwYyVg1Kg9Myop", "status": 2, "trade_no": "791668935178067969", "notify_id": "7342795997099393024", "usage_num": 0, "out_biz_no": "N123456001", "usable_num": 2, "usage_time": "2025-06-23 14:09:48", "valid_end_time": "2025-07-05 23:59:59", "settlement_price": 10.01, "valid_begin_time": "2025-02-21 00:00:00"}, "sign": "dPq0FrvOG7S+3eY8hQ6sg0uc+xb1F2ymWim0my+WGBHvg+U4qiT9HQ58ntXMxz/QAhEEpFGcoJXYqyYR1ZYEPnCdtkP0yCX3BucBW25NB7CWjoIO57akfKbNf9aZpx7xV3toYVrcIuXJNiJ8GxfIK1ybFZQlOvdlQdRE3NqAoAQtO0y6QZPAf3pziyRFHk77bIMoTXynmSf0FhtJYRDpURAntp8s4cJ5F/n3beAQkJnape8zHzkHKMIfr+3HEFYt3qEyoT3U2nRtxhWyrLSlo15KZ5a4yq4QdQJdcEM6KJedISKQnkeqIPRh7sKviVCqqRD/fMdak/Z/tqvNMGIbrQ==", "app_id": "lzm", "sign_type": "RSA", "timestamp": "2025-06-23 14:09:48"}`

	err := json.Unmarshal([]byte(reqStr), &req)
	if err != nil {
		t.Error(err)
		return
	}

	c := newCore()

	r, err := c.Notify(context.Background(), req)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("response1=%+v", r)

	request := &http.Request{
		Header: c.Headers,
		Body:   io.NopCloser(bytes.NewBuffer([]byte(reqStr))),
	}
	r, err = c.CallBack(context.Background(), request)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("response2=%+v", r)
}
