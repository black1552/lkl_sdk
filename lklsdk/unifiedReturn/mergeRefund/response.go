package mergerefund

// ResponseMergeRefund 统一退货响应结构体
type ResponseMergeRefund struct {
	Code     string                   `json:"code" dc:"Code 响应码"`
	Msg      string                   `json:"msg" dc:"Msg 响应描述"`
	RespTime string                   `json:"resp_time" dc:"RespTime 响应时间"`
	RespData *ResponseDataMergeRefund `json:"resp_data" dc:"RespData 响应数据"`
}

func (r *ResponseMergeRefund) SuccessOrFail() bool {
	return r.Code == "000000"
}

// ResponseDataMergeRefund 合单退货响应数据结构体
type ResponseDataMergeRefund struct {
	TradeState       string                   `json:"trade_state" dc:"TradeState 交易状态 必填，INIT-初始化；SUCCESS-交易成功；FAIL-交易失败；DEAL-交易处理中/未知；PROCESSING-交易已受理；TIMEOUT-超时未知；EXCEPTION-异常"`
	RefundType       string                   `json:"refund_type" dc:"RefundType 退货模式 必填"`
	MerchantNo       string                   `json:"merchant_no" dc:"MerchantNo 商户号 必填，拉卡拉分配的商户号"`
	OutTradeNo       string                   `json:"out_trade_no" dc:"OutTradeNo 商户请求流水号 必填，请求中的商户请求流水号"`
	TradeNo          string                   `json:"trade_no" dc:"TradeNo 拉卡拉交易流水号 必填"`
	LogNo            string                   `json:"log_no" dc:"LogNo 拉卡拉对账单流水号 必填，tradeNo的后14位"`
	AccTradeNo       string                   `json:"acc_trade_no,omitempty" dc:"AccTradeNo 账户端交易订单号 可选"`
	AccountType      string                   `json:"account_type,omitempty" dc:"AccountType 钱包类型 可选，微信：WECHAT；支付宝：ALIPAY；银联：UQRCODEPAY；翼支付:BESTPAY；苏宁易付宝:SUNING"`
	TotalAmount      string                   `json:"total_amount" dc:"TotalAmount 交易金额 必填，单位分，整数数字型字符串"`
	RefundAmount     string                   `json:"refund_amount" dc:"RefundAmount 申请退款金额 必填，单位分，整数数字型字符串"`
	PayerAmount      string                   `json:"payer_amount" dc:"PayerAmount 实际退款金额 必填，单位分，整数数字型字符串"`
	TradeTime        string                   `json:"trade_time,omitempty" dc:"TradeTime 退款时间 可选，实际退款时间。yyyyMMddHHmmss"`
	OriginLogNo      string                   `json:"origin_log_no,omitempty" dc:"OriginLogNo 原拉卡拉对账单流水号 可选，如果请求中携带，则返回"`
	OriginTradeNo    string                   `json:"origin_trade_no,omitempty" dc:"OriginTradeNo 原拉卡拉交易流水号 可选，如果请求中携带，则返回"`
	OriginOutTradeNo string                   `json:"origin_out_trade_no,omitempty" dc:"OriginOutTradeNo 原商户请求流水号 可选，如果请求中携带，则返回"`
	UpIssAddnData    string                   `json:"up_iss_addn_data,omitempty" dc:"UpIssAddnData 单品营销附加数据 可选，扫码交易，参与单品营销优惠时返回"`
	UpCouponInfo     string                   `json:"up_coupon_info,omitempty" dc:"UpCouponInfo 银联优惠信息 可选，扫码交易，参与单品营销优惠时返回"`
	TradeInfo        string                   `json:"trade_info,omitempty" dc:"TradeInfo 出资方信息 可选，扫码交易"`
	ChannelRetDesc   string                   `json:"channel_ret_desc" dc:"ChannelRetDesc 返回描述信息 必填"`
	RefundSplitInfo  []*RelateOutSplitRspInfo `json:"refund_split_info" dc:"RefundSplitInfo 响应合单域 必填"`
}

// RelateOutSplitRspInfo 响应合单域结构体
type RelateOutSplitRspInfo struct {
	OutSubTradeNo       string `json:"out_sub_trade_no" dc:"OutSubTradeNo 外部子退款交易流水号 必填，请求中的商户子流水号"`
	MerchantNo          string `json:"merchant_no" dc:"MerchantNo 商户号 必填，拉卡拉分配的商户号"`
	TermNo              string `json:"term_no" dc:"TermNo 终端号 必填，拉卡拉分配的业务终端号"`
	TradeState          string `json:"trade_state" dc:"TradeState 交易状态 必填，INIT-初始化；SUCCESS-交易成功；FAIL-交易失败；DEAL-交易处理中/未知；PROCESSING-交易已受理；TIMEOUT-超时未知；EXCEPTION-异常"`
	SubTradeNo          string `json:"sub_trade_no" dc:"SubTradeNo 拉卡拉子交易流水号 必填"`
	SubLogNo            string `json:"sub_log_no" dc:"SubLogNo 拉卡拉子对账单流水号 必填"`
	RefundAmount        string `json:"refund_amount" dc:"RefundAmount 申请退款金额 必填，单位分，整数型字符"`
	PayerAmount         string `json:"payer_amount" dc:"PayerAmount 实际退款金额 必填，单位分，整数型字符"`
	TotalAmount         string `json:"total_amount" dc:"TotalAmount 交易金额 必填，单位分，整数型字符"`
	OriginSubLogNo      string `json:"origin_sub_log_no,omitempty" dc:"OriginSubLogNo 原拉卡拉子对账单流水号 可选，如果请求中携带，则返回"`
	OriginSubTradeNo    string `json:"origin_sub_trade_no,omitempty" dc:"OriginSubTradeNo 原拉卡拉子交易流水号 可选，如果请求中携带，则返回"`
	OriginOutSubTradeNo string `json:"origin_out_sub_trade_no,omitempty" dc:"OriginOutSubTradeNo 原商户子交易流水号 可选，如果请求中携带，则返回"`
}
