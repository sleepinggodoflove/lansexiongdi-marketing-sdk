package key

type Request struct {
	AppId      string `json:"app_id"`
	SignType   string `json:"sign_type"`
	Method     string `json:"method"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
	Ciphertext string `json:"cipher_text"`
}

type Response struct {
	OutBizNo       string `json:"out_biz_no"`
	TradeNo        string `json:"trade_no"`
	Key            string `json:"key"`
	Status         string `json:"status"`
	Url            string `json:"url"`
	ValidBeginTime string `json:"valid_begin_time"`
	ValidEndTime   string `json:"valid_end_time"`
}

type AcquirePlaintext struct {
	OutBizNo   string `json:"out_biz_no"`
	ActivityNo string `json:"activity_no"`
	Number     string `json:"number"`
}

type AcquireReply struct {
	Code     string    `json:"code"`
	Msg      string    `json:"msg"`
	SubCode  string    `json:"subCode"`
	SubMsg   string    `json:"subMsg"`
	Response *Response `json:"response"`
	Sign     string    `json:"sign"`
}

type DiscardPlaintext struct {
	OutBizNo string `json:"out_biz_no"`
	TradeNo  string `json:"trade_no"`
	Key      string `json:"key"`
}

type DiscardReply struct {
	Code     string    `json:"code"`
	Msg      string    `json:"msg"`
	SubCode  string    `json:"subCode"`
	SubMsg   string    `json:"subMsg"`
	Response *Response `json:"response"`
	Sign     string    `json:"sign"`
}
