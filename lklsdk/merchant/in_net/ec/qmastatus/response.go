package qmastatus

// ECQmaStatusResponseData 电子合同人工复核结果查询响应数据结构体
type ECQmaStatusResponseData struct {
	Version        string `json:"version"`         // 版本号
	OrderNo        string `json:"order_no"`        // 四方机构自定义订单编号
	OrgID          int    `json:"org_id"`          // 机构号
	EcApplyID      string `json:"ec_apply_id"`     // 电子合同申请受理号
	AuditStatus    string `json:"audit_status"`    // 人工审核状态：WAIT_AUDIT待审核、PASS审核通过、REFUSE审核拒绝、CLOSE审核关闭
	AuditDesc      string `json:"audit_desc"`      // 人工审核结果说明
	SignH5URL      string `json:"sign_h5_url"`     // 签约H5地址，审核通过时返回
	SignH5URLExpTm string `json:"sign_h5_url_exp_tm"` // 签约H5地址过期时间，审核通过时返回
}

// ECQmaStatusResponse 电子合同人工复核结果查询响应结构体
type ECQmaStatusResponse struct {
	Code     string                  `json:"code"`      // 返回码，成功000000，其它失败
	Msg      string                  `json:"msg"`       // 返回码描述
	RespData *ECQmaStatusResponseData `json:"resp_data"` // 响应数据
}

// SuccessOfFail 判断响应是否成功
func (r *ECQmaStatusResponse) SuccessOfFail() bool {
	return r != nil && r.Code == "000000"
}