package transquery

// SplitInfo 拆单信息结构体
type SplitInfo struct {
	SubTradeNo    string `json:"sub_trade_no"`     // 子单交易流水号
	SubLogNo      string `json:"sub_log_no"`       // 子单对账单流水号
	OutSubTradeNo string `json:"out_sub_trade_no"` // 外部子交易流水号
	MerchantNo    string `json:"merchant_no"`      // 商户号
	MerchantName  string `json:"merchant_name"`    // 商户名称
	TermNo        string `json:"term_no"`          // 终端号
	Amount        string `json:"amount"`           // 金额
}

// RefundSplitInfo 合单退款拆单信息结构体
type RefundSplitInfo struct {
	OutSubTradeNo string `json:"out_sub_trade_no"`       // 外部子退款交易流水号
	MerchantNo    string `json:"merchant_no"`            // 商户号
	TermNo        string `json:"term_no"`                // 终端号
	RefundAmount  string `json:"refund_amount"`          // 申请退款金额
	SubTradeNo    string `json:"sub_trade_no,omitempty"` // 拉卡拉子交易流水号
	SubLogNo      string `json:"sub_log_no,omitempty"`   // 对账单子流水号
	TradeState    string `json:"trade_state,omitempty"`  // 子交易状态
	ResultCode    string `json:"result_code,omitempty"`  // 处理结果码
	ResultMsg     string `json:"result_msg,omitempty"`   // 处理描述
}

// AccRespFields 账户端返回信息域结构体
type AccRespFields struct {
	UserId            string `json:"user_id,omitempty"`             // 买家在支付宝的用户id
	StoreId           string `json:"store_id,omitempty"`            // 商户门店编号
	AlipayStoreId     string `json:"alipay_store_id,omitempty"`     // 支付宝店铺编号（不再使用）
	FundBillList      string `json:"fund_bill_list,omitempty"`      // 交易支付使用的资金渠道
	VoucherDetailList string `json:"voucher_detail_list,omitempty"` // 所有优惠券信息
	HbFqPayInfo       string `json:"hb_fq_pay_info,omitempty"`      // 花呗分期支付信息
}

// TransQueryResponseData 聚合扫码交易查询响应数据结构体
type TransQueryResponseData struct {
	MerchantNo             string             `json:"merchant_no"`                         // 拉卡拉分配的商户号
	OutTradeNo             string             `json:"out_trade_no"`                        // 商户请求流水号
	TradeNo                string             `json:"trade_no"`                            // 拉卡拉商户订单号
	LogNo                  string             `json:"log_no"`                              // 拉卡拉对账单流水号
	TradeMainType          string             `json:"trade_main_type,omitempty"`           // 交易大类
	SplitAttr              string             `json:"split_attr,omitempty"`                // 拆单属性
	SplitInfo              []*SplitInfo       `json:"split_info,omitempty"`                // 拆单信息
	AccTradeNo             string             `json:"acc_trade_no"`                        // 账户端交易订单号
	AccountType            string             `json:"account_type"`                        // 钱包类型
	TradeState             string             `json:"trade_state"`                         // 交易状态
	TradeStateDesc         string             `json:"trade_state_desc,omitempty"`          // 交易状态描述
	TotalAmount            string             `json:"total_amount"`                        // 订单金额
	PayerAmount            string             `json:"payer_amount,omitempty"`              // 付款人实付金额
	AccSettleAmount        string             `json:"acc_settle_amount,omitempty"`         // 账户端结算金额
	AccMDiscountAmount     string             `json:"acc_mdiscount_amount,omitempty"`      // 商户侧优惠金额（账户端）
	AccDiscountAmount      string             `json:"acc_discount_amount,omitempty"`       // 账户端优惠金额
	AccOtherDiscountAmount string             `json:"acc_other_discount_amount,omitempty"` // 账户端其它优惠金额
	TradeTime              string             `json:"trade_time"`                          // 交易完成时间
	UserId1                string             `json:"user_id1,omitempty"`                  // 用户标识1
	UserId2                string             `json:"user_id2,omitempty"`                  // 用户标识2
	BankType               string             `json:"bank_type,omitempty"`                 // 付款银行
	CardType               string             `json:"card_type,omitempty"`                 // 银行卡类型
	AccActivityId          string             `json:"acc_activity_id,omitempty"`           // 活动ID
	TradeReqDate           string             `json:"trade_req_date"`                      // 交易请求日期
	AccRespFields          *AccRespFields     `json:"acc_resp_fields,omitempty"`           // 账户端返回信息域
	RefundSplitInfo        []*RefundSplitInfo `json:"refund_split_info,omitempty"`         // 合单退款拆单信息
}

// TransQueryResponse 聚合扫码交易查询响应结构体
type TransQueryResponse struct {
	RespTime string                  `json:"resp_time"`
	Code     string                  `json:"code"`
	Msg      string                  `json:"msg"`
	RespData *TransQueryResponseData `json:"resp_data,omitempty"`
}

// SuccessOrFail 判断响应是否成功
func (r *TransQueryResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
