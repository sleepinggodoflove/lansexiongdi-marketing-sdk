package key

type AcquireRequest struct {
	AppId      string `json:"appIdd"`
	SignType   string `json:"signType"`
	Method     string `json:"method"`
	Version    string `json:"version"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
	BizContent string `json:"bizContent"`
}

type AcquireResponse struct {
	OutBizNo       string `json:"outBizNo"`
	OrderId        string `json:"orderId"`
	Key            string `json:"key"`
	ShortUrl       string `json:"ShortUrl"`
	ValidBeginDtTm string `json:"validBeginDtTm"`
	ValidEndDtTm   string `json:"validEndDtTm"`
}

type AcquireReply struct {
	Code     string           `json:"code"`
	Msg      string           `json:"msg"`
	SubCode  string           `json:"subCode"`
	SubMsg   string           `json:"subMsg"`
	Response *AcquireResponse `json:"response"`
	Sign     string           `json:"sign"`
}
