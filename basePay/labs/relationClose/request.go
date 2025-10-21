package relationclose

// LocationInfo 地址位置信息结构体
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`             // 请求方IP地址 (必填)
	BaseStation string `json:"base_station,omitempty"` // 基站信息 (可选)
	Location    string `json:"location,omitempty"`     // 纬度,经度 (可选)
}

// RelationCloseRequestData 聚合扫码关单请求数据结构体
type RelationCloseRequestData struct {
	MerchantNo           string        `json:"merchant_no" dc:"商户号 (必填)"`                             // 商户号 (必填)
	TermNo               string        `json:"term_no" dc:"终端号 (必填)"`                                 // 终端号 (必填)
	OriginOutTradeNo     string        `json:"origin_out_trade_no,omitempty" dc:"原商户交易流水号 (可选)"`      // 原商户交易流水号 (可选)
	OriginTradeNo        string        `json:"origin_trade_no,omitempty" dc:"原交易拉卡拉交易流水号 (可选)"`       // 原交易拉卡拉交易流水号 (可选)
	OriginOutOrderSource string        `json:"origin_out_order_source,omitempty" dc:"原订单外部订单来源 (可选)"` // 原订单外部订单来源 (可选)
	OriginOutOrderNo     string        `json:"origin_out_order_no,omitempty" dc:"原订单外部商户订单号 (可选)"`    // 原订单外部商户订单号 (可选)
	LocationInfo         *LocationInfo `json:"location_info" dc:"地址位置信息 (必填)"`                        // 地址位置信息 (必填)
}

// RelationCloseRequest 聚合扫码关单请求结构体
type RelationCloseRequest struct {
	ReqTime string                    `json:"req_time"`
	Version string                    `json:"version"`
	ReqData *RelationCloseRequestData `json:"req_data"`
}
