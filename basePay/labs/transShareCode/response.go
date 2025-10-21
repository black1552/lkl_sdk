package transShareCode

// AccRespFields 账户端返回信息域
type AccRespFields struct {
	ShareToken string `json:"share_token"` // 吱口令，必填
	ExpireDate string `json:"expire_date"` // 吱口令失效时间，可选，若为空则表示永久有效
}

// TransShareCodeResponseData 响应数据
type TransShareCodeResponseData struct {
	MerchantNo    string         `json:"merchant_no"`     // 商户号
	OutTradeNo    string         `json:"out_trade_no"`    // 商户请求流水号
	TradeNo       string         `json:"trade_no"`        // 拉卡拉交易流水号
	AccRespFields *AccRespFields `json:"acc_resp_fields"` // 账户端返回信息域
}

// TransShareCodeResponse 整体响应结构
type TransShareCodeResponse struct {
	Code     string                      `json:"code"`      // 响应码
	Msg      string                      `json:"msg"`       // 响应信息
	RespTime string                      `json:"resp_time"` // 响应时间
	RespData *TransShareCodeResponseData `json:"resp_data"` // 响应数据
}

// SuccessOrFail 判断请求是否成功
func (r *TransShareCodeResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
