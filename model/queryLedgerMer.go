package model

// QueryLedgerMerRequest 分账商户查询请求结构体
// 用于向拉卡拉接口发送分账商户查询请求
// 包含请求头信息和业务数据
type QueryLedgerMerRequest struct {
	// 请求业务数据
	ReqData *QueryLedgerMerReqData `json:"reqData"`
	// 接口版本号
	Version string `json:"version"`
	// 请求时间，格式为yyyyMMddHHmmss
	ReqTime string `json:"reqTime"`
}

// QueryLedgerMerReqData 分账商户查询请求业务数据结构体
// 包含分账商户查询所需的具体业务参数
// 用于查询商户的分账设置信息
type QueryLedgerMerReqData struct {
	// 接口版本号，必传，长度8，取值说明：1.0
	Version string `json:"version"`
	// 订单编号（便于后续跟踪排查问题及核对报文），必传，长度32，取值说明：14位年月日（24小时制）分秒+8位的随机数（不重复）
	OrderNo string `json:"orderNo"`
	// 机构代码，必传，长度32
	OrgCode string `json:"orgCode"`
	// 拉卡拉内部商户号，可选，长度32，取值说明：拉卡拉内部商户号和银联商户号必须传一个，都送以内部商户号为准
	MerInnerNo string `json:"merInnerNo"`
	// 银联商户号，可选，长度32，取值说明：拉卡拉内部商户号和银联商户号必须传一个，都送以内部商户号为准
	MerCupNo string `json:"merCupNo"`
}

// QueryLedgerMerResponse 分账商户查询响应结构体
// 拉卡拉接口返回的分账商户查询响应数据
// 包含响应状态码、消息和业务数据
type QueryLedgerMerResponse struct {
	// 响应状态码，000000表示成功
	Code string `json:"code"`
	// 响应消息
	Msg string `json:"msg"`
	// 响应业务数据，当code为000000时返回
	RespData *QueryLedgerMerRespData `json:"respData"`
}

// QueryLedgerMerRespData 分账商户查询响应业务数据结构体
// 包含分账商户查询返回的具体业务信息
type QueryLedgerMerRespData struct {
	// 分账商户机构号
	OrgId string `json:"orgId"`
	// 分账商户机构名称
	OrgName string `json:"orgName"`
	// 拉卡拉内部商户号
	MerInnerNo string `json:"merInnerNo"`
	// 银联商户号
	MerCupNo string `json:"merCupNo"`
	// 最低分账比例（百分比，支持2位精度），取值说明：70或70.50
	SplitLowestRatio string `json:"splitLowestRatio"`
	// 商户分账状态，取值说明：VALID启用，INVALID禁用
	SplitStatus string `json:"splitStatus"`
	// 分账范围，取值说明：ALL：全部交易分账(商户所有交易默认待分账)，MARK：标记交易分账(只有带分账标识交易待分账，其余交易正常结算)，默认：MARK
	SplitRange string `json:"splitRange"`
	// 分账依据，取值说明：TR或空：交易分账，BA：余额分账，默认：TR交易分账
	SepFundSource string `json:"sepFundSource"`
	// 平台ID，取值说明：如果商户和绑定平台分账，返回平台ID
	PlatformId string `json:"platformId"`
	// 分账发起方式，取值说明：AUTO：自动规则分账，POINTRULE：指定规则分账，MANUAL：手动规则分账
	SplitLaunchMode string `json:"splitLaunchMode"`
	// 分账规则来源，取值说明：MER：商户分账规则，PLATFORM：平台分账规则
	SplitRuleSource string `json:"splitRuleSource"`
	// 已绑定接收方列表
	BindRelations []*BindRelation `json:"bindRelations"`
}

// BindRelation 已绑定接收方信息结构体
// 用于表示分账商户已绑定的接收方信息
type BindRelation struct {
	// 拉卡拉内部商户号
	MerInnerNo string `json:"merInnerNo"`
	// 银联商户号
	MerCupNo string `json:"merCupNo"`
	// 接收方编号
	ReceiverNo string `json:"receiverNo"`
	// 接收方编号名称
	ReceiverName string `json:"receiverName"`
}

// SuccessOrFail 判断分账商户查询请求是否成功
// 成功条件：响应码为000000
// 返回值：true表示成功，false表示失败
func (resp *QueryLedgerMerResponse) SuccessOrFail() bool {
	return resp.Code == "000000"
}
