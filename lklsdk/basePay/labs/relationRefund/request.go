package relationRefund

// LocationInfo 地址位置信息结构体
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`             // 请求方IP地址，必填
	BaseStation string `json:"base_station,omitempty"` // 基站信息，可选
	Location    string `json:"location,omitempty"`     // 维度,经度，可选
}

// RelationRefundRequestData 聚合扫码退款请求数据结构体
type RelationRefundRequestData struct {
	MerchantNo       string       `json:"merchant_no"`                   // 商户号，必填
	TermNo           string       `json:"term_no"`                       // 终端号，必填
	OutTradeNo       string       `json:"out_trade_no"`                  // 商户交易流水号，必填
	RefundAmount     string       `json:"refund_amount"`                 // 退款金额，必填
	RefundReason     string       `json:"refund_reason"`                 // 退款原因，必填
	OriginOutTradeNo string       `json:"origin_out_trade_no,omitempty"` // 原商户交易流水号，可选（退款时origin_out_trade_no，origin_trade_no，origin_log_no必送其一）
	OriginTradeNo    string       `json:"origin_trade_no,omitempty"`     // 原拉卡拉交易流水号，可选
	OriginLogNo      string       `json:"origin_log_no,omitempty"`       // 原对账单流水号，可选
	TradeReqDate     string       `json:"trade_req_date,omitempty"`      // 交易请求日期，可选（送原商户交易流水号退款时必填）
	LocationInfo     LocationInfo `json:"location_info"`                 // 地址位置信息，必填
}

// RelationRefundRequest 聚合扫码退款请求结构体
type RelationRefundRequest struct {
	ReqTime string                     `json:"req_time"`
	Version string                     `json:"version"`
	ReqData *RelationRefundRequestData `json:"req_data"`
}
