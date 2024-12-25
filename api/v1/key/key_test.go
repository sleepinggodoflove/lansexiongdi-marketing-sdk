package key

import (
	"context"
	"github.com/sleepinggodoflove/lansexiongdi-marketing-sdk/core"
	"github.com/stretchr/testify/suite"
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

type KeySuite struct {
	suite.Suite

	kService *Key
}

func TestMathSuite(t *testing.T) {
	suite.Run(t, new(KeySuite))
}

func (k *KeySuite) SetupTest() {
	c, err := core.NewCore(&core.Config{
		AppID:      appId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Key:        key,
		SignType:   signType,
		BaseURL:    baseURL,
	})
	if err != nil {
		k.T().Fatal(err)
	}

	k.kService = &Key{c}
}

func (k *KeySuite) TestBuildParams() {
	req := &OrderRequest{
		OutBizNo:   "321312",
		ActivityNo: "lzm",
		Number:     1,
	}
	p, err := k.kService.BuildParams(req)
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("%+v", p)
}

func (k *KeySuite) TestOrder() {
	req := &OrderRequest{
		OutBizNo:   "20241211002",
		ActivityNo: "lzm",
		Number:     1,
	}
	r, err := k.kService.Order(context.Background(), req)
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("response=%+v", r)
	if !r.IsSuccess() {
		k.T().Errorf("获取key失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("data=%+v", data)
}

func (k *KeySuite) TestQuery() {
	req := &QueryRequest{
		OutBizNo: "20241211002",
		TradeNo:  "",
	}
	r, err := k.kService.Query(context.Background(), req)
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("response=%+v", r)
	if !r.IsSuccess() {
		k.T().Errorf("查询失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("data=%+v", data)
	//kService.Equal(true, data.Status.IsNormal(), "是否正常")
	//kService.Equal(true, data.Status.IsUsed(), "是否已核销")
	//kService.Equal(true, data.Status.IsDiscardIng(), "是否作废中")
	//kService.Equal(true, data.Status.IsDiscard(), "是否已作废")
}

func (k *KeySuite) TestDiscard() {
	req := &DiscardRequest{
		OutBizNo: "20241211002",
		TradeNo:  "",
		Reason:   "正常作废",
	}
	r, err := k.kService.Discard(context.Background(), req)
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("response=%+v", r)
	if !r.IsSuccess() {
		k.T().Errorf("作废收单失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("data=%+v", data)
	k.Equal(true, data.Status.IsDiscardIng(), "是否作废中")
}

func (k *KeySuite) TestNotify() {
	req := &Notify{
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
	}
	data, err := k.kService.Notify(context.Background(), req)
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("data=%+v", data)
}

func (k *KeySuite) TestCallback() {
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
	signStr, err := k.kService.CryptographySuite.Signer.Sign(n.SignString())
	if err != nil {
		k.T().Fatal(err)
	}
	n.Sign = signStr

	r, err := k.kService.Notify(context.Background(), n)
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("r=%+v", r)
}

func (k *KeySuite) TestResponse() {
	jsonBytes := []byte(`{"code":200,"data":{},"message":"成功"}`)
	resp, err := response(jsonBytes)
	if err != nil {
		k.T().Fatal(err)
	}
	result, err := resp.ConvertData()
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("%+v", resp)
	k.T().Logf("%s", string(resp.Data))
	k.T().Logf("%+v", result)

	jsonBytes2 := []byte(`{"code":200,"message":"成功","data":{"out_biz_no":"123","trade_no":"456"}}`)
	resp2, err := response(jsonBytes2)
	if err != nil {
		k.T().Fatal(err)
	}
	result2, err := resp2.ConvertData()
	if err != nil {
		k.T().Fatal(err)
	}
	k.T().Logf("%+v", resp2)
	k.T().Logf("%s", string(resp2.Data))
	k.T().Logf("%+v", result2)
}
