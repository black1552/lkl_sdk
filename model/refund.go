package model

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gtime"
)

// Refund 退款请求结构体
type Refund struct {
	ReqTime string         `json:"req_time"` // 请求时间，格式：YYYYMMDDHHMMSS
	Version string         `json:"version"`  // 接口版本号，固定值"3.0"
	ReqData *RefundReqData `json:"req_data"` // 请求业务数据
}

// RefundReqData 退款请求业务数据
type RefundReqData struct {
	MerchantNo    string               `json:"merchant_no"`     // 商户号，拉卡拉分配的商户号，String(32)，必填
	TermNo        string               `json:"term_no"`         // 终端号，拉卡拉分配的业务终端号，String(32)，必填
	OutTradeNo    string               `json:"out_trade_no"`    // 商户请求流水号，商户系统唯一，String(32)，必填
	RefundAmount  string               `json:"refund_amount"`   // 退款金额，单位分，整数型字符，String(12)，必填
	RefundAccMode consts.RefundAccMode `json:"refund_acc_mode"` // 退款账户模式，String(2)，必填，00-调用户余额 65-调商户余额 66-调终端余额 30-调账户
	LocationInfo  struct {
		RequestIp string `json:"request_ip"` // 请求方IP地址，请求方的IP地址，存在必填，格式如36.45.36.95，String(64)，必填
	} `json:"location_info"` // 地址位置信息，Object，必填
	NotifyUrl        string              `json:"notify_url"`          // 后台通知地址，交易结果通知地址，String(128)，选填
	RefundAmtSts     consts.RefundAmtSts `json:"refund_amt_sts"`      // 退货资金状态 String(2) 00-为默认，01-为分账；分账交易退款必须填写
	OriginLogNo      string              `json:"origin_log_no"`       // 拉卡拉对账单流水号，正常退款的拉卡拉对账单流水号，String(14)，选填
	OriginOutTradeNo string              `json:"origin_out_trade_no"` // 原始交易商户流水号，String(32)，选填
	OriginTradeNo    string              `json:"origin_trade_no"`     // 原交易拉卡拉交易订单号，String(32)，选填
	RefundSplitMsg   string              `json:"refund_split_msg"`    // 退款分账状态，String(2)，选填，00-为默认，01-为分账；分账交易退款必须填写。需要退款上送该笔的分账状态，为分账时，是退分账前处理，还是退分账后处理
}

// NewRefund 创建退款请求
func NewRefund(param *RefundReqData) *Refund {
	return &Refund{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: param,
	}
}

// RefundResponse 退款响应结构体
type RefundResponse struct {
	Code     string         `json:"code"`      // 响应码，"000000"表示成功
	Msg      string         `json:"msg"`       // 响应信息，对响应码的文字描述
	RespData RefundRespData `json:"resp_data"` // 响应业务数据
	RespTime string         `json:"resp_time"` // 响应时间
}

// RefundRespData 退款响应业务数据
type RefundRespData struct {
	TradeState       string `json:"trade_state"`         // 交易状态，String(15)，必填，INIT-初始化（需商户确认结果）；SUCCESS-交易成功；FAIL-交易失败；REFUND-交易退款中（需商户确认结果）；PROCESSING-交易处理中（需商户确认结果）；TIMEOUT-请求超时（需商户确认结果）；EXCEPTION-异常（失败）
	RefundType       string `json:"refund_type"`         // 退款模式，String(20)，必填
	MerchantNo       string `json:"merchant_no"`         // 商户号，拉卡拉分配的商户号，String(32)，必填
	OutTradeNo       string `json:"out_trade_no"`        // 商户请求流水号，请求报文中的商户请求流水号，String(32)，必填
	TradeNo          string `json:"trade_no"`            // 拉卡拉交易流水号，String(32)，必填
	LogNo            string `json:"log_no"`              // 拉卡拉对账单流水号，String(14)，必填
	AccType          string `json:"acc_type"`            // 账户类型，String(32)，必填
	TotalAmount      string `json:"total_amount"`        // 交易金额，单位分，整数型字符，String(12)，必填
	RefundAmount     string `json:"refund_amount"`       // 申请退款金额，单位分，整数型字符，String(12)，必填
	PayedAmount      string `json:"payed_amount"`        // 实际退款金额，单位分，整数型字符，String(12)，必填
	TradeTime        string `json:"trade_time"`          // 退款时间，实际退款时间，格式：yyyyMMddHHmmss，String(14)，选填
	OriginLogNo      string `json:"origin_log_no"`       // 原拉卡拉对账单流水号，原交易的拉卡拉对账单流水号，String(14)，选填
	OriginOutTradeNo string `json:"origin_out_trade_no"` // 原商户请求流水号，原交易中的商户请求流水号，String(32)，选填
	OriginTradeNo    string `json:"origin_trade_no"`     // 原交易拉卡拉交易订单号，String(32)，选填
	UpCouponInfo     string `json:"up_coupon_info"`      // 银联优惠券信息，目标字段，单位是银联侧返回的四部分内容：{"fundChannel": "BOC", "amount": "10"}，String(500)，选填
	TanteInfo        string `json:"tante_info"`          // 淘方信息，目标字段，数据是淘方侧返回的四部分内容：{"fundChannel": "BOC", "amount": "10"}，String(32)，选填
	ChannelRetDesc   string `json:"channel_ret_desc"`    // 渠道返回描述，String，必填，codeMsg: "R000000-成功", "R011122-渠道处理超时"
}

// SuccessOrFail 判断退款交易是否成功
func (r *RefundResponse) SuccessOrFail() bool {
	return r.Code == "000000"
}
