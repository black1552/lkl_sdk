package refundquery

// ResponseRefundQuery 统一退货查询响应结构体
// API文档: https://o.lakala.com/#/home/document/detail?id=893
type ResponseRefundQuery struct {
	Code string                   `json:"code" dc:"Code 响应码 RFD00000#成功、RFD11112#网络请求超时"`
	Msg  string                   `json:"msg" dc:"Msg 响应描述"`
	Data *ResponseDataRefundQuery `json:"data,omitempty" dc:"Data 响应数据"`
	Sign string                   `json:"sign" dc:"Sign 签名"`
}

// ResponseDataRefundQuery 统一退货查询响应数据结构体
// API文档: https://o.lakala.com/#/home/document/detail?id=893
type ResponseDataRefundQuery struct {
	OutTradeNo         string                   `json:"out_trade_no" dc:"OutTradeNo 商户请求流水号 必填，原退款交易商户请求流水号（扫码交易返回）"`
	TradeTime          string                   `json:"trade_time" dc:"TradeTime 交易时间 必填，交易时间yyyyMMddHHmmss"`
	TradeState         string                   `json:"trade_state" dc:"TradeState 交易状态 必填，INIT-初始化；SUCCESS-交易成功；FAIL-交易失败；DEAL-交易处理中/未知；TIMEOUT-超时未知；EXCEPTION-异常"`
	TradeNo            string                   `json:"trade_no" dc:"TradeNo 拉卡拉商户订单号 必填，拉卡拉生成的交易流水"`
	LogNo              string                   `json:"log_no" dc:"LogNo 拉卡拉对账单流水号 必填，拉卡拉生成的对账单流水号（新增）"`
	AccTradeNo         string                   `json:"acc_trade_no,omitempty" dc:"AccTradeNo 账户端交易订单号 选填，账户端交易流水号，String(32)"`
	RefundAmount       string                   `json:"refund_amount" dc:"RefundAmount 交易金额 必填"`
	PayMode            string                   `json:"pay_mode,omitempty" dc:"PayMode 支付方式 选填，00 借记卡 01 贷记卡 02 准贷记卡 银行卡交易返回"`
	CrdNo              string                   `json:"crd_no,omitempty" dc:"CrdNo 卡号 选填，脱敏卡号，前六后四，中间用*替换"`
	AccountType        string                   `json:"account_type,omitempty" dc:"AccountType 钱包类型 选填，微信：WECHAT 支付宝：ALIPAY 银联：UNION 翼支付: BESTPAY 苏宁易付宝: SUNING 扫码交易返回"`
	PayerAmount        string                   `json:"payer_amount,omitempty" dc:"PayerAmount 付款人实付金额 选填，实际退款金额，单位分 扫码交易返回"`
	AccSettleAmount    string                   `json:"acc_settle_amount,omitempty" dc:"AccSettleAmount 账户端结算金额 选填，账户端应结订单金额，单位分 扫码交易返回"`
	AccMdiscountAmount string                   `json:"acc_mdiscount_amount,omitempty" dc:"AccMdiscountAmount 商户侧优惠金额（账户端） 选填，商户优惠金额，单位分 扫码交易返回"`
	AccDiscountAmount  string                   `json:"acc_discount_amount,omitempty" dc:"AccDiscountAmount 账户端优惠金额 选填，拉卡拉优惠金额， 扫码交易返回"`
	ChannelRetDesc     string                   `json:"channel_ret_desc" dc:"ChannelRetDesc 返回描述信息 必填，code#msg：RFD00000#成功、RFD11112#网络请求超时"`
	RefundSplitInfo    []*RelateOutSplitRspInfo `json:"refund_split_info,omitempty" dc:"RefundSplitInfo 退款拆单信息 选填，合单交易退款查询时返回"`
	OriginLogNo        string                   `json:"origin_log_no" dc:"OriginLogNo 拉卡拉对账单流水号 必填，原交易拉卡拉对账单流水号"`
	OriginOutTradeNo   string                   `json:"origin_out_trade_no" dc:"OriginOutTradeNo 原商户交易流水号 必填"`
	OriginTradeNo      string                   `json:"origin_trade_no" dc:"OriginTradeNo 原交易拉卡拉交易订单号 必填"`
	OriginTotalAmount  string                   `json:"origin_total_amount" dc:"OriginTotalAmount 原交易金额 必填，原正交易订单金额"`
	RefundType         string                   `json:"refund_type,omitempty" dc:"RefundType 退货模式 选填"`
}

// RelateOutSplitRspInfo 退款拆单信息结构体
// API文档: https://o.lakala.com/#/home/document/detail?id=893
type RelateOutSplitRspInfo struct {
	OutSubTradeNo string `json:"out_sub_trade_no" dc:"OutSubTradeNo 外部子退款交易流水号 必填，商户子交易流水号，商户号下唯一"`
	MerchantNo    string `json:"merchant_no" dc:"MerchantNo 商户号 必填，拉卡拉分配的商户号"`
	TermNo        string `json:"term_no" dc:"TermNo 终端号 必填，拉卡拉分配的业务终端号"`
	RefundAmount  string `json:"refund_amount" dc:"RefundAmount 申请退款金额 必填，单位分，整数型字符"`
	SubTradeNo    string `json:"sub_trade_no,omitempty" dc:"SubTradeNo 拉卡拉子交易流水号 选填"`
	SubLogNo      string `json:"sub_log_no,omitempty" dc:"SubLogNo 对账单子流水号 选填，sub_trade_no后14位"`
	TradeState    string `json:"trade_state,omitempty" dc:"TradeState 子交易状态 选填，SUCCESS-交易成功 FAIL-交易失败"`
	ResultCode    string `json:"result_code,omitempty" dc:"ResultCode 处理结果码 选填"`
	ResultMsg     string `json:"result_msg,omitempty" dc:"ResultMsg 处理描述 选填"`
}

// SuccessOrFail 判断交易是否成功
// 返回true表示交易成功，false表示交易失败或处理中
func (r *ResponseRefundQuery) SuccessOrFail() bool {
	return r.Code == "000000"
}
