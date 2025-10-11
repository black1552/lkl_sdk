package refundquery

// RequestRefundQuery 统一退货查询请求结构体
// API文档: https://o.lakala.com/#/home/document/detail?id=893
type RequestRefundQuery struct {
	ReqTime string                  `json:"req_time" dc:"ReqTime 请求时间 格式yyyyMMddHHmmss"`
	Version string                  `json:"version" dc:"Version 版本号 固定值3.0"`
	ReqData *RequestDataRefundQuery `json:"req_data" dc:"ReqData 请求数据"`
}

// RequestDataRefundQuery 统一退货查询请求数据结构体
// API文档: https://o.lakala.com/#/home/document/detail?id=893
type RequestDataRefundQuery struct {
	MerchantNo string `json:"merchant_no" dc:"MerchantNo 商户号 必填，拉卡拉分配的商户号，String(15)"`
	TermNo     string `json:"term_no" dc:"TermNo 终端号 必填，拉卡拉分配的业务终端号，String(8)"`
	OutTradeNo string `json:"out_trade_no,omitempty" dc:"OutTradeNo 商户交易流水号 选填，下单时的商户请求流水号，String(32) 与拉卡拉交易流水号二选一"`
	TradeNo    string `json:"trade_no,omitempty" dc:"TradeNo 拉卡拉交易流水号 选填，拉卡拉交易流水号，String(32) 与商户交易流水号二选一"`
}
