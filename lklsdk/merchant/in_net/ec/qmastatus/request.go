package qmastatus

// ECQmaStatusRequestData 电子合同人工复核结果查询请求数据结构体
type ECQmaStatusRequestData struct {
	Version   string `json:"version"`   // 版本号，默认1.0
	OrderNo   string `json:"order_no"`  // 四方机构自定义订单编号，必填，建议平台编号+14位年月日时分秒+8位随机数
	OrgID     int    `json:"org_id"`    // 机构号，必填，签约方所属拉卡拉机构
	EcApplyID string `json:"ec_apply_id"` // 电子合同申请受理号，必填，申请接口反馈编号
}

// ECQmaStatusRequest 电子合同人工复核结果查询请求结构体
type ECQmaStatusRequest struct {
	ReqTime string                  `json:"req_time"` // 请求时间
	ReqData *ECQmaStatusRequestData `json:"req_data"` // 请求数据
	Version string                  `json:"version"`  // 接口版本号
}