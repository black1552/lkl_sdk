package model

// SplitBalanceRequest 可分账金额查询请求结构体
// 用于查询指定交易的可分账金额信息
// 拉卡拉SDK接口文档：可分账金额查询接口

type SplitBalanceRequest struct {
	ReqData *SplitBalanceReqData `json:"req_data"` // 请求业务数据
	Version string               `json:"version"`   // 接口版本号
	ReqTime string               `json:"req_time"`  // 请求时间，格式为yyyyMMddHHmmss
}

// SplitBalanceReqData 可分账金额查询请求业务数据结构体
// 包含可分账金额查询所需的查询参数

type SplitBalanceReqData struct {
	MerchantNo string `json:"merchant_no"` // 商户号，必传，长度32
	LogDate    string `json:"log_date"`    // 拉卡拉对账单交易日期，必传，长度8，格式为yyyyMMdd，用于查询结算
	LogNo      string `json:"log_no"`      // 拉卡拉对账单流水号，必传，长度14，对应pos流水号，用于查询结算
}
// SplitBalanceResponse 可分账金额查询响应结构体
// 包含可分账金额查询的结果信息

type SplitBalanceResponse struct {
	Msg      string `json:"msg"`       // 响应消息，描述响应结果
	RespTime string `json:"resp_time"` // 响应时间，格式为yyyyMMddHHmmss
	Code     string `json:"code"`      // 响应码，SACS0000表示成功
	RespData struct {
		MerchantNo       string `json:"merchant_no"`       // 商户号，请求返回
		TotalSeparateAmt string `json:"total_separate_amt"` // 总分账金额，必传，长度15，单位为分
		CanSeparateAmt   string `json:"can_separate_amt"`   // 可分账金额，必传，长度15，单位为分
		LogDate          string `json:"log_date"`          // 拉卡拉对账单交易日期，必传，长度8，格式为yyyyMMdd
		LogNo            string `json:"log_no"`            // 拉卡拉对账单流水号，必传，长度14，请求返回
	} `json:"resp_data"` // 响应业务数据，当code为SACS0000时返回
}

func (s *SplitBalanceResponse) SuccessOrFail() bool {
	return s.Code == "SACS0000"
}
