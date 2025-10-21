package transMicropayEncry

// TransMicropayEncryResponseData 聚合扫码被扫接口（全报文加密）响应数据
type TransMicropayEncryResponseData struct {
	MerchantNo  string `json:"merchant_no"`  // 商户号
	TermNo      string `json:"term_no"`      // 终端号
	OutTradeNo  string `json:"out_trade_no"` // 商户交易流水号
	TotalAmount string `json:"total_amount"` // 金额，单位分
	PayOrderNo  string `json:"pay_order_no"` // 拉卡拉订单系统订单号
	TradeType   string `json:"trade_type"`   // 交易类型
	Status      string `json:"status"`       // 交易状态
	PayTime     string `json:"pay_time"`     // 支付完成时间
}

// TransMicropayEncryResponse 聚合扫码被扫接口（全报文加密）响应结构
type TransMicropayEncryResponse struct {
	RespTime string                          `json:"resp_time"` // 响应时间
	Version  string                          `json:"version"`   // 版本号
	Code     string                          `json:"code"`      // 响应码
	Msg      string                          `json:"msg"`       // 响应信息
	RespData *TransMicropayEncryResponseData `json:"resp_data"` // 响应数据
}

// SuccessOrFail 判断响应是否成功
func (r *TransMicropayEncryResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
