package model

// TradeQuery 交易查询请求结构体
type TradeQuery struct {
	ReqTime    string             `json:"req_time"`     // 请求时间
	Version    string             `json:"version"`      // API版本号
	OutOrgCode string             `json:"out_org_code"` // 外部机构码
	ReqData    *TradeQueryReqData `json:"req_data"`     // 请求数据
}

// TradeQueryReqData 交易查询请求数据结构体
type TradeQueryReqData struct {
	MerchantNo string `json:"merchant_no"`  // 商户号，必传
	TermNo     string `json:"term_no"`      // 终端号，必传
	OutTradeNo string `json:"out_trade_no"` // 商户交易流水号，条件必传，与trade_no必传其一
}

// TradeQueryResponse 交易查询响应结构体
type TradeQueryResponse struct {
	Msg      string `json:"msg"`       // 响应消息
	RespTime string `json:"resp_time"` // 响应时间
	Code     string `json:"code"`      // 响应码，000000表示成功
	RespData struct {
		MerchantNo    string `json:"merchant_no"`     // 商户号，必传
		OutTradeNo    string `json:"out_trade_no"`    // 商户请求流水号，必传
		TradeNo       string `json:"trade_no"`        // 拉卡拉商户订单号，必传
		LogNo         string `json:"log_no"`          // 拉卡拉对账流水号，必传
		TradeMainType string `json:"trade_main_type"` // 交易大类，条件必传（PREORDER-主扫，MICROPAY-被扫，REFUND-退款，CANCEL-撤销）
		SplitAttr     string `json:"split_attr"`      // 拆单属性，条件必传（M-主单，S-子单）
		SplitInfo     []struct {
			SubTradeNo    string `json:"sub_trade_no"`     // 子单交易流水号，必传
			SubLogNo      string `json:"sub_log_no"`       // 子单对账单单流水号，必传
			OutSubTradeNo string `json:"out_sub_trade_no"` // 外部子交易流水号，必传
			MerchantNo    string `json:"merchant_no"`      // 商户号，必传
			MerchantName  string `json:"merchant_name"`    // 商户名称，必传
			TermNo        string `json:"term_no"`          // 终端号，必传
			Amount        string `json:"amount"`           // 金额，必传（单位分）
		} `json:"split_info"` // 拆单信息，条件必传（如果查询订单是主单，则返回）
		RefundSplitInfo []struct {
			OutSubTradeNo string `json:"out_sub_trade_no"` // 外部子退款交易流水号，必传
			MerchantNo    string `json:"merchant_no"`      // 商户号，必传
			TermNo        string `json:"term_no"`          // 终端号，必传
			RefundAmount  string `json:"refund_amount"`    // 申请退款金额，必传（单位分）
			SubTradeNo    string `json:"sub_trade_no"`     // 拉卡分子交易流水号，条件必传
			SubLogNo      string `json:"sub_log_no"`       // 对账单子流水号，条件必传
			TradeState    string `json:"trade_state"`      // 子交易状态，条件必传（SUCCESS-交易成功 FAIL-交易失败）
			ResultCode    string `json:"result_code"`      // 处理结果码，条件必传
			ResultMsg     string `json:"result_msg"`       // 处理描述，条件必传
		} `json:"refund_split_info"` // 合单退款拆单信息，条件必传（如果查询订单是退款主单，则返回）
		AccTradeNo             string                 `json:"acc_trade_no"`              // 账户端交易订单号，必传
		AccountType            string                 `json:"account_type"`              // 钱包类型，必传（微信: WECHAT 支付宝: ALIPAY 银联: UQRCODEPAY 翼支付: BESTPAY 苏宁支付: SUNING）
		TradeState             string                 `json:"trade_state"`               // 交易状态，必传（INIT-初始化 CREATE-下单成功 SUCCESS-交易成功 FAIL-交易失败 DEAL-交易处理中 UNKNOWN-未知状态 CLOSE-订单关闭 PART_REFUND-部分退款 REFUND-全部退款）
		TradeStateDesc         string                 `json:"trade_state_desc"`          // 交易状态描述，条件必传
		TotalAmount            string                 `json:"total_amount"`              // 订单金额，必传（单位分）
		PayerAmount            string                 `json:"payer_amount"`              // 付款人实付金额，条件必传（单位分）
		AccSettleAmount        string                 `json:"acc_settle_amount"`         // 账户端结算金额，条件必传（单位分）
		AccMdiscountAmount     string                 `json:"acc_mdiscount_amount"`      // 商户侧优惠金额，条件必传（单位分）
		AccDiscountAmount      string                 `json:"acc_discount_amount"`       // 账户端优惠金额，条件必传（单位分）
		AccOtherDiscountAmount string                 `json:"acc_other_discount_amount"` // 账户端其它优惠金额，条件必传（单位分）
		TradeTime              string                 `json:"trade_time"`                // 交易完成时间，条件必传（yyyyMMddHHmmss）
		UserId1                string                 `json:"user_id1"`                  // 用户标识1，条件必传（微信sub_open_id 支付宝buyer_login_id 买家支付账号）
		UserId2                string                 `json:"user_id2"`                  // 用户标识2，条件必传（微信open_id 支付宝buyer_user_id 银user_id）
		BankType               string                 `json:"bank_type"`                 // 付款银行，条件必传
		CardType               string                 `json:"card_type"`                 // 银行卡类型，条件必传（00: 借记卡 01: 贷记卡 02: 微信零钱 03: 支付宝花呗 04: 支付宝其他 05: 数字货币 06: 拉卡拉支付账户 99: 未知）
		AccActivityId          string                 `json:"acc_activity_id"`           // 活动ID，条件必传（在账户端商户后台配置的批次ID）
		TradeReqDate           string                 `json:"trade_req_date"`            // 交易请求日期，必传（yyyyMMdd）
		AccRespFields          map[string]interface{} `json:"acc_resp_fields"`           // 账户端返回信息域，条件必传
	} `json:"resp_data"` // 响应数据
}

func (t *TradeQueryResponse) SuccessOrFail() bool {
	return t.Code == "BBS00000"
}
