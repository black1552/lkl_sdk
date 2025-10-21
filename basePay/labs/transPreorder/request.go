package transPreorder

// LocationInfo 地址位置信息结构体
type LocationInfo struct {
	RequestIP   string `json:"request_ip"`             // 请求方IP地址 (必填)
	BaseStation string `json:"base_station,omitempty"` // 基站信息 (可选)
	Location    string `json:"location,omitempty"`     // 纬度,经度 (可选，银联二维码交易必填)
}

// GoodsInfo 商品信息结构体
type GoodsInfo struct {
	GoodsId    string `json:"goods_id,omitempty"`    // 商品ID (可选)
	GoodsName  string `json:"goods_name,omitempty"`  // 商品名称 (可选)
	Price      string `json:"price,omitempty"`       // 商品单价 (可选)
	Quantity   string `json:"quantity,omitempty"`    // 商品数量 (可选)
	CategoryId string `json:"category_id,omitempty"` // 商品类目ID (可选)
	Body       string `json:"body,omitempty"`        // 商品描述 (可选)
}

// AccBusiFields 账户端业务信息域结构体
type AccBusiFields struct {
	UserId            string `json:"user_id,omitempty"`             // 买家在支付宝的用户id (支付宝JSAPI支付必填)
	TimeoutExpress    string `json:"timeout_express,omitempty"`     // 预下单有效时间 (可选)
	ExtendParams      string `json:"extend_params,omitempty"`       // 业务扩展参数 (可选)
	GoodsDetail       string `json:"goods_detail,omitempty"`        // 商品详情 (可选)
	StoreId           string `json:"store_id,omitempty"`            // 商户门店编号 (可选)
	BusinessParams    string `json:"business_params,omitempty"`     // 商户传入业务信息 (可选)
	MinAge            string `json:"min_age,omitempty"`             // 允许的最小买家年龄 (可选)
	PromoParams       string `json:"promo_params,omitempty"`        // 优惠明细 (可选)
	EnablePayChannels string `json:"enable_pay_channels,omitempty"` // 支付通道 (可选)
}

// TransPreorderRequestData 预订单请求数据结构体
type TransPreorderRequestData struct {
	MerchantNo        string         `json:"merchant_no"`                   // 商户号 (必填)
	TermNo            string         `json:"term_no"`                       // 终端号 (必填)
	OutTradeNo        string         `json:"out_trade_no"`                  // 商户交易流水号 (必填)
	AccountType       string         `json:"account_type"`                  // 钱包类型 (必填)
	TransType         string         `json:"trans_type"`                    // 接入方式 (必填)
	TotalAmount       string         `json:"total_amount"`                  // 金额 (必填)
	LocationInfo      *LocationInfo  `json:"location_info"`                 // 地址位置信息 (必填)
	BusiMode          string         `json:"busi_mode,omitempty"`           // 业务模式 (可选)
	Subject           string         `json:"subject,omitempty"`             // 订单标题 (可选，微信支付必送)
	PayOrderNo        string         `json:"pay_order_no,omitempty"`        // 支付业务订单号 (可选)
	NotifyURL         string         `json:"notify_url,omitempty"`          // 商户通知地址 (可选)
	SettleType        string         `json:"settle_type,omitempty"`         // 结算类型 (可选)
	Remark            string         `json:"remark,omitempty"`              // 备注 (可选)
	AccBusiFields     *AccBusiFields `json:"acc_busi_fields,omitempty"`     // 账户端业务信息域 (可选)
	CompleteNotifyURL string         `json:"complete_notify_url,omitempty"` // 发货确认通知地址 (可选)
	GoodsTag          string         `json:"goods_tag,omitempty"`           // 商品标签 (可选)
	GoodsInfo         []*GoodsInfo   `json:"goods_info,omitempty"`          // 商品信息 (可选)
}

// TransPreorderRequest 预订单请求结构体
type TransPreorderRequest struct {
	ReqTime string                    `json:"req_time"`
	Version string                    `json:"version"`
	ReqData *TransPreorderRequestData `json:"req_data"`
}
