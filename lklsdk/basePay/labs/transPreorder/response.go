package transPreorder

// ResponseData 聚合主扫响应数据结构体
type ResponseData struct {
	MerchantNo     string `json:"merchant_no"`                // 商户号
	TermNo         string `json:"term_no"`                    // 终端号
	OutTradeNo     string `json:"out_trade_no"`               // 商户交易流水号
	TradeNo        string `json:"trade_no"`                   // 拉卡拉交易流水号
	LogNo          string `json:"log_no"`                     // 拉卡拉对账单流水号
	TotalAmount    string `json:"total_amount"`               // 交易总金额
	TradeTime      string `json:"trade_time"`                 // 交易时间
	PayStatus      string `json:"pay_status"`                 // 支付状态
	OrderDesc      string `json:"order_desc,omitempty"`       // 订单描述
	PayAmount      string `json:"pay_amount,omitempty"`       // 实际支付金额
	PayChannel     string `json:"pay_channel,omitempty"`      // 支付渠道
	ChannelTradeNo string `json:"channel_trade_no,omitempty"` // 渠道交易号
	QrCode         string `json:"qr_code,omitempty"`          // 支付二维码内容 (NATIVE支付时返回)
	CodeURL        string `json:"code_url,omitempty"`         // 二维码链接 (NATIVE支付时返回)
	PayParams      string `json:"pay_params,omitempty"`       // 支付参数 (JSAPI/APP支付时返回)
	AppId          string `json:"app_id,omitempty"`           // 应用ID
	MchId          string `json:"mch_id,omitempty"`           // 商户号
	PrepayId       string `json:"prepay_id,omitempty"`        // 预支付交易会话ID
}

// Response 预订单响应结构体
type TransPreorderResponse struct {
	RespTime string        `json:"resp_time"`
	Code     string        `json:"code"`
	Msg      string        `json:"msg"`
	RespData *ResponseData `json:"resp_data,omitempty"`
}

// SuccessOrFail 判断响应是否成功
func (r *TransPreorderResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
