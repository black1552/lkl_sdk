package apply

// ECApplyResponseData 电子合同申请响应数据
type ECApplyResponseData struct {
	OrderNo   string `json:"order_no" dc:"请求上送的订单号"`         // 请求上送的订单号
	EcApplyID int64  `json:"ec_apply_id" dc:"电子签约申请受理编号"`    // 电子签约申请受理编号
	ResultUrl string `json:"result_url" dc:"电子签约申请结果H5链接地址"` // 申请成功时：待签约合同H5链接；申请失败时：错误信息结果H5链接
}

// ECApplyResponse 电子合同申请响应
type ECApplyResponse struct {
	Code     string               `json:"code" dc:"返回码"`         // 成功 000000 其它失败
	Msg      string               `json:"msg" dc:"返回码描述"`        // 返回码描述
	RespData *ECApplyResponseData `json:"resp_data" dc:"结果信息集合"` // 结果信息集合
	RespTime string               `json:"resp_time" dc:"响应时间"`   // 格式：yyyyMMddHHmmss
}

func (r *ECApplyResponse) SuccessOfFail() bool {
	return r.Code == "000000"
}
