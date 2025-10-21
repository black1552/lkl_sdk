package transMicropayEncry

// LocationInfo 地址位置信息，风控要求必送
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`   // 请求方IP地址，必填
	BaseStation string `json:"base_station"` // 客户端设备的基站信息
	Location    string `json:"location"`     // 商户终端的地理位置，整体格式：纬度,经度
}

// AccBusiFields 账户端业务信息域
type AccBusiFields struct {
	ExtendParams       map[string]interface{} `json:"extend_params"`        // 支付宝业务扩展参数，主要用于花呗分期
	BusinessParams     string                 `json:"business_params"`      // 商户传入业务信息
	GoodsDetail        string                 `json:"goods_detail"`         // 订单包含的商品列表信息，Json数组
	StoreID            string                 `json:"store_id"`             // 商户门店编号
	TimeoutExpress     string                 `json:"timeout_express"`      // 交易有效时间，以分钟为单位
	DisablePayChannels string                 `json:"disable_pay_channels"` // 支付宝禁用支付渠道
	MinAge             string                 `json:"min_age"`              // 允许的最小买家年龄
	PriorityPayAssets  map[string]interface{} `json:"priority_pay_assets"`  // 优先使用资产
}

// TransMicropayEncryRequestData 聚合扫码被扫接口（全报文加密）请求数据
type TransMicropayEncryRequestData struct {
	MerchantNo    string        `json:"merchant_no"`     // 商户号，必填
	TermNo        string        `json:"term_no"`         // 终端号，必填
	OutTradeNo    string        `json:"out_trade_no"`    // 商户交易流水号，必填
	AuthCode      string        `json:"auth_code"`       // 支付授权码，必填
	TotalAmount   string        `json:"total_amount"`    // 金额，必填，单位分，整数型字符
	LocationInfo  LocationInfo  `json:"location_info"`   // 地址位置信息，必填
	BusiMode      string        `json:"busi_mode"`       // 业务模式，可选：ACQ-收单，不填默认为"ACQ-收单"
	Subject       string        `json:"subject"`         // 订单标题，可选，微信支付必送
	PayOrderNo    string        `json:"pay_order_no"`    // 拉卡拉支付业务订单号，可选
	NotifyURL     string        `json:"notify_url"`      // 商户通知地址，可选
	SettleType    string        `json:"settle_type"`     // 结算类型，可选："0"或者空，常规结算方式
	Remark        string        `json:"remark"`          // 备注，可选
	ScanType      string        `json:"scan_type"`       // 扫码类型，可选：0或不填：扫码支付，1：支付宝刷脸支付，2: 微信刷脸支付
	IdentityInfo  string        `json:"identity_info"`   // 实名支付信息，可选，json字符串
	AccBusiFields AccBusiFields `json:"acc_busi_fields"` // 账户端业务信息域，可选
}

// TransMicropayEncryRequest 聚合扫码被扫接口（全报文加密）请求结构
type TransMicropayEncryRequest struct {
	ReqTime string                         `json:"req_time"` // 请求时间
	Version string                         `json:"version"`  // 版本号
	ReqData *TransMicropayEncryRequestData `json:"req_data"` // 请求数据
}
