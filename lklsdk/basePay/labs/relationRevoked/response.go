package relationRevoked

// RelationRevokedResponseData 聚合扫码撤销响应数据结构体
type RelationRevokedResponseData struct {
	MerchantNo       string `json:"merchant_no"`                   // 商户号
	OutTradeNo       string `json:"out_trade_no"`                  // 商户请求流水号
	TradeNo          string `json:"trade_no,omitempty"`            // 拉卡拉商户订单号
	LogNo            string `json:"log_no"`                        // 拉卡拉对账单流水号
	AccTradeNo       string `json:"acc_trade_no,omitempty"`        // 账户端交易订单号
	AccountType      string `json:"account_type"`                  // 钱包类型
	TotalAmount      string `json:"total_amount"`                  // 交易金额
	TradeTime        string `json:"trade_time"`                    // 交易完成时间
	OriginOutTradeNo string `json:"origin_out_trade_no,omitempty"` // 原商户请求流水号
	OriginTradeNo    string `json:"origin_trade_no,omitempty"`     // 原拉卡拉交易流水号
}

// RelationRevokedResponse 聚合扫码撤销响应结构体
type RelationRevokedResponse struct {
	Code     string                      `json:"code"`                // 响应码
	Msg      string                      `json:"msg"`                 // 响应信息
	RespTime string                      `json:"resp_time"`           // 响应时间
	RespData RelationRevokedResponseData `json:"resp_data,omitempty"` // 响应数据
}

// SuccessOrFail 判断响应是否成功
func (r *RelationRevokedResponse) SuccessOrFail() bool {
	return r.Code == "BBS00000"
}
