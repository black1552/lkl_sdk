package model

// BalanceQueryRequest 余额查询请求结构体
// 用于向拉卡拉接口发送余额查询请求
// 包含请求头信息和业务数据
type BalanceQueryRequest struct {
	// 请求业务数据
	ReqData *BalanceQueryReqData `json:"req_data"`
	// 接口版本号
	Version string `json:"version"`
	// 请求时间，格式为yyyyMMddHHmmss
	ReqTime string `json:"req_time"`
}

// BalanceQueryReqData 余额查询请求业务数据结构体
// 包含余额查询所需的具体业务参数
// 用于查询账户的实时余额信息
type BalanceQueryReqData struct {
	// 机构号，必传，最大长度32
	OrgNo string `json:"org_no"`
	// 商户号或receiveNo或商户用户编号，必传，最大长度32
	MerchantNo string `json:"merchant_no"`
	// 账号，若该参数上送，则payType将无效，非必传，最大长度32
	PayNo string `json:"pay_no"`
	// 账号类型（01：收款账户，02：付款账户，03：分账商户账户，04：分账接收方账户，05：充值代付账户，06：结算代付账户）-未上送则默认01，非必传，最大长度32
	PayType string `json:"pay_type"`
	// 账户标志（01:一般账户;03:虚户）-未上送则默认01，非必传，最大长度32
	MgtFlag string `json:"mgt_flag"`
}

// BalanceQueryResponse 余额查询响应结构体
// 拉卡拉接口返回的余额查询响应数据
// 包含响应状态码、消息和业务数据
type BalanceQueryResponse struct {
	// 响应状态码，000000
	Code string `json:"code"`
	// 响应消息
	Msg string `json:"msg"`
	// 响应业务数据，当code为SACS0000时返回
	RespData *BalanceQueryRespData `json:"resp_data"`
}

// BalanceQueryRespData 余额查询响应业务数据结构体
// 包含余额查询返回的具体账户信息
type BalanceQueryRespData struct {
	// 账号，必传
	PayNo string `json:"pay_no"`
	// 账户类型，必传
	PayType string `json:"pay_type"`
	// 账户状态，必传，取值说明：CLOSE销户，NORMAL正常，FREEZE冻结，STOPAY止付
	AcctSt string `json:"acct_st"`
	// 预付余额（单位元），必传
	ForceBalance string `json:"force_balance"`
	// 上日余额（单位元）-该字段已废弃使用，必传
	HisBalance string `json:"his_balance"`
	// 实时余额（单位元），必传
	ReBalance string `json:"re_balance"`
	// 当前可用余额（单位元），必传
	CuBalance string `json:"cu_balance"`
}

// SuccessOrFail 判断余额查询请求是否成功
// 成功条件：响应码为SACS0000
// 返回值：true表示成功，false表示失败
func (resp *BalanceQueryResponse) SuccessOrFail() bool {
	return resp.Code == "000000"
}
