// Package model 提供数据模型定义
package model

// MerValidateRequest 商户进件校验请求
// 接入方在入网前，对商户名称、法人信息、结算人信息做敏感词和黑名单、反洗钱的校验
// 参考文档: https://o.lakala.com/#/home/document/detail?id=93

type MerValidateRequest struct {
	// ReqData 请求业务参数
	ReqData *MerValidateRequestData `json:"reqData"`
	// Ver 版本号
	Ver string `json:"ver"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// ReqId 请求ID
	ReqId string `json:"reqId"`
}

// MerValidateRequestData 商户进件校验请求业务参数

type MerValidateRequestData struct {
	// Version 接口版本号
	Version string `json:"version"`
	// OrderNo 订单编号（便于后续跟踪排查问题及核对报文）
	// 14位年月日时（24小时制）分秒+8位的随机数（不重复）如：2021020112000012345678
	OrderNo string `json:"orderNo"`
	// OrgCode 机构代码
	OrgCode string `json:"orgCode"`
	// MerRegName 商户注册名称（可传）
	MerRegName string `json:"merRegName"`
	// MerBizName 商户经营名称（可传）
	MerBizName string `json:"merBizName"`
	// MerBlis 营业执照（可传）
	MerBlis string `json:"merBlis"`
	// LarIdcard 法人身份证号（可传）
	LarIdcard string `json:"larIdcard"`
	// AcctNo 结算账号（可传）
	AcctNo string `json:"acctNo"`
	// AcctIdcard 结算人证件号（可传）
	AcctIdcard string `json:"acctIdcard"`
}

// MerValidateResponse 商户进件校验响应

type MerValidateResponse struct {
	// Appid 应用ID
	Appid string `json:"appid"`
	// CmdRetCode 全局返回码
	CmdRetCode string `json:"cmdRetCode"`
	// ReqId 请求ID
	ReqId string `json:"reqId"`
	// RetCode 返回码
	RetCode string `json:"retCode"`
	// RetMsg 返回消息
	RetMsg string `json:"retMsg"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// Ver 版本号
	Ver string `json:"ver"`
}

func (t *MerValidateResponse) SuccessOrFail() bool {
	return t.RetCode == "000000"
}
