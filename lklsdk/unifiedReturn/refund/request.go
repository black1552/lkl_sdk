package refund

// RequestRefund 统一退货基础请求结构体
// 参考文档：https://o.lakala.com/#/home/document/detail?id=894
type RequestRefund struct {
	ReqTime string             `json:"req_time"`
	Version string             `json:"version"`
	ReqData *RequestDataRefund `json:"req_data"`
}

// RequestDataRefund 统一退货请求数据结构体
// 注意：origin_out_trade_no、origin_log_no、origin_trade_no至少一个必填（调用收银台下单接口拉起交易后发起退款时至少要传两个）
// 优先级顺序：origin_trade_no > origin_log_no > origin_out_trade_no
type RequestDataRefund struct {
	MerchantNo       string        `json:"merchant_no"`                   // merchant_no: 商户号 (必填，拉卡拉分配的商户号，String(15))
	TermNo           string        `json:"term_no"`                       // term_no: 终端号 (必填，拉卡拉分配的终端号，String(8))
	OutTradeNo       string        `json:"out_trade_no"`                  // out_trade_no: 商户请求流水号 (必填，商户系统唯一，String(32))
	RefundAmount     string        `json:"refund_amount"`                 // refund_amount: 退款金额 (必填，单位分，整数数字型字符，String(12))
	RefundReason     string        `json:"refund_reason,omitempty"`       // refund_reason: 退货原因 (可选，String(32))
	OriginLogNo      string        `json:"origin_log_no,omitempty"`       // origin_log_no: 拉卡拉对账单流水号 (可选，正交易返回的拉卡拉对账单流水号，String(14))
	OriginOutTradeNo string        `json:"origin_out_trade_no,omitempty"` // origin_out_trade_no: 原商户交易流水号 (可选，String(32))
	OriginTradeNo    string        `json:"origin_trade_no,omitempty"`     // origin_trade_no: 原交易拉卡拉交易订单号 (可选，String(32))
	LocationInfo     *LocationInfo `json:"location_info"`                 // location_info: 地址位置信息 (必填)
	RefundAccMode    string        `json:"refund_acc_mode,omitempty"`     // refund_acc_mode: 退货账户模式 (可选，00退货账户余额 05商户余额 06终端余额 30终点账户，String(2))
	NotifyURL        string        `json:"notify_url,omitempty"`          // notify_url: 后台通知地址 (可选，交易结果通知地址，String(128))
	RefundAmtSts     string        `json:"refund_amt_sts,omitempty"`      // refund_amt_sts: 退货资金状态 (可选，00 分账前，01 分账后；分账交易部分退货的情况，需要前端上送交易的分账状态，String(2))
}

// LocationInfo 地址位置信息结构体
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`             // request_ip: 请求IP (必填)
	Location    string `json:"location"`               // location: 位置信息 (必填，如经纬度等)
	BaseStation string `json:"base_station,omitempty"` // base_station: 基站信息 (可选)
}
