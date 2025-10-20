package querystatus

// ECQueryStatusResponseData 电子合同查询状态响应数据
type ECQueryStatusResponseData struct {
	OrderNo   string `json:"order_no" dc:"请求上送的订单号"`               // 请求上送的订单号
	EcApplyID int64  `json:"ec_apply_id" dc:"电子签约申请受理编号"`          // 电子签约申请受理编号
	EcStatus  string `json:"ec_status" dc:"电子合同签署状态"`               // 电子合同签署状态，UNDONE 未完成，COMPLETED 已完成
	EcNo      string `json:"ec_no" dc:"电子合同号"`                       // 签署完成反馈
}

// ECQueryStatusResponse 电子合同查询状态响应
type ECQueryStatusResponse struct {
	Code     string                     `json:"code" dc:"返回码"`         // 成功 000000 其它失败
	Msg      string                     `json:"msg" dc:"返回码描述"`        // 返回码描述
	RespData *ECQueryStatusResponseData `json:"resp_data" dc:"结果信息集合"` // 结果信息集合
	RespTime string                     `json:"resp_time" dc:"响应时间"`   // 格式：yyyyMMddHHmmss
}

// SuccessOfFail 检查响应是否成功
func (r *ECQueryStatusResponse) SuccessOfFail() bool {
	return r.Code == "000000"
}