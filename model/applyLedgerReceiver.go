package model

import "github.com/black1552/lkl_sdk/consts"

// ApplyLedgerReceiverRequest 分账接收方创建请求结构体
// 用于向拉卡拉接口发送分账接收方创建请求
// 包含请求头信息和业务数据
type ApplyLedgerReceiverRequest struct {
	// 请求业务数据
	ReqData *ApplyLedgerReceiverReqData `json:"reqData"`
	// 请求ID，用于幂等性校验
	ReqId string `json:"reqId"`
	// 接口版本号
	Version string `json:"version"`
	// 请求时间，格式为yyyyMMddHHmmss
	ReqTime string `json:"reqTime"`
}

// ApplyLedgerReceiverReqData 分账接收方创建请求业务数据结构体
// 包含分账接收方创建所需的具体业务参数
// 用于创建分账接收方信息
type ApplyLedgerReceiverReqData struct {
	// 接口版本号，必传，长度8，取值说明：1.0
	Version string `json:"version"`
	// 订单编号（便于后续跟踪排查问题及核对报文），必传，长度32，取值说明：14位年月日（24小时制）分秒+8位的随机数（不重复）
	OrderNo string `json:"orderNo"`
	// 机构代码，必传，长度32
	OrgCode string `json:"orgCode"`
	// 分账接收方名称，必传，长度64
	ReceiverName string `json:"receiverName"`
	// 联系手机号，必传，长度16
	ContactMobile string `json:"contactMobile"`
	// 营业执照号码，可选，长度32，取值说明：收款账户类型为对公，必须上送
	LicenseNo string `json:"licenseNo"`
	// 营业执照名称，可选，长度128，取值说明：收款账户类型为对公，必须上送
	LicenseName string `json:"licenseName"`
	// 法人姓名，可选，长度32，取值说明：收款账户类型为对公，必须上送
	LegalPersonName string `json:"legalPersonName"`
	// 法人证件类型，可选，长度32，取值说明：17身份证，18护照，19港澳居民来往内地通行证，20台湾居民来往内地通行证，收款账户类型为对公，必须上送，身份证外类型先咨询后再使用
	LegalPersonCertificateType string `json:"legalPersonCertificateType"`
	// 法人证件号，可选，长度32，取值说明：收款账户类型为对公，必须上送
	LegalPersonCertificateNo string `json:"legalPersonCertificateNo"`
	// 收款账户卡号，必传，长度32
	AcctNo string `json:"acctNo"`
	// 收款账户名称，必传，长度32
	AcctName string `json:"acctName"`
	// 收款账户类型代码，必传，长度32，取值说明：57：对公，58：对私
	AcctTypeCode consts.AcctTypeCode `json:"acctTypeCode"`
	// 收款账户证件类型，必传，长度32，取值说明：17身份证，18护照，19港澳居民来往内地通行证，20台湾居民来往内地通行证，身份证外类型先咨询后再使用
	AcctCertificateType string `json:"acctCertificateType"`
	// 收款账户证件号，必传，长度32
	AcctCertificateNo string `json:"acctCertificateNo"`
	// 收款账户开户行号，必传，长度32，取值说明：参照FBI.N信息查询，仅支持对私结算账户
	AcctOpenBankCode string `json:"acctOpenBankCode"`
	// 收款账户开户行名称，必传，长度64，取值说明：参照FBI.N信息查询，仅支持对私结算账户
	AcctOpenBankName string `json:"acctOpenBankName"`
	// 收款账户清算行行号，必传，长度32，取值说明：参照FBI.N信息查询，仅支持对私结算账户
	AcctClearBankCode string `json:"acctClearBankCode"`
	// 接收方附件资料，可选，集合
	AttachList []*ApplyBindAttachment `json:"attachList"`
	// 提款类型，可选，长度32，取值说明：01：主动提款，03：交易自动结算，不填默认01
	SettleType string `json:"settleType"`
}

// ApplyLedgerReceiverResponse 分账接收方创建响应结构体
// 拉卡拉接口返回的分账接收方创建响应数据
// 包含响应状态码、消息和业务数据
type ApplyLedgerReceiverResponse struct {
	// 响应状态码，000000表示成功
	RetCode string `json:"retCode"`
	// 响应消息
	RetMsg string `json:"retMsg"`
	// 响应业务数据，当retCode为000000时返回
	RespData *ApplyLedgerReceiverRespData `json:"respData"`
}

// ApplyLedgerReceiverRespData 分账接收方创建响应业务数据结构体
// 包含分账接收方创建返回的具体业务信息
type ApplyLedgerReceiverRespData struct {
	// 接口版本号(回传)
	Version string `json:"version"`
	// 订单编号(回传)
	OrderNo string `json:"orderNo"`
	// 申请机构代码(回传)
	OrgCode string `json:"orgCode"`
	// 接收方所属机构
	OrgId string `json:"orgId"`
	// 接收方所属机构名称
	OrgName string `json:"orgName"`
	// 接收方编号
	ReceiverNo string `json:"receiverNo"`
}

// SuccessOrFail 判断分账接收方创建请求是否成功
// 成功条件：响应码为000000
// 返回值：true表示成功，false表示失败
func (resp *ApplyLedgerReceiverResponse) SuccessOrFail() bool {
	return resp.RetCode == "000000"
}
