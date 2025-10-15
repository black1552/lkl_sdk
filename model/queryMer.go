// Package model 提供数据模型定义
package model

// QueryMerRequest 商户进件信息查询请求
// 接入方通过开放平台查询进件信息，返回报文同进件审核完成主动通知报文
// 参考文档: https://o.lakala.com/#/home/document/detail?id=102

type QueryMerRequest struct {
	// ReqData 请求业务参数
	ReqData *QueryMerRequestData `json:"reqData"`
	// Ver 版本号
	Ver string `json:"ver"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// ReqId 请求ID
	ReqId string `json:"reqId"`
}

// QueryMerRequestData 商户进件信息查询请求业务参数

type QueryMerRequestData struct {
	// Version 接口版本号
	Version string `json:"version"`
	// OrderNo 订单编号（便于后续跟踪排查问题及核对报文）
	// 14位年月日时（24小时制）分秒+8位的随机数（不重复）如：2021020112000012345678
	OrderNo string `json:"orderNo"`
	// OrgCode 机构代码
	OrgCode string `json:"orgCode"`
	// ContractId 进件ID
	ContractId string `json:"contractId"`
}

// QueryMerResponse 商户进件信息查询响应

type QueryMerResponse struct {
	// CmdRetCode 全局返回码
	CmdRetCode string `json:"cmdRetCode"`
	// ReqId 请求ID
	ReqId string `json:"reqId"`
	// RespData 响应业务参数
	RespData *QueryMerResponseData `json:"respData"`
	// RetCode 返回码
	RetCode string `json:"retCode"`
	// md 随机字符串
	Md string `json:"md"`
	// RetMsg 返回消息
	RetMsg string `json:"retMsg"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// Ver 版本号
	Ver string `json:"ver"`
}

// QueryMerResponseData 商户进件信息查询响应业务参数

type QueryMerResponseData struct {
	// OrgCode 机构代码
	OrgCode string `json:"orgCode"`
	// OrderNo 订单号
	OrderNo string `json:"orderNo"`
	// ContractId 进件ID
	ContractId string `json:"contractId"`
	// ContractStatus 进件状态
	// 未提交：NO_COMMIT
	// 已提交：COMMIT
	// 提交失败：COMMIT_FAIL
	// 转人工审核：MANUAL_AUDIT
	// 审核中：REVIEW_ING
	// 审核通过：WAIT_FOR_CONTACT
	// 审核驳回：INNER_CHECK_REJECTED
	ContractStatus string `json:"contractStatus"`
	// ContractMemo 进件描述
	// 进件审核通过，返回"审核通过"
	// 进件审核驳回，返回具体的驳回理由
	ContractMemo string `json:"contractMemo"`
	// MerInnerNo 拉卡拉内部商户号（该属性审核通过才有）
	MerInnerNo string `json:"merInnerNo"`
	// MerCupNo 银联商户号（该属性审核通过才有）
	MerCupNo string `json:"merCupNo"`
	// TermDatas 终端列表信息（该属性审核通过并且是增商、增终进件才有）
	TermDatas []*TermData `json:"termDatas"`
}

// TermData 终端数据信息

type TermData struct {
	// ShopId 网点编号
	ShopId string `json:"shopId"`
	// TermId 终端编号
	TermId string `json:"termId"`
	// TermNo 终端号
	TermNo string `json:"termNo"`
	// BusiTypeCode 业务代码（参考【业务类型字典表】文档）
	BusiTypeCode string `json:"busiTypeCode"`
	// BusiTypeName 业务名称
	BusiTypeName string `json:"busiTypeName"`
	// ProductName 产品名称
	ProductName string `json:"productName"`
	// ProductCode 产品代码
	ProductCode string `json:"productCode"`
	// DevSerialNo 终端设备序列号
	DevSerialNo string `json:"devSerialNo"`
}

func (t *QueryMerResponse) SuccessOrFail() bool {
	return t.RetCode == "000000"
}
