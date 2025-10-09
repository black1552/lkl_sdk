package model

type SeparateRequest struct {
	ReqData *SeparateReqData `json:"req_data"` // 请求数据
	Version string           `json:"version"`  // 版本号
	ReqTime string           `json:"req_time"` // 请求时间
}

type SeparateReqData struct {
	MerchantNo    string               `json:"merchant_no"`     // 商户号
	LogDate       string               `json:"log_date"`        // 交易日期 yyyyMMdd，查清结算用
	LogNo         string               `json:"log_no"`          // 拉卡拉对账单流水号
	OutSeparateNo string               `json:"out_separate_no"` // 商户分账指令流水号
	TotalAmt      string               `json:"total_amt"`       // 分账总金额 [单位：分]
	LklOrgNo      string               `json:"lkl_org_no"`      // 拉卡拉机构编号 非必填
	CalType       string               `json:"cal_type"`        // 分账计算类型 0- 按照指定金额，1- 按照指定比例。默认 0 非必填
	SeparateType  string               `json:"separate_type"`
	NotifyUrl     string               `json:"notify_url"` // 回调地址 分账，分账撤销或分账回退时，是异步接口。通过该地址通知商户最终处理结果。不传时，不回调
	RecvDatas     []*SeparateRecvDatas `json:"recv_datas"` // 分账接收数据对象 分账接收方编号必须已创建
}
type SeparateRecvDatas struct {
	RecvMerchantNo string `json:"recv_merchant_no,omitempty"` // 接收方商户号 分账接收方商户号，分给自己时使用，需和merchantNo一样，否则检查报错；分账接收方商户号 和 分账接收方 只能填写一个。
	SeparateValue  string `json:"separate_value"`             // 分账数值 calType为0时，按照固定金额分账，单位：分 calType为1时，按照比例分账，单位：百分比的小数值，比如0.55 （55%）
	RecvNo         string `json:"recv_no,omitempty"`          // 分账接收方编号 分账接收方编号, 分给他人时使用；分账接收方商户号 和 分账接收方 只能填写一个。
}

type SeparateResponse struct {
	Msg      string `json:"msg"`
	RespTime string `json:"resp_time"`
	Code     string `json:"code"`
	RespData struct {
		TotalAmt      string `json:"total_amt"`       // 分账总金额 单位：分
		LogDate       string `json:"log_date"`        // 拉卡拉订单日期 posp日期，yyyyMMdd，查清结算用
		SeparateNo    string `json:"separate_no"`     // 分账指令流水号 分账系统生成唯一流水
		LogNo         string `json:"log_no"`          // 拉卡拉对账单流水号 请求透返
		OutSeparateNo string `json:"out_separate_no"` // 商户订单号 请求报文中的商户外部订单号
		Status        string `json:"status"`          // 分账状态 处理中：PROCESSING, 已受理：ACCEPTED, 成功：SUCCESS, 失败：FAIL
	} `json:"resp_data"`
}

func (s *SeparateResponse) SuccessOrFail() bool {
	return s.Code == "SACS0000"
}
