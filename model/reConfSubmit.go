package model

// ReConfSubmitRequest 进件复议提交请求
// 接入方在进件被驳回后（由于系统自动校验无法通过），调用此接口提交进件，转人工审核
// 参考文档: https://o.lakala.com/#/home/document/detail?id=94

type ReConfSubmitRequest struct {
	// ReqData 请求业务参数
	ReqData *ReConfSubmitRequestData `json:"reqData"`
	// Ver 版本号
	Ver string `json:"ver"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// ReqId 请求ID
	ReqId string `json:"reqId"`
}

// ReConfSubmitRequestData 进件复议提交请求业务参数

type ReConfSubmitRequestData struct {
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

// ReConfSubmitResponse 进件复议提交响应

type ReConfSubmitResponse struct {
	// CmdRetCode 全局返回码
	CmdRetCode string `json:"cmdRetCode"`
	// ReqId 请求ID
	ReqId string `json:"reqId"`
	// RetCode 返回码
	RetCode string `json:"retCode"`
	// md 随机字符串
	Md string `json:"md"`
	// RespData 响应业务参数
	RespData *ReConfSubmitResponseData `json:"respData"`
	// RetMsg 返回消息
	RetMsg string `json:"retMsg"`
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// Ver 版本号
	Ver string `json:"ver"`
}
type ReConfSubmitResponseData struct {
	// OrgCode 机构代码
	OrgCode string `json:"orgCode"`
	// OrderNo 订单号
	OrderNo string `json:"orderNo"`
}

func (t *ReConfSubmitResponse) SuccessOrFail() bool {
	return t.RetCode == "000000"
}
