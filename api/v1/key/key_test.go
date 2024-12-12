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

	c *core.Core
	k *Key
}

func TestMathSuite(t *testing.T) {
	suite.Run(t, new(KeySuite))
}

func (s *KeySuite) SetupTest() {
	c, err := core.NewCore(&core.Config{
		AppID:      appId,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Key:        key,
		SignType:   signType,
		BaseURL:    baseURL,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.c = c
	s.k = &Key{c}
}

func (s *KeySuite) TestBuildParams() {
	req := &OrderRequest{
		OutBizNo:   "321312",
		ActivityNo: "lzm",
		Number:     1,
	}
	p, err := s.c.BuildParams(req)
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("%+v", p)
}

func (s *KeySuite) TestOrder() {
	r, err := s.k.Order(context.Background(), &OrderRequest{
		OutBizNo:   "20241211002",
		ActivityNo: "lzm",
		Number:     1,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("response=%+v", r)
	if !r.IsSuccess() {
		s.T().Errorf("获取key失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("data=%+v", data)
}

func (s *KeySuite) TestQuery() {
	r, err := s.k.Query(context.Background(), &QueryRequest{
		OutBizNo: "20241211002",
		TradeNo:  "",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("response=%+v", r)
	if !r.IsSuccess() {
		s.T().Errorf("查询失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("data=%+v", data)
	//s.Equal(true, data.Status.IsNormal(), "是否正常")
	//s.Equal(true, data.Status.IsUsed(), "是否已核销")
	//s.Equal(true, data.Status.IsDiscardIng(), "是否作废中")
	//s.Equal(true, data.Status.IsDiscard(), "是否已作废")
}

func (s *KeySuite) TestDiscard() {
	r, err := s.k.Discard(context.Background(), &DiscardRequest{
		OutBizNo: "20241211002",
		TradeNo:  "",
		Reason:   "正常作废",
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("response=%+v", r)
	if !r.IsSuccess() {
		s.T().Errorf("作废收单失败:%s", r.Message)
		return
	}
	data, err := r.ConvertData()
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("data=%+v", data)
	s.Equal(true, data.Status.IsDiscardIng(), "是否作废中")
}

func (s *KeySuite) TestNotify() {
	data, err := s.k.Notify(context.Background(), &Notify{
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
		s.T().Fatal(err)
	}
	s.T().Logf("data=%+v", data)
}

func (s *KeySuite) TestCallback() {
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
	signStr, err := s.k.CryptographySuite.Signer.Sign(n.SignString())
	if err != nil {
		s.T().Fatal(err)
	}
	n.Sign = signStr

	r, err := s.k.Notify(context.Background(), n)
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("r=%+v", r)
}

func (s *KeySuite) TestResponse() {
	jsonBytes := []byte(`{"code":200,"data":{},"message":"成功"}`)
	resp, err := response(jsonBytes)
	if err != nil {
		s.T().Fatal(err)
	}
	result, err := resp.ConvertData()
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("%+v", resp)
	s.T().Logf("%s", string(resp.Data))
	s.T().Logf("%+v", result)

	jsonBytes2 := []byte(`{"code":200,"message":"成功","data":{"out_biz_no":"123","trade_no":"456"}}`)
	resp2, err := response(jsonBytes2)
	if err != nil {
		s.T().Fatal(err)
	}
	result2, err := resp2.ConvertData()
	if err != nil {
		s.T().Fatal(err)
	}
	s.T().Logf("%+v", resp2)
	s.T().Logf("%s", string(resp2.Data))
	s.T().Logf("%+v", result2)
}
