package download

// ECDownloadRequestData 电子合同下载请求数据
type ECDownloadRequestData struct {
	OrderNo   string `json:"order_no" dc:"四方机构自定义订单编号"`   // 四方机构自定义订单编号
	OrgCode   int    `json:"org_code" dc:"机构号"`   // 机构号
	EcApplyID int64  `json:"ec_apply_id" dc:"电子合同申请受理号"` // 电子合同申请受理号
}

// ECDownloadRequest 电子合同下载请求结构
type ECDownloadRequest struct {
	ReqTime string               `json:"req_time"` // 请求时间
	ReqData *ECDownloadRequestData `json:"req_data"` // 请求数据
	Version string               `json:"version"`  // 接口版本号
}