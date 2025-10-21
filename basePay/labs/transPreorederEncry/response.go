package transPreorederEncry

// TransPreorederEncryResponseData 聚合扫码主扫交易（全报文加密）响应数据
type TransPreorederEncryResponseData struct {
	MerchantNo  string `json:"merchant_no"`  // 商户号
	TermNo      string `json:"term_no"`      // 终端号
	OutTradeNo  string `json:"out_trade_no"` // 商户交易流水号
	TotalAmount string `json:"total_amount"` // 金额，单位分
	PayOrderNo  string `json:"pay_order_no"` // 拉卡拉订单系统订单号
	CodeUrl     string `json:"code_url"`     // 二维码内容，用于生成支付二维码
	TradeType   string `json:"trade_type"`   // 交易类型
}

// TransPreorederEncryResponse 聚合扫码主扫交易（全报文加密）响应结构
type TransPreorederEncryResponse struct {
	RespTime string                          `json:"resp_time"` // 响应时间
	Version  string                          `json:"version"`   // 版本号
	Code     string                          `json:"code"`      // 响应码
	Msg      string                          `json:"msg"`       // 响应信息
	RespData TransPreorederEncryResponseData `json:"resp_data"` // 响应数据
}

// SuccessOrFail 判断响应是否成功
func (r *TransPreorederEncryResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
