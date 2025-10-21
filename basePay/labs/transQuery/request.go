package transquery

// TransQueryRequestData 聚合扫码交易查询请求数据结构体
type TransQueryRequestData struct {
	MerchantNo string `json:"merchant_no"`            // 商户号 (必填)
	TermNo     string `json:"term_no"`                // 终端号 (必填)
	OutTradeNo string `json:"out_trade_no,omitempty"` // 商户交易流水号 (可选，与trade_no二选一)
	TradeNo    string `json:"trade_no,omitempty"`     // 拉卡拉交易流水号 (可选，与out_trade_no二选一)
}

// TransQueryRequest 聚合扫码交易查询请求结构体
type TransQueryRequest struct {
	ReqTime string                 `json:"req_time"`
	Version string                 `json:"version"`
	ReqData *TransQueryRequestData `json:"req_data"`
}
