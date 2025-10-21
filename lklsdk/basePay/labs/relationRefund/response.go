package relationRefund

// RelationRefundResponseData 聚合扫码退款响应数据结构体
type RelationRefundResponseData struct {
	MerchantNo       string `json:"merchant_no"`                   // 商户号
	OutTradeNo       string `json:"out_trade_no"`                  // 商户请求流水号
	TradeNo          string `json:"trade_no"`                      // 拉卡拉退款单号
	LogNo            string `json:"log_no"`                        // 拉卡拉对账单流水号
	AccTradeNo       string `json:"acc_trade_no,omitempty"`        // 账户端交易订单号
	AccountType      string `json:"account_type,omitempty"`        // 钱包类型
	TotalAmount      string `json:"total_amount"`                  // 交易金额
	RefundAmount     string `json:"refund_amount"`                 // 申请退款金额
	PayerAmount      string `json:"payer_amount"`                  // 实际退款金额
	TradeTime        string `json:"trade_time,omitempty"`          // 退款时间
	OriginTradeNo    string `json:"origin_trade_no,omitempty"`     // 原拉卡拉订单号
	OriginOutTradeNo string `json:"origin_out_trade_no,omitempty"` // 原商户请求流水号
	UpIssAddnData    string `json:"up_iss_addn_data,omitempty"`    // 单品营销附加数据
	UpCouponInfo     string `json:"up_coupon_info,omitempty"`      // 银联优惠信息、出资方信息
	TradeInfo        string `json:"trade_info,omitempty"`          // 出资方信息
	FundBillList     string `json:"fund_bill_list,omitempty"`      // 交易支付使用的资金渠道
}

// RelationRefundResponse 聚合扫码退款响应结构体
type RelationRefundResponse struct {
	Code     string                      `json:"code"`
	Msg      string                      `json:"msg"`
	RespTime string                      `json:"resp_time"`
	RespData *RelationRefundResponseData `json:"resp_data,omitempty"`
}

// SuccessOrFail 判断响应是否成功
func (r *RelationRefundResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
