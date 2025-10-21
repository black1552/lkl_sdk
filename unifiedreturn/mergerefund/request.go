package mergerefund

// RequestMergeRefund 统一退货基础请求结构体
type RequestMergeRefund struct {
	ReqTime string                  `json:"req_time" dc:"ReqTime 请求时间 必填，格式yyyyMMddHHmmss"`
	Version string                  `json:"version" dc:"Version 版本号 必填，固定为\"3\""`
	ReqData *RequestDataMergeRefund `json:"req_data" dc:"ReqData 请求参数 必填，具体字段见对应的数据结构体"`
}

// RequestDataMergeRefund 合单退货请求参数结构体
type RequestDataMergeRefund struct {
	MerchantNo       string                `json:"merchant_no" dc:"MerchantNo 商户号 必填，拉卡拉分配的商户号"`
	TermNo           string                `json:"term_no" dc:"TermNo 终端号 必填，拉卡拉分配的终端号"`
	OutTradeNo       string                `json:"out_trade_no" dc:"OutTradeNo 商户请求流水号 必填，商户系统唯一"`
	RefundAmount     string                `json:"refund_amount" dc:"RefundAmount 退款金额 必填，单位分，整数数字型字符"`
	RefundReason     string                `json:"refund_reason,omitempty" dc:"RefundReason 退货原因 可选"`
	OriginLogNo      string                `json:"origin_log_no,omitempty" dc:"OriginLogNo 拉卡拉对账单流水号 可选，正交易返回的拉卡拉对账单流水号"`
	OriginOutTradeNo string                `json:"origin_out_trade_no,omitempty" dc:"OriginOutTradeNo 原商户交易流水号 可选"`
	OriginTradeNo    string                `json:"origin_trade_no,omitempty" dc:"OriginTradeNo 原交易拉卡拉交易订单号 可选"`
	LocationInfo     *LocationInfo         `json:"location_info" dc:"LocationInfo 地址位置信息 必填"`
	NotifyURL        string                `json:"notify_url,omitempty" dc:"NotifyURL 后台通知地址 可选，交易结果通知地址"`
	RefundAccMode    string                `json:"refund_acc_mode,omitempty" dc:"RefundAccMode 退货账户模式 可选，00退货账户余额 05商户余额 06终端余额"`
	RefundSplitInfo  []*RelateOutSplitInfo `json:"refund_split_info,omitempty" dc:"RefundSplitInfo 请参合单域 可选"`
}

// LocationInfo 地址位置信息结构体
type LocationInfo struct {
	RequestIP   string `json:"request_ip" dc:"RequestIP 请求方IP地址 必填，格式如36.45.36.95"`
	Location    string `json:"location,omitempty" dc:"Location 维度,经度 可选，商户终端的地理位置，格式：纬度,经度，+表示北纬、东经，-表示南纬、西经，精度最长支持小数点后9位"`
	BaseStation string `json:"base_station,omitempty" dc:"BaseStation 基站信息 可选，客户端设备的基站信息"`
}

// RelateOutSplitInfo 请求合单域结构体
type RelateOutSplitInfo struct {
	OutSubTradeNo       string `json:"out_sub_trade_no" dc:"OutSubTradeNo 外部子退款交易流水号 必填，商户子交易退款流水号，商户号下唯一"`
	MerchantNo          string `json:"merchant_no" dc:"MerchantNo 商户号 必填，拉卡拉分配的商户号"`
	TermNo              string `json:"term_no" dc:"TermNo 终端号 必填，拉卡拉分配的业务终端号"`
	RefundAmount        string `json:"refund_amount" dc:"RefundAmount 申请退款金额 必填，单位分，整数型字符"`
	OriginOutSubTradeNo string `json:"origin_out_sub_trade_no,omitempty" dc:"OriginOutSubTradeNo 原商户交易流水号 可选，下单时的商户子单请求流水号"`
	OriginSubTradeNo    string `json:"origin_sub_trade_no,omitempty" dc:"OriginSubTradeNo 原拉卡拉子交易流水号 可选，下单成功时，返回的拉卡拉交易子流水"`
	OriginSubLogNo      string `json:"origin_sub_log_no,omitempty" dc:"OriginSubLogNo 原对账单子流水号 可选，原交易的tradeNo的后14位，必须是66开头的"`
}
