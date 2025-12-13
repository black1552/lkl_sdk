package model

type MrchAuthStateQueryRequest struct {
	ReqData   *MrchAuthStateQueryReqData `json:"reqData"`
	Ver       string                     `json:"ver"`
	Timestamp int64                      `json:"timestamp"`
	ReqId     string                     `json:"reqId"`
}

type MrchAuthStateQueryReqData struct {
	TradeMode     string `json:"tradeMode" dc:"交易模式"`
	SubMerchantId string `json:"subMerchantId" dc:"子商户号"`
	MerchantNo    string `json:"merchantNo" dc:"商户号"`
}

type MrchAuthStateQueryResponse struct {
	RetCode  string                      `json:"retCode"`
	RetMsg   string                      `json:"retMsg"`
	RespData *MrchAuthStateQueryRespData `json:"respData"`
}

type MrchAuthStateQueryRespData struct {
	SubMerchantId string `json:"subMerchantId"`
	CheckResult   string `json:"checkResult"`
}

func (t *MrchAuthStateQueryResponse) SuccessOrFail() bool {
	return t.RetCode == "000000"
}
