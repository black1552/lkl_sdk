package transPreorederEncry

// LocationInfo 地址位置信息，风控要求必送
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`   // 请求方IP地址，必填
	BaseStation string `json:"base_station"` // 客户端设备的基站信息（主扫时基站信息使用该字段）
	Location    string `json:"location"`     // 商户终端的地理位置，整体格式：纬度,经度
}

// AccBusiFields 账户端业务信息域
type AccBusiFields struct {
	UserID             string                 `json:"user_id"`              // 买家在支付宝的用户id
	TimeoutExpress     string                 `json:"timeout_express"`      // 预下单有效时间，以分钟为单位
	ExtendParams       map[string]interface{} `json:"extend_params"`        // 支付宝业务扩展参数
	GoodsDetail        string                 `json:"goods_detail"`         // 订单包含的商品列表信息，Json数组
	StoreID            string                 `json:"store_id"`             // 商户门店编号
	DisablePayChannels string                 `json:"disable_pay_channels"` // 支付宝禁用支付渠道
	BusinessParams     string                 `json:"business_params"`      // 商户传入业务信息
	MinAge             string                 `json:"min_age"`              // 允许的最小买家年龄
}

// RequestData 聚合扫码主扫交易（全报文加密）请求数据
type TransPreorederEncryRequestData struct {
	MerchantNo    string         `json:"merchant_no"`     // 商户号，必填
	TermNo        string         `json:"term_no"`         // 终端号，必填
	OutTradeNo    string         `json:"out_trade_no"`    // 商户交易流水号，必填
	AccountType   string         `json:"account_type"`    // 钱包类型，必填：WECHAT、ALIPAY、UQRCODEPAY等
	TransType     string         `json:"trans_type"`      // 接入方式，必填：41:NATIVE、51:JSAPI、71:微信小程序支付、61:APP支付
	TotalAmount   string         `json:"total_amount"`    // 金额，必填，单位分，整数型字符
	LocationInfo  *LocationInfo  `json:"location_info"`   // 地址位置信息，必填
	BusiMode      string         `json:"busi_mode"`       // 业务模式，可选：ACQ-收单，不填默认为"ACQ-收单"
	Subject       string         `json:"subject"`         // 订单标题，可选，微信支付必送
	PayOrderNo    string         `json:"pay_order_no"`    // 支付业务订单号，可选
	NotifyURL     string         `json:"notify_url"`      // 商户通知地址，可选
	SettleType    string         `json:"settle_type"`     // 结算类型，可选："0"或者空为常规结算，"1"为拉卡拉分账
	Remark        string         `json:"remark"`          // 备注，可选
	IdentityInfo  string         `json:"identity_info"`   // 实名支付信息，可选，json字符串
	AccBusiFields *AccBusiFields `json:"acc_busi_fields"` // 账户端业务信息域，可选
}

// Request 聚合扫码主扫交易（全报文加密）请求结构
type TransPreorederEncryRequest struct {
	ReqTime string                          `json:"req_time"` // 请求时间
	Version string                          `json:"version"`  // 版本号
	ReqData *TransPreorederEncryRequestData `json:"req_data"` // 请求数据
}
