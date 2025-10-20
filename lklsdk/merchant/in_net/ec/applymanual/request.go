package applymanual

import (
	"github.com/black1552/lkl_sdk/consts"
)

// ECApplyManualRequestData 电子合同人工复核申请请求数据结构体
type ECApplyManualRequestData struct {
	Version   string        `json:"version"`           // 版本号，默认1.0
	OrderNo   string        `json:"order_no"`          // 四方机构自定义订单编号，必填，建议平台编号+14位年月日时分秒+8位随机数
	OrgID     int           `json:"org_id"`            // 机构号，必填，签约方所属拉卡拉机构
	EcApplyID string        `json:"ec_apply_id"`       // 电子合同申请受理号，必填，申请接口反馈编号
	ApplyDesc string        `json:"apply_desc"`        // 申请说明，必填，复议提交的原因说明
	FileData  []*ECFileData `json:"file_data"`         // 附件信息集合，必填，提供给审核人员核对信息
	RetURL    string        `json:"ret_url,omitempty"` // 回调地址，可选
}

// ECFileData 附件信息结构体
type ECFileData struct {
	AttachType      consts.AttType `json:"attach_type"`       // 附件类型，必填，枚举值见文档
	AttachName      string `json:"attach_name"`       // 附件名称，必填，最大32字符
	AttachExtName   string `json:"attach_ext_name"`   // 附件格式，必填，jpg、pdf等
	AttachStorePath string `json:"attach_store_path"` // 附件坐标URL，必填，如G1/M00/00/61/CrFdEl3IyceAVVd8AAA0ADuZsA0911.jpg
}

// ECApplyManualRequest 电子合同人工复核申请请求结构体
type ECApplyManualRequest struct {
	ReqTime string                    `json:"req_time"` // 请求时间
	ReqData *ECApplyManualRequestData `json:"req_data"` // 请求数据
	Version string                    `json:"version"`  // 接口版本号
}
