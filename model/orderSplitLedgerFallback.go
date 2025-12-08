package model

type OrderSplitLedgerFallbackRequest struct {
	ReqData *OrderSplitLedgerFallbackReqData `json:"req_data"` // 请求业务数据
	Version string                           `json:"version"`  // 接口版本号
	ReqTime string                           `json:"req_time"` // 请求时间，格式为yyyyMMddHHmmss
}

type OrderSplitLedgerFallbackReqData struct {
	MerchantNo          string                             `json:"merchant_no"`            // 商户号，必传，长度32
	OriginSeparateNo    string                             `json:"origin_separate_no"`     // 原分账单号，必传，长度32
	OutSeparateNo       string                             `json:"out_separate_no"`        // 外部分账单号，必传，长度32
	OriginOutSeparateNo string                             `json:"origin_out_separate_no"` // 原外部分账单号，必传，长度32
	FallbackReason      string                             `json:"fallback_reason"`        // 回退原因，必传，长度255
	TotalAmt            string                             `json:"total_amt"`              // 总金额，必传，长度15
	OriginRecvDatas     []*OrderSplitLedgerOriginRecvDatas `json:"origin_recv_datas"`      // 原分账接收数据，必传，数组长度1-100
}

type OrderSplitLedgerOriginRecvDatas struct {
	RecvNo string `json:"recv_no"` // 原分账接收号，必传，长度32
	Amt    string `json:"amt"`     // 原分账接收金额，必传，长度15
}

type OrderSplitLedgerFallbackResponse struct {
	Msg      string `json:"msg"`       // 消息
	RespTime string `json:"resp_time"` // 响应时间
	Code     string `json:"code"`      // 响应码 SACS0000表示成功
	RespData struct {
		OutSeparateNo       string `json:"out_separate_no"`        // 外部分账单号，必传，长度32
		TotalAmt            string `json:"total_amt"`              // 总金额，必传，长度15
		OriginOutSeparateNo string `json:"origin_out_separate_no"` // 原外部分账单号，必传，长度32
		OriginSeparateNo    string `json:"origin_separate_no"`     // 原分账单号，必传，长度32
		Status              string `json:"status"`                 // 状态，必传，长度1
		SeparateNo          string `json:"separate_no"`            // 分账单号，必传，长度32
	}
}

func (s *OrderSplitLedgerFallbackResponse) SuccessOrFail() bool {
	return s.Code == "SACS0000"
}
