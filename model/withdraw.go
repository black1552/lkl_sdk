package model

import "github.com/black1552/lkl_sdk/consts"

// WithdrawRequest 提现请求结构体
// 用于向拉卡拉接口发送提现请求
// 包含请求头信息和业务数据
type WithdrawRequest struct {
	// 请求业务数据
	ReqData *WithdrawReqData `json:"req_data"`
	// 接口版本号
	Version string `json:"version"`
	// 请求时间，格式为yyyyMMddHHmmss
	ReqTime string `json:"req_time"`
}

// WithdrawReqData 提现请求业务数据结构体
// 包含提现所需的具体业务参数
// 用于申请账户资金提现
type WithdrawReqData struct {
	// bmcp机构号，必传，最大长度32
	OrgNo string `json:"org_no"`
	// 商户号，必传，最大长度32，822商户号，SR分账接收方编号
	MerchantNo string `json:"merchant_no"`
	// 提现金额（单位：元），必传，最大长度32
	DrawAmt string `json:"draw_amt"`
	// 通知地址，非必传，最大长度256
	NotifyUrl string `json:"notify_url"`
	// 商户订单号（商户系统唯一），非必传，最大长度256
	MerOrderNo string `json:"mer_order_no"`
	// 账号（若该参数上送，则payType将无效），非必传，最大长度32
	PayNo string `json:"pay_no"`
	// 账号类型（01：收款账户，04：分账接收方账户）未上送则默认01，必传，最大长度32，分账接收方提现时需填04
	PayType consts.PayType `json:"pay_type"`
	// 备注信息，非必传，最大长度64
	Remark string `json:"remark"`
	// 摘要，非必传，最大长度64
	Summary string `json:"summary"`
	// 结算银行ID，非必传，最大长度32
	BankId string `json:"bank_id"`
}

// WithdrawResponse 提现响应结构体
// 拉卡拉接口返回的提现响应数据
// 包含响应状态码、消息和业务数据
type WithdrawResponse struct {
	// 响应状态码，000000表示成功
	Code string `json:"code"`
	// 响应消息
	Msg string `json:"msg"`
	// 响应业务数据，当code为000000时返回
	RespData *WithdrawRespData `json:"resp_data"`
}

// WithdrawRespData 提现响应业务数据结构体
// 包含提现返回的具体业务信息
type WithdrawRespData struct {
	// 提现流水号，必传
	DrawJnl string `json:"draw_jnl"`
	// 商户订单号，必传
	MerOrderNo string `json:"mer_order_no"`
}

// SuccessOrFail 判断提现请求是否成功
// 成功条件：响应码为000000
// 返回值：true表示成功，false表示失败
func (resp *WithdrawResponse) SuccessOrFail() bool {
	return resp.Code == "000000"
}
