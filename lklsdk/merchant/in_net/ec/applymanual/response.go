package applymanual

// ECApplyManualResponseData 电子合同人工复核申请响应数据结构体
type ECApplyManualResponseData struct {
	// 响应数据为空对象，根据API文档，成功时返回空的resp_data
}

// ECApplyManualResponse 电子合同人工复核申请响应结构体
type ECApplyManualResponse struct {
	Code     string                  `json:"code"`     // 返回码，成功为000000，其他为失败
	Msg      string                  `json:"msg"`      // 返回码描述
	RespData ECApplyManualResponseData `json:"resp_data"` // 响应数据
}

// SuccessOfFail 检查响应是否成功
// 返回true表示成功，false表示失败
func (r *ECApplyManualResponse) SuccessOfFail() bool {
	return r.Code == "000000"
}