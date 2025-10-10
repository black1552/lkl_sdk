package model

import (
	"github.com/black1552/lkl_sdk/consts"
)

// MerchantApplyRequest 商户进件请求结构体

type MerchantApplyRequest struct {
	ReqId     string                `json:"reqId"`     // 请求时间，格式为yyyyMMddHHmmss，必填
	Timestamp int64                 `json:"timestamp"` // 接口版本，固定值"1.0"，必填
	Ver       string                `json:"ver"`       // 接口版本，固定值"1.0"，必填
	ReqData   *MerchantApplyReqData `json:"req_data"`  // 请求业务参数，必填
}

// MerchantApplyReqData 商户进件请求业务数据结构体

type MerchantApplyReqData struct {
	Version           string                `json:"version"`           // 接口版本号，固定值"1.0"，必填
	OrderNo           string                `json:"orderNo"`           // 订单编号，用于后续联调时查询及核对报文，必填，建议14位年月日时分秒+8位随机数
	PosType           consts.PosType        `json:"posType"`           // POS类型，必填，按接入系统做控制，参见POS类型字典表
	OrgCode           string                `json:"orgCode"`           // 机构代码，合作方在拉卡拉的标识，请联系业务员，必填
	MerRegName        string                `json:"merRegName"`        // 商户注册名称，必填，长度不小于4个汉字，8~40字符，不可为纯数字
	MerBizName        string                `json:"merBizName"`        // 商户经营名称，选填，为空时同商户注册名称，4~64字符，不可为纯数字
	MerRegDistCode    consts.AddrCode       `json:"merRegDistCode"`    // 商户地区代码，必填，参见地区文档
	MerRegAddr        string                `json:"merRegAddr"`        // 商户详细地址，必填，去除省、市、区后的详细地址，6-200字符
	MccCode           consts.Mcc            `json:"mccCode"`           // 商户MCC编号，必填，银联商户类别代码
	MerBlisName       string                `json:"merBlisName"`       // 营业执照名称，选填，小微商户可不传，其它必传
	MerBlis           string                `json:"merBlis"`           // 营业执照号，选填，小微商户可不传，对公进件必传，且不可与法人证件相同
	MerBlisStDt       string                `json:"merBlisStDt"`       // 营业执照开始日期，选填，格式yyyy-MM-dd，有营业执照时必传
	MerBlisExpDt      string                `json:"merBlisExpDt"`      // 营业执照有效期，选填，格式yyyy-MM-dd，有营业执照时必传
	MerBusiContent    consts.MerBusiContent `json:"merBusiContent"`    // 商户经营内容，必填，参见经营内容字典表
	LarName           string                `json:"larName"`           // 商户法人姓名，必填
	LarIdType         consts.AccIdType      `json:"larIdType"`         // 法人证件类型，必填，支持其他证件类型，见证件类型字典表
	LarIdcard         string                `json:"larIdcard"`         // 法人身份证号码，必填
	LarIdcardStDt     string                `json:"larIdcardStDt"`     // 法人身份证开始日期，必填，格式yyyy-MM-dd
	LarIdcardExpDt    string                `json:"larIdcardExpDt"`    // 法人身份证有效期，必填，格式yyyy-MM-dd
	MerContactMobile  string                `json:"merContactMobile"`  // 商户联系人手机号码，必填
	MerContactName    string                `json:"merContactName"`    // 商户联系人姓名，必填
	ShopName          string                `json:"shopName"`          // 网点名称，选填，不填取商户注册名称
	ShopDistCode      string                `json:"shopDistCode"`      // 网点地址区划代码，选填，不填取商户地区代码
	ShopAddr          string                `json:"shopAddr"`          // 网点详细地址，选填，不填取商户详细地址
	ShopContactName   string                `json:"shopContactName"`   // 网点联系人名称，选填，不填取商户联系人姓名
	ShopContactMobile string                `json:"shopContactMobile"` // 网点联系人手机号，选填，不填取商户联系人手机号码
	OpenningBankCode  string                `json:"openningBankCode"`  // 结算账户开户行号，必填，可根据结算卡信息进行查询
	OpenningBankName  string                `json:"openningBankName"`  // 结算账户开户行名称，必填，可根据结算卡信息进行查询
	ClearingBankCode  string                `json:"clearingBankCode"`  // 结算账户清算行号，必填，可根据结算卡信息进行查询
	AcctNo            string                `json:"acctNo"`            // 结算账户账号，必填
	AcctName          string                `json:"acctName"`          // 结算账户名称，必填
	AcctTypeCode      consts.AcctTypeCode   `json:"acctTypeCode"`      // 结算账户性质，必填，57为对公，58为对私
	SettlePeriod      consts.SettlePeriod   `json:"settlePeriod"`      // 结算周期，必填，参见结算周期表
	ClearDt           consts.ClearDt        `json:"clearDt"`           // 日切时间，选填，参见日切时间字典表，默认TWENTY_THREE
	AcctIdType        consts.AccIdType      `json:"acctIdType"`        // 结算人证件类型，选填，为空时判断为同法人
	AcctIdcard        string                `json:"acctIdcard"`        // 结算人证件号码，选填，为空时判断为同法人
	AcctIdDt          string                `json:"acctIdDt"`          // 结算人证件有效期，选填，为空时判断为同法人
	DevSerialNo       string                `json:"devSerialNo"`       // 终端设备序列号，选填
	DevTypeName       string                `json:"devTypeName"`       // 设备型号，选填
	TermVer           string                `json:"termVer"`           // 终端版本号，选填
	SalesStaff        string                `json:"salesStaff"`        // 销售人员，选填
	TermNum           string                `json:"termNum"`           // 终端数量，选填，1-5，最大5个终端
	RetUrl            string                `json:"retUrl"`            // 回调地址，必填
	FeeData           []*FeeData            `json:"feeData"`           // 费率信息集合，必填
	FileData          []*FileData           `json:"fileData"`          // 附件集合，选填
	ContractNo        string                `json:"contractNo"`        // 电子合同编号，选填，部分进件类型要求录入
	FeeAssumeType     string                `json:"feeAssumeType"`     // 大额理财-手续费承担方，选填，大额理财进件时必传
	AmountOfMonth     string                `json:"amountOfMonth"`     // 大额理财-最小月交易额，选填
	ServiceFee        string                `json:"serviceFee"`        // 大额理财-收取服务费，选填
}

// FeeData 费率信息集合结构体

type FeeData struct {
	FeeRateTypeCode consts.FeeRateTypeCode `json:"feeRateTypeCode"` // 费率类型，必填，参见费率类型字典表
	FeeRateTypeName string                 `json:"feeRateTypeName"` // 费率类型名称，必填，如银行卡借记卡
	FeeRatePct      string                 `json:"feeRatePct"`      // 手续费费率(%)，必填，如0.6
	FeeUpperAmtPcnt string                 `json:"feeUpperAmtPcnt"` // 单笔交易手续费封顶，选填，默认不封顶，单位（元）
	FeeLowerAmtPcnt string                 `json:"feeLowerAmtPcnt"` // 单笔交易手续费保底，选填，默认无保底，单位（元）
	FeeRateSdt      string                 `json:"feeRateSdt"`      // 手续费生效日期，选填，默认为进件日期
}

// FileData 附件集合结构体

type FileData struct {
	AttField string `json:"attField"` // 文件编号/附件上传后返回的编号，必填
	AttType  string `json:"attType"`  // 附件类型，必填
}

// MerchantApplyResponse 商户进件响应结构体

type MerchantApplyResponse struct {
	RetCode    string                 `json:"retCode"`    // 响应码，成功为"0000"，其他为错误码
	RetMsg     string                 `json:"retMsg"`     // 响应描述
	Timestamp  int64                  `json:"timestamp"`  // 响应描述
	Ver        string                 `json:"ver"`        // 响应描述
	ReqId      string                 `json:"reqId"`      // 响应描述
	CmdRetCode string                 `json:"cmdRetCode"` // 响应描述
	RespData   *MerchantApplyRespData `json:"respData"`   // 响应业务数据，成功时返回
}

// MerchantApplyRespData 商户进件响应业务数据结构体

type MerchantApplyRespData struct {
	OrgCode    string `json:"orgCode"`    // 机构代码
	OrderNo    string `json:"orderNo"`    // 订单号
	ContractId string `json:"contractId"` // 进件ID，用于后续查询进件结果
}

func (t *MerchantApplyResponse) SuccessOrFail() bool {
	return t.RetCode == "000000"
}
