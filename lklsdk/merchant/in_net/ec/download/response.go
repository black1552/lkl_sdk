package download

// ECDownloadResponseData 电子合同下载响应数据
type ECDownloadResponseData struct {
	OrderNo   string `json:"order_no"`   // 请求上送的订单号
	EcApplyID int64  `json:"ec_apply_id"` // 电子签约申请受理编号
	EcStatus  string `json:"ec_status"`  // 电子合同状态: UNDONE 未完成, COMPLETED 已完成
	EcNo      string `json:"ec_no"`      // 电子合同号（合同状态为COMPLETED时返回）
	EcFile    string `json:"ec_file"`    // 电子合同pdf文件(base64格式字符，合同状态为COMPLETED时返回)
}

// ECDownloadResponse 电子合同下载响应结构
type ECDownloadResponse struct {
	Code     string               `json:"code"`      // 返回码
	Msg      string               `json:"msg"`       // 返回码描述
	RespData *ECDownloadResponseData `json:"resp_data"` // 结果信息集合
}

// SuccessOfFail 检查响应是否成功
func (r *ECDownloadResponse) SuccessOfFail() bool {
	return r.Code == "000000"
}