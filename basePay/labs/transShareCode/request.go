package transShareCode

// LocationInfo 地址位置信息
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`   // 请求方IP地址，必填
	BaseStation string `json:"base_station"` // 客户端设备的基站信息
	Location    string `json:"location"`     // 商户终端的地理位置，整体格式：纬度,经度
}

// ExtInfo 扩展内容，主要满足花呗分期相关的额鉴权验等功能
type ExtInfo struct {
	FqNumber        string `json:"fq_number"`         // 花呗分期期数，支付宝花呗分期必送字段: 花呗分期数 3：3期 6：6期 12：12期
	FqSellerPercent string `json:"fq_seller_percent"` // 卖家承担手续费比例，支付宝花呗分期必送字段: 卖家承担收费比例，商家承担手续费传入100，用户承担手续费传入0
}

// AccBusiFields 账户端业务信息域
type AccBusiFields struct {
	Source   string   `json:"source"`    // 业务来源，必填
	BizLink  string   `json:"biz_link"`  // 跳转业务链接，必填
	SellerID string   `json:"seller_id"` // 卖家支付宝ID，可选
	ExtInfo  *ExtInfo `json:"ext_info"`  // 扩展内容，可选
}

// TransShareCodeRequestData 请求数据
type TransShareCodeRequestData struct {
	MerchantNo      string         `json:"merchant_no"`       // 商户号，必填
	TermNo          string         `json:"term_no"`           // 终端号，必填
	OutTradeNo      string         `json:"out_trade_no"`      // 商户交易流水号，必填
	AccountType     string         `json:"account_type"`      // 钱包类型，必填，支付宝：ALIPAY
	TotalAmount     string         `json:"total_amount"`      // 金额，必填，单位分，整数型字符
	LocationInfo    LocationInfo   `json:"location_info"`     // 地址位置信息，必填
	CodeValidPeriod string         `json:"code_valid_period"` // 码有效期，可选，秒为单位,整型
	AccBusiFields   *AccBusiFields `json:"acc_busi_fields"`   // 账户端业务信息域，可选
}

// TransShareCodeRequest 整体请求结构
type TransShareCodeRequest struct {
	ReqTime string                     `json:"req_time"` // 请求时间
	Version string                     `json:"version"`  // 版本号
	ReqData *TransShareCodeRequestData `json:"req_data"` // 请求数据
}
