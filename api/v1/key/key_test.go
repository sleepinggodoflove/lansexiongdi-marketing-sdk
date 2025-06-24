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
	appId      = "iLqdJUryfuUF1"
	privateKey = "MIIEogIBAAKCAQEAxKGQ3r0zrGiuHTiYNv2aPxKYNiBkXfbXDD2OR1dYWRAyKXerQbaxK0h8xBn70RDYRX93JojJg3TJjA16EHzO+Py3xxmzpqFNRNvcx5kdqhcprWiecO7AplbYmWmwqn53V/ScHiWWVlQeVeAz9sAp04cE1EpSRcXjqjYzUn6KlIatkFDwdhmAp/hsVb6nllj1YLDk8kQ4ynMs4gnyeAtNcyCDsowXG+gPvguLJ7XqlztTk/a3Omb2uQYkWY8v5d2SIm8aI6ik8w0WqLq/y8fNREU7sO4CUPOVmUryM5wYCjzlfYwTs7Fp784t4zSvV/u07w+4HjRSaib6SsJH6A8NSQIDAQABAoIBADSGwG9v9XTSBekSbD8MYf8FVZnn70fWn73KV86g+53XIL9JE3ubdRqN455zHnzL1ipCka6+ja5LLWf3Vas/2/5RbDqImCAkKMMDRKghM/Zy95Q2RVT/woDs5DGfNS1mWTUqx+WZvt5S3EOuf9MjpEi4YcbgD7hJ07ZtgipAWX6xPu0/CeBc1Cfdn4+vlLBlhKF3L2YwgBQk/eOpWMxySSI/BUQTwWa4Un8ikOSKi9LbdmapUvuy1A05dIC8Eni/9cmcY8NPzxC0KajO7WJvmxjHGTJVO/7lhzpLE4FOh3MD0XfWHDyGbwnC+6ZLuTIOiHge+rn8cMLUqtECWkeNwDECgYEA+eT/XWVZ5IPS3krK7SbcivNmIibqayayZBoO98KYCfmLJ+RulkKEn5SXEw5mcTsfDOZ4mLqewO+qX4cjI95jNK4D/mC1D1VX4K23D9gJDQJla/IytsqU7siI70540eUhcHaZk/mWiWTWtkRTSOfyDz7vDRw/AOIAnV0IBppLPG0CgYEAyW9sqyhl0Z73Vz82Kx5Xa3rhFaBnw5la83/Mlt4sXTe/6vmLXxfNYGgVShgenTzx+7/CrPfO4bhB0UcaXIenipnzy9LPDSY45dV0ZXzoCawifdJ+yAkTGB1rm8ydb5ArwSD7W/cSar1Evgoh94cLDtSXEgZ1qnLaWDmO+3PQEs0CgYANPCqmOKr1JmIxscZjnw5JMbD0GBmMSUVjddnbF2xUAupy129f9+/sP8Nsl/OnBZmUZlR2ylOEJm1gl9itmqaocJr3iwmr4TpBRRIP/cIk9T6H5BD5i2st5mMSQZa2jyshOLbTloF5j2SrzJyYnOg+FHg2uos/sbiUnQvCxcM8VQKBgD6Lf2VhpMPCpTAM64fV4vT3cX4ikTV1n3zt2JubnpDPJ6MRGspK/LULfFjGRnMyIjy13P+R7kW9zYnqlu/WGxp5FO7bpPiDPrV2Yq7EatPeA9OnkUAROUGKmQgGL3gdfsh3sjRq8ef8nqSXRtaxsqhHMCUjplNnWA5+yLugySGdAoGAMcghxV7PWFlVOxGYwwQlblpsxC5ObcJdP5PUWWxgjHguOl+SsivrZ+XZpjODuXe0kqyxpazbb8hhSDnSPrMMrQP+6jxoS8mkdCIZZ2ytuI8vMUoCMlrN0+JfbXUk+apzEtNZLFbfga4DKqEMC0mTxSUIL0nWIrHvid6yt4wSbCw="
	publicKey  = "MIIBCgKCAQEAxKGQ3r0zrGiuHTiYNv2aPxKYNiBkXfbXDD2OR1dYWRAyKXerQbaxK0h8xBn70RDYRX93JojJg3TJjA16EHzO+Py3xxmzpqFNRNvcx5kdqhcprWiecO7AplbYmWmwqn53V/ScHiWWVlQeVeAz9sAp04cE1EpSRcXjqjYzUn6KlIatkFDwdhmAp/hsVb6nllj1YLDk8kQ4ynMs4gnyeAtNcyCDsowXG+gPvguLJ7XqlztTk/a3Omb2uQYkWY8v5d2SIm8aI6ik8w0WqLq/y8fNREU7sO4CUPOVmUryM5wYCjzlfYwTs7Fp784t4zSvV/u07w+4HjRSaib6SsJH6A8NSQIDAQAB"
	key        = "6bf9bf97d2dca73c262172bd7488b849"
	baseURL    = "https://xytrans.86698.cn/ymttb/pacific/v1/notify"
	signType   = core.SignRSA
)

//var (
//	appId      = "lzm"
//	privateKey = "MIIEogIBAAKCAQEA0w4XGS1eEO9gAtWoB0E1vi1QH3xZAiHnkzZMhZRJOKeZhNUb9nmPzrGtCFD1c+to9/hxKnWZnRi1dklRGI4uXaB7PKuDhifHarBTTPBzW/8m+YqKEwjT2XWYvnG1Zeek4a45xze5cHhLA7Ow1Lwgy0u1rhalvz8GbCa9A7ZHKvZtIJJzfPSIV6gZIz5b7+v7rXZzMNNxvC7m+cwtvvERPjhJoj3O7ithcgdiT3JkZd1fZxkA6HCJx1I+TElt4qA9WnV+rqQwjka1gxBO497c0MUq+4Tx+lLGKpb61RPja4+9wiLFvEiS80WyZYkptWlA0Z5mhsxURs/OqaMyVXzbhwIDAQABAoIBACKQgy3nZSlm2pV8Qjl174RGzYFqjvUvckqEsQGLaHZz1EuRzzONcwTJymm7QIeMfTNnJ7lpaw/0VPubREG+P0+sEaK9ABw/dYQ+flXyZyIg5lQl4Tj+0BskDDDcVHXs+u7O5r1+ncsSmE5x5jrg2IoSyx1Irjpk6vtZWMk71+nsB3gC7AqynfzFmX8orIK7tdfWE1wyMHQ9czxdxcWqB2TyFZSBAf9INvEPyj9iBViRQ+P2K5qWURLmfdYwCCIPKEWa+CjATU1pP4o2FHnAwvsgaAIV1qOjmbkvxkemAnq8P1buJn4p82WcoklIyFEaJYrdljGMI8rOEL0rVzuNrnkCgYEA893qdXTrEAlBlytKsCAkgv/6tNfaAeKDKeD5np1BYvDBi4xVZt2aPPMz8bZT4OLfhSqxEbn+jVSVVvn6d4SrEKzUfnm8zVQrvw4T6qShx+IBJrt4Jz8xEbIE/L1dGiiO8fVTXqxGxtR4tg98jsBNssZrwc+/PhD5tmWHicaqkIMCgYEA3Y5Btw7Rqh+ElAewm62eSidkJkli2IZoLet017OpycGoY6BB4tKGEe6mI1y+hGdtqH/z1bCU9fSSqybCKIZhHoo0E2n2fAQFPCoZyMkqyzvSOp7K31nuLTyw0las7rDrCqUb+u92d2LA3xiO2E6EOXclvP4rOif1QYYajTn1ka0CgYAmt7PyxAZR/HY6bvgjsGa7mbKPJboKFJFog7x970+jSsAfzL7+Xu6PALndhWoZyUtdlCKawuHkRGqVbYjTku+p7Rarod5U5yku4yhMV6kL2BkAskDoUkMTISVjjxkJ/yh6x81duZJfHPqxRRIsg+GSIaiYE8i0LPPIfqQfPrhzywKBgG2HWUd1RQOUh0djMdUUlM9V//XJi9s9Px2MbHwCbuq2GVf/LvNCXlNZJrsOq4TrWNPXRaUbodih6yw/gfbkz/h4HFyIovkWR7xBl/OiN8y3Kywdum+Glu+4NDYX8XAi+F+P1nBMl8VXhcAE9QFMd6OtGnP/N0GN5XpaIKA2ygcVAoGAbJpF93LT92HkzUqxh05PX1QcmW9gAU24cO/Mc1WRUg7BQz+JR7AeheZX430Qk9j0JP+CtKWnqaaa5K1ym6m/r3uTj8Xx+CF9DOy9OZNrELSgveVSN+oAXAUeSemUbSV7499oLLYQmYBhWB7a6lREk2O8FEfe2Yb8ySNApWUv/js="
//	publicKey  = "MIIBCgKCAQEA0w4XGS1eEO9gAtWoB0E1vi1QH3xZAiHnkzZMhZRJOKeZhNUb9nmPzrGtCFD1c+to9/hxKnWZnRi1dklRGI4uXaB7PKuDhifHarBTTPBzW/8m+YqKEwjT2XWYvnG1Zeek4a45xze5cHhLA7Ow1Lwgy0u1rhalvz8GbCa9A7ZHKvZtIJJzfPSIV6gZIz5b7+v7rXZzMNNxvC7m+cwtvvERPjhJoj3O7ithcgdiT3JkZd1fZxkA6HCJx1I+TElt4qA9WnV+rqQwjka1gxBO497c0MUq+4Tx+lLGKpb61RPja4+9wiLFvEiS80WyZYkptWlA0Z5mhsxURs/OqaMyVXzbhwIDAQAB"
//	key        = "5f42e758a38cc003c5da7cee814ddfd5"
//	baseURL    = "https://gateway.dev.cdlsxd.cn"
//	signType   = core.SignRSA
//)

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

	_, r, err := c.Order(context.Background(), &OrderRequest{
		OutBizNo:   "N123456001",
		ActivityNo: "N123456",
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

	t.Logf("response=%+v", r)
	if !r.IsSuccess() {
		t.Errorf("查询失败:%s", r.Message)
		return
	}

	t.Logf("data=%s", string(r.Data))
	//t.Log(result.Status.IsNormal())
	//t.Log(result.Status.IsUsed())
	//t.Log(result.Status.IsDiscardIng())
	//t.Log(result.Status.IsDiscard())
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

	t.Logf("response=%+v", r)
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
	t.Logf("%+v", resp)
	t.Logf("%s", string(resp.Data))
	t.Logf("%+v", result)

	jsonBytes2 := []byte(`{"code":200,"message":"成功","data":{"out_biz_no":"123","trade_no":"456"}}`)
	resp2, err := core.BuildResponse(jsonBytes2)
	if err != nil {
		t.Error(err)
		return
	}
	result2, err := ConvertData(resp2.Data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", resp2)
	t.Logf("%s", string(resp2.Data))
	t.Logf("%+v", result2)
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

	// 测试
	reqStr := `{"data": {"key": "3GZwYyVg1Kg9Myop", "status": 2, "trade_no": "791668935178067969", "notify_id": "7342795997099393024", "usage_num": 0, "out_biz_no": "N123456001", "usable_num": 2, "usage_time": "2025-06-23 14:09:48", "valid_end_time": "2025-07-05 23:59:59", "settlement_price": 10.01, "valid_begin_time": "2025-02-21 00:00:00"}, "sign": "dPq0FrvOG7S+3eY8hQ6sg0uc+xb1F2ymWim0my+WGBHvg+U4qiT9HQ58ntXMxz/QAhEEpFGcoJXYqyYR1ZYEPnCdtkP0yCX3BucBW25NB7CWjoIO57akfKbNf9aZpx7xV3toYVrcIuXJNiJ8GxfIK1ybFZQlOvdlQdRE3NqAoAQtO0y6QZPAf3pziyRFHk77bIMoTXynmSf0FhtJYRDpURAntp8s4cJ5F/n3beAQkJnape8zHzkHKMIfr+3HEFYt3qEyoT3U2nRtxhWyrLSlo15KZ5a4yq4QdQJdcEM6KJedISKQnkeqIPRh7sKviVCqqRD/fMdak/Z/tqvNMGIbrQ==", "app_id": "lzm", "sign_type": "RSA", "timestamp": "2025-06-23 14:09:48"}`
	// 生产
	//reqStr := `{"data": {"url": "https://market.86698.cn/3GZwYQNqjKeR4yop", "status": 2, "trade_no": "789856868900282369", "notify_id": "7341346520502038528", "usage_num": 1, "out_biz_no": "25061814091189343", "usable_num": 1, "usage_time": "2025-06-19 14:10:06", "valid_end_time": "2025-06-30 09:35:01", "settlement_price": 0.1, "valid_begin_time": "2025-05-27 09:35:04"}, "sign": "Ki+JXSz0so3JmKOK6kKvBaLlsgcypmC0Rv0IG39ZIErWn82XfvxhdXh2y9sQAW7BplTCNzlMeSLp/ml+SXoxrgAaswxwNP3B5kMWyXNF6K/S7RWrZFFmW50PrnKZSnOOmP1nM7T0X/+rhkaN+DF+WcMK2FA/BIgoauBHJg+75UqI5YROpXj4fJjKXe8AAUthhv6Zngtx/Ep50KqgfxsHXnvlXrxxLhXrXRjyXoA0rwD7uD44Wl1e4R041pdxa/+I/BrJeLTPW++MYkuOZ7uRAWn/4LdMB04KfFi9s4ex4dT4RtKSB8El3n5nOYBlAX9ksYeCL+So6NPwj9VpSpF8bQ==", "app_id": "iLqdJUryfuUF1", "sign_type": "RSA", "timestamp": "2025-06-19 14:10:06"}`
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
