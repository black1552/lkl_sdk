package model

type QuerySubMerInfoRequest struct {
	ReqData   *QuerySubMerInfoReqData `json:"reqData"`
	Ver       string                  `json:"ver"`
	Timestamp int64                   `json:"timestamp"`
	ReqId     string                  `json:"reqId"`
}

type QuerySubMerInfoReqData struct {
	Version         string `json:"version" dc:"接口版本号"`
	OrderNo         string `json:"orderNo" dc:"订单编号,保证唯一"`
	OrgCode         string `json:"orgCode" dc:"机构代码"`
	MerInnerNo      string `json:"merInnerNo" dc:"拉卡拉内部商户号和银联商户号必须传一个，都送以内部商户号为准。"`
	MerCupNo        string `json:"merCupNo" dc:"拉卡拉内部商户号和银联商户号必须传一个，都送以内部商户号为准。"`
	RegisterChannel string `json:"registerChannel" dc:"报备渠道"`
	RegisterType    string `json:"registerType" dc:"报备类型"`
	RegisterStatus  string `json:"registerStatus" dc:"报备状态 SUCCESS：成功；FAIL：失败"`
	SubMchId        string `json:"subMchId" dc:"子商户号"`
}

type QuerySubMerInfoResponse struct {
	RetCode  string                `json:"retCode"`
	RetMsg   string                `json:"retMsg"`
	RespData *QueryMerResponseData `json:"respData"`
}

type QuerySubMerInfoRespData struct {
	OrgCode string          `json:"orgCode"`
	OrderNo string          `json:"orderNo"`
	List    []*RegisterList `json:"list"`
}
type RegisterList struct {
	MerInnerNo      string `json:"merInnerNo" dc:"内部商户号"`
	SubMchId        string `json:"subMchId" dc:"子商户号"`
	SubMchIdBank    string `json:"subMchIdBank" dc:"交易子商户号"`
	DcWalletId      string `json:"dcWalletId" dc:"数币钱包ID"`
	ChannelId       string `json:"channelId" dc:"渠道号"`
	ReceOrgNo       string `json:"receOrgNo" dc:"从业机构号"`
	RegisterChannel string `json:"registerChannel" dc:"报备渠道"`
	RegisterType    string `json:"registerType" dc:"报备类型"`
	RegisterTm      string `json:"registerTm" dc:"报备时间"`
	RegisterStatus  string `json:"registerStatus" dc:"报备状态"`
	ResultCode      string `json:"resultCode" dc:"结果返回码"`
	ResultMessage   string `json:"resultMessage" dc:"结果描述"`
}

func (t *QuerySubMerInfoResponse) SuccessOrFail() bool {
	return t.RetCode == "000000"
}
