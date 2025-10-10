package model

import "github.com/black1552/lkl_sdk/consts"

// OrderSplitLedgerRequest 订单分账请求结构体
// 用于发起订单分账操作，支持向多个接收方进行分账
// 拉卡拉SDK接口文档：订单分账接口

type OrderSplitLedgerRequest struct {
	ReqData *OrderSplitLedgerReqData `json:"req_data"` // 请求业务数据
	Version string                   `json:"version"`  // 接口版本号
	ReqTime string                   `json:"req_time"` // 请求时间，格式为yyyyMMddHHmmss
}

// OrderSplitLedgerReqData 订单分账请求业务数据结构体
// 包含订单分账所需的详细业务参数

type OrderSplitLedgerReqData struct {
	MerchantNo    string                       `json:"merchant_no"`          // 商户号，必传，长度32
	LogNo         string                       `json:"log_no"`               // 拉卡拉对账单流水号，必传，长度14，posp流水号，用于查询结算
	LogDate       string                       `json:"log_date"`             // 交易日期，必传，长度8，posp日期，格式为yyyyMMdd，用于查询结算
	OutSeparateNo string                       `json:"out_separate_no"`      // 商户分账指令流水号，必传，长度32，每个商户号下唯一，否则会校验失败
	TotalAmt      string                       `json:"total_amt"`            // 分账总金额，必传，长度15，单位为分
	LklOrgNo      string                       `json:"lkl_org_no"`           // 拉卡拉机构编号，条件必传，长度16
	CalType       consts.CalType               `json:"cal_type"`             // 分账计算类型，条件必传，长度2，取值说明：0-按照指定金额，1-按照指定比例，默认0
	NotifyUrl     string                       `json:"notify_url"`           // 回调地址，条件必传，长度128，分账、分账撤销或分账回退时，通过该地址通知商户最终处理结果，不传时不回调
	RecvDatas     []*OrderSplitLedgerRecvDatas `json:"recv_datas,omitempty"` // 分账接收数据对象，条件必传，列表类型，分账接收方编号必须已创建
}

type OrderSplitLedgerRecvDatas struct {
	RecvMerchantNo string `json:"recv_merchant_no"` // 接收方商户号，条件必传，长度32，分账接收方商户号，分给自己时使用，需和merchantNo一样，否则检查报错；分账接收方商户号和分账接收方只能填写一个
	RecvNo         string `json:"recv_no"`          // 分账接收方编号，条件必传，长度32，分账接收方编号，分给他人时使用；分账接收方商户号和分账接收方只能填写一个
	SeparateValue  string `json:"separate_value"`   // 分账数值，必传，长度32，calType为0时按照固定金额分账（单位：分），calType为1时按照比例分账（单位：百分比的小数值，如0.55表示55%）
}

// OrderSplitLedgerResponse 订单分账响应结构体
// 包含订单分账操作的处理结果信息

type OrderSplitLedgerResponse struct {
	Msg      string `json:"msg"`       // 响应消息，描述响应结果
	RespTime string `json:"resp_time"` // 响应时间，格式为yyyyMMddHHmmss
	Code     string `json:"code"`      // 响应码，SACS0000表示成功
	RespData struct {
		SeparateNo    string `json:"separate_no"`     // 分账指令流水号，必传，长度32，分账系统生成唯一流水
		OutSeparateNo string `json:"out_separate_no"` // 商户订单号，必传，长度32，请求报文中的商户外部订单号
		Status        string `json:"status"`          // 分账状态，条件必传，长度16，取值说明：处理中-PROCESSING，已受理-ACCEPTED，成功-SUCCESS，失败-FAIL
		LogNo         string `json:"log_no"`          // 拉卡拉对账单流水号，条件必传，长度14，请求透返
		LogDate       string `json:"log_date"`        // 拉卡拉订单日期，条件必传，长度8，posp日期，格式为yyyyMMdd，用于查询结算
		TotalAmt      string `json:"total_amt"`       // 分账总金额，条件必传，长度15，单位为分
	} `json:"resp_data"` // 响应业务数据，当code为SACS0000时返回
}

// SuccessOrFail 判断分账请求是否成功
// 返回true表示成功，false表示失败
func (s *OrderSplitLedgerResponse) SuccessOrFail() bool {
	return s.Code == "SACS0000"
}
