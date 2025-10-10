package model

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gtime"
)

// Preorder 预下单请求结构体
type Preorder struct {
	ReqTime string           `json:"req_time"` // 请求时间，格式：YYYYMMDDHHMMSS
	Version string           `json:"version"`  // 接口版本号，固定值"3.0"
	ReqData *PreorderReqData `json:"req_data"` // 请求业务数据
}

// PreorderReqData 预下单请求业务数据
type PreorderReqData struct {
	MerchantNo   string             `json:"merchant_no"`  // 商户号，拉卡拉分配的商户号，String(32)
	TermNo       string             `json:"term_no"`      // 终端号，拉卡拉分配的业务终端号，String(32)
	OutTradeNo   string             `json:"out_trade_no"` // 商户交易流水号，商户系统唯一，对应数据库表中外请求流水号，String(32)
	AccountType  consts.AccountType `json:"account_type"` // 钱包类型，微信：WECHAT 支付宝：ALIPAY 银联：UQRCODEPAY 翼支付：BESTPAY 苏宁易付宝：SUNING 拉卡拉支付账户：LKLACC 网联小钱包：NUCSPAY 京东钱包：JD，String(32)
	TransType    consts.TransType   `json:"trans_type"`   // 接入方式，41:NATIVE（ALIPAY, 云闪付支持, 京东白条分期）51:JSAPI（微信公众号支付, 支付宝服务窗、JS支付, 翼支付JS支付, 拉卡拉钱包支付, 京东白条分期）71:微信小程序支付61:APP支付（微信APP支付），String(2)
	TotalAmount  string             `json:"total_amount"` // 金额，单位分，整数型字符，String(12)
	NotifyUrl    string             `json:"notify_url"`   // 商户通知地址，商户通知地址，如果上传，且 pay_order_no 不存在情况下，则按此地址通知商户，String(128)
	SettleType   consts.SettleType  `json:"settle_type"`  // “0”或者空，常规结算方式，如需接拉卡拉分账通需传“1”，商户未开通分账之前切记不用上送此参数。
	LocationInfo struct {
		RequestIp string `json:"request_ip"` // 请求方IP地址，存在必填，格式如36.45.36.95，String(64)
		Location  string `json:"location"`   // 纬度,经度，商户终端的地理位置，银联二维码交易必填，整体格式：纬度,经度，+表示北纬、东经，-表示南纬、西经。经度格式：1位正负号+3位整数+1位小数点+5位小数；纬度格式：1位正负号+2位整数+1位小数点+6位小数；举例：+31.221345,+121.12345，String(32)
	} `json:"location_info"` // 地址位置信息，Object
	Subject string `json:"subject"` // 订单标题，用于简单描述订单或商品主题，传输给账户端（账户端控制，实际最多42个字节），微信支付必送，String(42)
}

func NewPreorder(param *PreorderReqData) *Preorder {
	return &Preorder{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: param,
	}
}

// PreorderResponse 预下单响应结构体
type PreorderResponse struct {
	Code     string   `json:"code"`      // 响应码，BBS00000表示成功
	Msg      string   `json:"msg"`       // 响应信息，对响应码的文字描述
	ReqData  *ReqData `json:"resp_data"` // 响应业务数据
	RespTime string   `json:"resp_time"` // 响应时间
}

// ReqData 响应业务数据
type ReqData struct {
	MerchantNo       string              `json:"merchant_no"`        // 商户号
	OutTradeNo       string              `json:"out_trade_no"`       // 外部订单号（商户订单号）
	TradeNo          string              `json:"trade_no"`           // 交易号，拉卡拉生成的订单号
	LogNo            string              `json:"log_no"`             // 拉卡拉对账单流水号
	SettleMerchantNo string              `json:"settle_merchant_no"` // 结算商户号
	SettleTermNo     string              `json:"settle_term_no"`     // 结算终端号
	AccRespFields    *WxPreorderResponse `json:"acc_resp_fields"`    // 支付通道返回的具体信息
}

// WxPreorderResponse 支付通道返回的具体信息
type WxPreorderResponse struct {
	Code        string `json:"code"`          // 返回码
	CodeImage   string `json:"code_image"`    // 二维码图片Base64编码
	PrepayId    string `json:"prepay_id"`     // 预支付ID
	AppId       string `json:"app_id"`        // 应用ID
	PaySign     string `json:"pay_sign"`      // 签名
	TimeStamp   string `json:"time_stamp"`    // 时间戳
	NonceStr    string `json:"nonce_str"`     // 随机字符串
	Package     string `json:"package"`       // 订单详情扩展字符串
	SignType    string `json:"sign_type"`     // 签名方式
	FormData    string `json:"form_data"`     // 表单数据
	RedirectUrl string `json:"redirect_url"`  // 重定向URL
	BestPayInfo string `json:"best_pay_info"` // 翼支付信息
	PartnerId   string `json:"partner_id"`    // 合作伙伴ID
	SubMchId    string `json:"sub_mch_id"`    // 子商户ID
}

func (p *PreorderResponse) SuccessOrFail() bool {
	return p.Code == "BBS00000"
}
