package model

// ApplyLedgerMerRequest 商户分账业务开通申请请求结构体
type ApplyLedgerMerRequest struct {
	ReqData *ApplyLedgerMerReqData `json:"reqData"` // 请求业务数据
	ReqId   string                 `json:"reqId"`   // 请求ID
	Version string                 `json:"version"` // 接口版本号
	ReqTime string                 `json:"reqTime"` // 请求时间
}

// ApplyLedgerMerReqData 商户分账业务开通申请请求业务数据结构体
type ApplyLedgerMerReqData struct {
	Version              string  `json:"version"`              // 版本号，必传，长度8，取值说明：1.0
	OrderNo              string  `json:"orderNo"`              // 订单编号，必传，长度32，用于后续处理查询及回调通知消息标识，2014年月日时分秒毫秒组成
	OrgCode              string  `json:"orgCode"`              // 机构代码，必传，长度12
	MerInnerNo           string  `json:"merInnerNo"`           // 拉卡拉内部商户号，可选，长度32，拉卡拉内部商户号和银联商户号必须传一个，默认以内部商户号为准
	MerCupNo             string  `json:"merCupNo"`             // 银联商户号，可选，长度32，拉卡拉内部商户号和银联商户号必须传一个，默认以内部商户号为准
	ContactMobile        string  `json:"contactMobile"`        // 联系手机号，必传，长度32
	SplitLowestRatio     float64 `json:"splitLowestRatio"`     // 最低分账比例，必传，长度12，百分比，支持2位精度，取值范围：70-70.50
	SplitEntrustFileName string  `json:"splitEntrustFileName"` // 分账授权委托书文件名称，必传，长度64，文件格式：pdf
	SplitEntrustFilePath string  `json:"splitEntrustFilePath"` // 分账授权委托书文件路径，必传，长度64，调用附件上传接口获取
	SplitRange           string  `json:"splitRange"`           // 分账范围，必传，长度32，取值说明：ALL-全部交易分账(所有交易默认都分账)，MARK-标记交易分账(只有带标记交易才分账，其余交易正常结算)
	SplitFundSource      string  `json:"splitFundSource"`      // 分账依据，非必传，长度32，取值说明：TRA-交易分账，BAR-金额分账
	ElecContractId       string  `json:"elecContractId"`       // 电子合同编号，非必传，长度32，收单已签约交易电子合同编号，供审核人员复核使用
	SplitLaunchMode      string  `json:"splitLaunchMode"`      // 分账发起方式，非必传，长度32，取值说明：AUTO-自动触发分账，POINTTRUE-指定规则分账，MANUAL-手动分账
	SettleType           string  `json:"settleType"`           // 结算类型，非必传，长度32，取值说明：01-主扫现结，02-复扫现结，03-交易自动结算
	SplitRuleSource      string  `json:"splitRuleSource"`      // 分账规则来源，条件必传，长度32，取值说明：MER-商户自定规则，PLATFORM-平台分润规则(分润规则必传)
	RetUrl               string  `json:"retUrl"`               // 回调通知地址，必传，长度128，分账申请结果以异步消息或同步返回的方式通知，如需无线路由处理，也可以通过第三方商户信息查询接口确定结算结果
	Attachments          []struct {
		AttachType      string `json:"attachType"`      // 附件类型编码，必传，长度32
		AttachName      string `json:"attachName"`      // 附件名称，必传，长度32
		AttachStorePath string `json:"attachStorePath"` // 附件路径，必传，长度128，调用附件上传接口获取
	} `json:"attachments,omitempty"` // 附加资料，可选，集合，其他需附加的文件信息
}

// ApplyLedgerMerResponse 商户分账业务开通申请响应结构体
type ApplyLedgerMerResponse struct {
	RetCode  string `json:"retCode"` // 返回码，000000表示成功
	RetMsg   string `json:"retMsg"`  // 返回消息
	RespData struct {
		Version string `json:"version"` // 接口版本号，例如：547110502170558464
		OrderNo string `json:"orderNo"` // 订单编号，例如：2021020112000012345678
		OrgCode string `json:"orgCode"` // 机构代码，例如：200669
		ApplyId int64  `json:"applyId"` // 受理编号，例如：548099616395186176
	} `json:"respData"` // 响应数据
}

func (a *ApplyLedgerMerResponse) SuccessOrFail() bool {
	return a.RetCode == "000000"
}
