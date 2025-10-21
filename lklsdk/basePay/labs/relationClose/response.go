package relationclose

// RelationCloseResponseData 聚合扫码关单响应数据结构体
type RelationCloseResponseData struct {
	OriginTradeNo        string `json:"origin_trade_no"`                   // 原拉卡拉订单号
	OriginOutTradeNo     string `json:"origin_out_trade_no"`               // 原商户请求流水号
	OriginOutOrderSource string `json:"origin_out_order_source,omitempty"` // 原订单外部订单来源
	OriginOutOrderNo     string `json:"origin_out_order_no,omitempty"`     // 原订单外部商户订单号
	TradeTime            string `json:"trade_time"`                        // 交易时间
}

// RelationCloseResponse 聚合扫码关单响应结构体
type RelationCloseResponse struct {
	RespTime string                     `json:"resp_time"`
	Code     string                     `json:"code"`
	Msg      string                     `json:"msg"`
	RespData *RelationCloseResponseData `json:"resp_data,omitempty"`
}

// SuccessOrFail 判断响应是否成功
func (r *RelationCloseResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
