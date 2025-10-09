package model

import "github.com/gogf/gf/v2/os/gtime"

// MergePreorder 主扫合单交易请求结构体
type MergePreorder struct {
	ReqTime string                 `json:"req_time"` // 请求时间，格式：YYYYMMDDHHMMSS
	Version string                 `json:"version"`  // 接口版本号，固定值"3.0"
	ReqData *MergePreorderReqData `json:"req_data"` // 请求业务数据
}

// MergePreorderReqData 主扫合单交易请求业务数据
type MergePreorderReqData struct {
	MerchantNo   string       `json:"merchant_no"`   // 商户号，拉卡拉分配的商户号，String(32)，必填
	TermNo       string       `json:"term_no"`       // 终端号，拉卡拉分配的业务终端号，String(32)，必填
	OutTradeNo   string       `json:"out_trade_no"`  // 商户交易流水号，商户系统唯一，String(32)，必填
	OutSplitInfo []*OutSplitInfo `json:"out_split_info"` // 拆单信息，List，必填
	AccountType  string       `json:"account_type"`  // 钱包类型，微信：WECHAT 支付宝：ALIPAY 银联：UQRCODEPAY 京东钱包：JD，String(32)，必填
	TransType    string       `json:"trans_type"`    // 接入方式，41:NATIVE（扫码支付）(仅ALIPAY支持) 51:JSAPI（微信公众号支付，支付宝服务窗支付，银联JS支付，支付宝JS支付、拉卡拉钱包支付）71:微信小程序支付 81:支付宝H5支付（需特殊商户账户端支持），String(2)，必填
	TotalAmount  string       `json:"total_amount"`  // 金额，单位分，整数型字符，String(12)，必填
	LocationInfo *LocationInfo `json:"location_info"` // 地址位置信息，Object，风控要求必送，必填
	BusiMode     string       `json:"busi_mode"`     // 业务模式，ACQ-收单 PAY-付款不填，默认为"ACQ-收单"，String(8)，选填
	Subject      string       `json:"subject"`       // 订单标题，用于简单描述订单或商品主题，传递给账户端（账户端控制，实际最多42个字节），String(42)，选填
	NotifyUrl    string       `json:"notify_url"`    // 商户通知地址，如果上传，且 pay_order_no 不存在情况下，则按此地址通知商户，String(128)，选填
	Remark       string       `json:"remark"`        // 备注，String(128)，选填
	IdentityInfo string       `json:"identity_info"` // 实名支付信息，json字符串，如{"identityNo": "3200000000000000XX", "name": "张三"}，然后国密sm2加密，String(1024)，选填
	AccBusiFields *AccBusiFields `json:"acc_busi_fields"` // 账户端业务信息域，Object，选填
	CompleteNotifyUrl string `json:"complete_notify_url"` // 发货确认通知地址，发货类小程序确认收货后通知商户的地址，String(128)，选填
}

// OutSplitInfo 拆单信息
type OutSplitInfo struct {
	OutSubTradeNo string `json:"out_sub_trade_no"` // 外部子交易流水号，商户子交易流水号，商户号下唯一，String(32)，必填
	MerchantNo    string `json:"merchant_no"`      // 商户号，拉卡拉分配的商户号，String(32)，必填
	TermNo        string `json:"term_no"`          // 终端号，拉卡拉分配的业务终端号，String(32)，必填
	Amount        string `json:"amount"`           // 金额，单位分，整数型字符，String(12)，必填
	SettleType    string `json:"settle_type"`      // 结算类型（单台），"0"或者空，常规结算方式，String(4)，选填
	SubRemark     string `json:"sub_remark"`       // 子备注，子单备注信息，String(64)，选填
}

// LocationInfo 地址位置信息
type LocationInfo struct {
	RequestIp string `json:"request_ip"` // 请求方IP地址，请求方的IP地址，存在必填，格式如36.45.36.95，String(64)，必填
	BaseStation string `json:"base_station"` // 基站信息，客户端设备的基站信息（主扫时基站信息使用该字段），String(128)，选填
	Location string `json:"location"`     // 维度,经度，商户终端的地理位置，存在必填格式：纬度,经度，+表示北纬、东经，-表示南纬、西经，精度最长支持小数点后9位。举例:+37.123456789,121.123456789，String(32)，选填
}

// AccBusiFields 账户端业务信息域，微信主扫场景
type AccBusiFields struct {
	TimeoutExpress string       `json:"timeout_express"` // 预下单的订单的有效时间，以分钟为单位。如果在有效时间内没有完成付款，则在账户端该订单失效。如用户超时，则账户端完成失效处理，建议不超过15分钟。不传值则默认5分钟，String(2)，选填
	SubAppId       string       `json:"sub_appid"`       // 子商户公众账号ID，sub_appid（即微信小程序支付-71、公众号支付-51、微信支付-61），此参数必传，只对微信支付生效；拉卡拉钱包情况下，该字段上送LAKALA的openid，String(32)，选填
	UserId         string       `json:"user_id"`         // 用户标识，用户在子商户sub_appid下的唯一标识，sub_openid，（即微信小程序支付-71、公众号支付-51），此参数必传，只对微信支付有效，String(64)，选填
	Detail         string       `json:"detail"`          // 商品详情，单品优惠功能字段，详见下文说明，String(1024)，选填
	GoodsTag       string       `json:"goods_tag"`       // 订单优惠标记，微信平台配置的商品标记，用于优惠券或者满减使用，accountType为WECHAT时，可选填此字段，String(32)，选填
	Attach         string       `json:"attach"`          // 附加域，附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据。商户定制字段，直接送到账户端，String(128)，选填
	GoodsDetail    []*GoodsDetail `json:"goods_detail"`    // 商品详情列表，微信商品详情字段说明
}

// GoodsDetail 微信商品详情字段说明
type GoodsDetail struct {
	GoodsId      string `json:"goods_id"`      // 商品ID，由半角的大小写字母、数字、中划线、下划线中的一种或几种组成。如"商品编码"，必填
	WxpayGoodsId string `json:"wxpay_goods_id"` // 微信支付定义的统一商品编号，String(32)，选填
	GoodsName    string `json:"goods_name"`    // 商品的实际名称，String(256)，选填
	Quantity     string `json:"quantity"`      // 用户购买的数量，String(12)，必填
	Price        string `json:"price"`         // 单价，单位为：分。如果商户有优惠，需传输商户优惠后的单价，String(12)，必填
}

// NewMergePreorder 创建主扫合单交易请求
func NewMergePreorder(param *MergePreorderReqData) *MergePreorder {
	return &MergePreorder{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: param,
	}
}

// MergePreorderResponse 主扫合单交易响应结构体
type MergePreorderResponse struct {
	Code     string            `json:"code"`      // 响应码，BBS00000表示成功
	Msg      string            `json:"msg"`       // 响应信息，对响应码的文字描述
	ReqData  MergePreorderRespData `json:"resp_data"` // 响应业务数据
	RespTime string            `json:"resp_time"` // 响应时间
}

// MergePreorderRespData 主扫合单交易响应业务数据
type MergePreorderRespData struct {
	MerchantNo string       `json:"merchant_no"`    // 商户号（待上线），拉卡拉分配的商户号（请求接口中商户号）
	OutTradeNo string       `json:"out_trade_no"`   // 商户请求流水号，请求报文中的商户请求流水号
	TradeNo    string       `json:"trade_no"`       // 拉卡拉交易流水号
	LogNo      string       `json:"log_no"`         // 拉卡拉对账单流水号
	SplitInfo  []*SplitInfo `json:"split_info"`     // 拆单信息
	AccRespFields interface{} `json:"acc_resp_fields"` // 账户端返回信息域
}

// SplitInfo 拆单信息
type SplitInfo struct {
	SubTradeNo    string `json:"sub_trade_no"`    // 子单交易流水号
	SubLogNo      string `json:"sub_log_no"`      // 子单对账单流水号
	OutSubTradeNo string `json:"out_sub_trade_no"` // 外部子交易流水号，商户子交易流水号，商户号下唯一
	MerchantNo    string `json:"merchant_no"`     // 商户号，拉卡拉分配的商户号
	MerchantName  string `json:"merchant_name"`   // 商户名称
	TermNo        string `json:"term_no"`         // 终端号，拉卡拉分配的业务终端号
	Amount        string `json:"amount"`          // 金额，单位为：分。整数型字符
}

// WxAccRespFields 微信(71-小程序/微信(51-JSAPI)场景下账户端返回信息域
type WxAccRespFields struct {
	PrepayId   string `json:"prepay_id"`   // 预下单ID，预支付交易会话ID
	PaySign    string `json:"pay_sign"`    // 支付签名信息
	AppId      string `json:"app_id"`      // 小程序ID，商户注册具有支付权限的小程序成功后即可获得小程序ID
	TimeStamp  string `json:"time_stamp"`  // 时间戳，当前的时间
	NonceStr   string `json:"nonce_str"`   // 随机字符串
	Package    string `json:"package"`     // 订单详情扩展字符串
	SignType   string `json:"sign_type"`   // 签名方式，签名类型，支持RSA
	SubMchId   string `json:"sub_mch_id"`  // 子商户号，账户端子商户号
}

// SuccessOrFail 判断主扫合单交易是否成功
func (m *MergePreorderResponse) SuccessOrFail() bool {
	return m.Code == "BBS00000"
}