package model

type OrderSplitLedgerResultQueryRequest struct {
	ReqData *OrderSplitLedgerResultQueryReqData `json:"req_data"` // 请求业务数据
	Version string                              `json:"version"`  // 接口版本号
	ReqTime string                              `json:"req_time"` // 请求时间，格式为yyyyMMddHHmmss
}

type OrderSplitLedgerResultQueryReqData struct {
	MerchantNo    string `json:"merchant_no"`     // 商户号，必传，长度32
	SeparateNo    string `json:"separate_no"`     // 分账核心系统返回的分账流水号，入参中的separate_no,out_separate_no至少二选一，优先级: separate_no> out_separate_no
	OutSeparateNo string `json:"out_separate_no"` // 商户分账指令流水号，入入参中的separate_no,out_separate_no至少二选一，优先级: separate_no> out_separate_no
}

type OrderSplitLedgerResultQueryResponse struct {
	Msg      string `json:"msg"`       // 消息
	RespTime string `json:"resp_time"` // 响应时间
	Code     string `json:"code"`      // 响应码 SACS0000表示成功
	RespData struct {
		SeparateNo        string                                        `json:"separate_no"`         // 分账指令流水号
		OutSeparateNo     string                                        `json:"out_separate_no"`     // 商户分账指令流水号
		CmdType           string                                        `json:"cmd_type"`            // 指令类型 SEPARATE：分账 CANCEL：分账撤销 FALLBACK：分账回退
		LogNo             string                                        `json:"log_no"`              // 拉卡拉对账单流水号
		LogDate           string                                        `json:"log_date"`            // 交易日期  posp日期，yyyyMMdd，查清结算用
		CalType           string                                        `json:"cal_type"`            // 分账计算类型 0 按照指定金额。1 按照指定比例，默认 0 （cmd_type为SEPARATE分账指令类型才有值）
		SeparateDate      string                                        `json:"separate_date"`       // 分账日期
		FinishDate        string                                        `json:"finish_date"`         // 完成日期
		TotalAmt          string                                        `json:"total_amt"`           // 发生总金额，单位为分
		Status            string                                        `json:"status"`              // 分账状态，ACCEPTED:已受理, PROCESSING:处理中, FAIL:失败, SUCCESS:成功, （如果分账指令后有反向操作指令，则原分账指令会变更成以下的状态之一：） CANCELING:撤销中, CANCELED:撤销成功, CANCEL_FAIL:撤销失败, FALLBACKING:回退中, FALLBACK_END:回退结束
		FinalStatus       string                                        `json:"final_status"`        // 处理状态 ACCEPTED:已受理, PROCESSING:处理中, FAIL:失败, SUCCESS:成功
		FrontRuleId       string                                        `json:"front_rule_id"`       // 分账前置规则ID
		ActualSeparateAmt string                                        `json:"actual_separate_amt"` // 实分金额
		TotalFeeAmt       string                                        `json:"total_fee_amt"`       // 手续费金额
		AccResultDesc     string                                        `json:"acc_result_desc"`     // 账户处理错误描述
		DetailDatas       []*OrderSplitLedgerResultQueryRespDetailDatas `json:"detail_datas"`        // 明细数据
	}
}

type OrderSplitLedgerResultQueryRespDetailDatas struct {
	RecvMerchantNo string `json:"recv_merchant_no"` // 接收方商户号
	RecvNo         string `json:"recv_no"`          // 接收方编号
	Amt            string `json:"amt"`              // 分账金额
	ActualAmt      string `json:"actual_amt"`       // 实分金额
	FeeAmt         string `json:"fee_amt"`          // 手续费金额
}

func (s *OrderSplitLedgerResultQueryResponse) SuccessOrFail() bool {
	return s.Code == "SACS0000"
}
