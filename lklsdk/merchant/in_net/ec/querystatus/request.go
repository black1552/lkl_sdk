package querystatus

// ECQueryStatusRequestData 电子合同查询状态请求数据
// 提供申请过与拉卡拉电子签约用户查询电子合同签署状态
// 请求URL：
// 测试环境：https://test.wsmsd.cn/sit/api/v3/mms/open_api/ec/q_status
// 生产环境：https://s2.lakala.com/api/v3/mms/open_api/ec/q_status
type ECQueryStatusRequestData struct {
	OrderNo   string `json:"order_no" dc:"四方机构自定义订单编号 必选，建议：平台编号+14位年月日时（24小时制）分秒+8位的随机数（同一接入机构不重复）"` // 必选，建议：平台编号+14位年月日时（24小时制）分秒+8位的随机数（同一接入机构不重复）
	OrgID     int    `json:"org_id" dc:"机构号 必选，签约方所属拉卡拉机构"`                                     // 必选，签约方所属拉卡拉机构
	EcApplyID int64  `json:"ec_apply_id" dc:"电子合同申请受理号 必选，申请接口反馈编号"`                               // 必选，申请接口反馈编号
}

// ECQueryStatusRequest 电子合同查询状态请求
type ECQueryStatusRequest struct {
	ReqTime string                    `json:"req_time" dc:"请求时间"` // 必选，格式：yyyyMMddHHmmss
	ReqData *ECQueryStatusRequestData `json:"req_data" dc:"请求数据"` // 请求数据
	Version string                    `json:"version" dc:"版本号"`   // 必选，固定值：1.0
}