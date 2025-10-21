package relationRevoked

// LocationInfo 地址位置信息结构体
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`             // 请求方IP地址，必填
	BaseStation string `json:"base_station,omitempty"` // 基站信息，可选
	Location    string `json:"location,omitempty"`     // 维度,经度，可选
}

// RelationRevokedRequestData 聚合扫码撤销请求数据结构体
type RelationRevokedRequestData struct {
	MerchantNo       string       `json:"merchant_no"`                   // 商户号，必填
	TermNo           string       `json:"term_no"`                       // 终端号，必填
	OutTradeNo       string       `json:"out_trade_no"`                  // 商户交易流水号，必填
	OriginOutTradeNo string       `json:"origin_out_trade_no,omitempty"` // 原商户交易流水号，可选（撤销时origin_out_trade_no，origin_trade_no必送其一）
	OriginTradeNo    string       `json:"origin_trade_no,omitempty"`     // 原拉卡拉交易流水号，可选（撤销时origin_out_trade_no，origin_trade_no必送其一）
	LocationInfo     LocationInfo `json:"location_info"`                 // 地址位置信息，必填
}

// RelationRevokedRequest 聚合扫码关单请求结构体
type RelationRevokedRequest struct {
	ReqTime string                      `json:"req_time"`
	Version string                      `json:"version"`
	ReqData *RelationRevokedRequestData `json:"req_data"`
}
