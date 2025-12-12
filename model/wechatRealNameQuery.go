package model

type WechatRealNameQueryRequest struct {
	// 请求业务数据
	ReqData *WechatRealNameQueryReqData `json:"reqData"`
	// 接口版本号
	Ver string `json:"ver"`
	// 请求时间，格式为yyyyMMddHHmmss
	Timestamp int64  `json:"timestamp"`
	ReqId     string `json:"reqId"`
}

type WechatRealNameQueryReqData struct {
	Version    string `json:"version" dc:"接口版本号"`
	OrderNo    string `json:"orderNo" dc:"订单编号 14位年月日时（24小时制）分秒+8位的随机数（不重复）如：2021020112000012345678"`
	OrgCode    string `json:"orgCode" dc:"机构代码"`
	MerInnerNo string `json:"merInnerNo" dc:"拉卡拉内部商户号"`
	SubMchId   string `json:"subMchId" dc:"子商户号"`
	ChannelId  string `json:"channelId" dc:"渠道号 （建议传入，能具体定位用的渠道，仅支持拉卡拉渠道查询）"`
}

type WechatRealNameQueryResponse struct {
	// 响应状态码，000000表示成功
	RetCode string `json:"retCode"`
	// 响应消息
	RetMsg string `json:"retMsg"`
	// 响应业务数据，当code为000000时返回
	RespData *WechatRealNameQueryRespData `json:"respData"`
}

type WechatRealNameQueryRespData struct {
	MerInnerNo      string `json:"merInnerNo" dc:"拉卡拉内部商户号"`
	SubMchId        string `json:"subMchId" dc:"账户端子商户号	"`
	ChannelId       string `json:"channelId" dc:"渠道号"`
	ReceOrgNo       string `json:"receOrgNo" dc:"从业机构号"`
	ApplymentId     string `json:"applymentId" dc:"申请编号"`
	ApplymentState  string `json:"applymentState" dc:"申请状态"`
	AuthorizeState  string `json:"authorizeState" dc:"认证状态"`
	QrcodeData      string `json:"qrcodeData" dc:"	小程序码图片"`
	RejectParameter string `json:"rejectParameter" dc:"驳回参数"`
	RejectReason    string `json:"rejectReason" dc:"驳回原因"`
}

func (resp *WechatRealNameQueryResponse) SuccessOrFail() bool {
	return resp.RetCode == "000000"
}
