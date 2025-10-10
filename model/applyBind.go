package model

import "github.com/black1552/lkl_sdk/consts"

// ApplyBindRequest 分账关系绑定请求结构体
// 用于发起分账接收方与商户的关系绑定申请
// 拉卡拉SDK接口文档：分账关系绑定接口

type ApplyBindRequest struct {
	ReqData *ApplyBindReqData `json:"reqData"` // 请求业务数据
	ReqId   string            `json:"reqId"`   // 请求ID，唯一标识一笔请求
	Version string            `json:"version"` // 接口版本号
	ReqTime string            `json:"reqTime"` // 请求时间，格式为yyyyMMddHHmmss
}

// ApplyBindReqData 分账关系绑定请求业务数据结构体
// 包含分账关系绑定所需的详细业务参数

type ApplyBindReqData struct {
	Version         string                 `json:"version"`               // 接口版本号，必传，长度8，取值说明：1.0
	OrderNo         string                 `json:"orderNo"`               // 订单编号，必传，长度32，用于后续跟踪排查问题及核对报文，格式为14位年月日(24小时制)分秒+8位随机数（不重复）
	OrgCode         string                 `json:"orgCode"`               // 分账接收方所属机构代码，必传，长度32
	MerInnerNo      string                 `json:"merInnerNo"`            // 分账商户内部商户号，必传，长度32，与MerCupNo选传其一，不能都为空
	MerCupNo        string                 `json:"merCupNo"`              // 分账商户银联商户号，必传，长度32，与MerInnerNo选传其一，不能都为空
	ReceiverNo      string                 `json:"receiverNo"`            // 分账接收方编号，必传，长度32
	EntrustFileName string                 `json:"entrustFileName"`       // 合作协议附件名称，必传，长度32
	EntrustFilePath string                 `json:"entrustFilePath"`       // 合作协议附件路径，必传，长度32，通过调用附件上传接口获取
	RetUrl          string                 `json:"retUrl"`                // 回调通知地址，必传，长度128，审核通过后通知地址
	Attachments     []*ApplyBindAttachment `json:"attachments,omitempty"` // 附加资料，可选，集合类型，其他附加资料文件信息
}

type ApplyBindAttachment struct {
	AttachType      consts.AttType `json:"attachType"`      // 附件类型编码，必传，长度32
	AttachName      string         `json:"attachName"`      // 附件名称，必传，长度32
	AttachStorePath string         `json:"attachStorePath"` // 附件路径，必传，长度128，通过调用附件上传接口获取
}

// ApplyBindResponse 分账关系绑定响应结构体
// 包含分账关系绑定申请的处理结果信息

type ApplyBindResponse struct {
	RetCode  string `json:"retCode"` // 响应码，000000表示成功
	RetMsg   string `json:"retMsg"`  // 响应消息，描述响应结果
	RespData struct {
		Version string `json:"version"` // 接口版本号，例如：547110502170558464
		OrderNo string `json:"orderNo"` // 订单编号，例如：2021020112000012345678
		OrgCode string `json:"orgCode"` // 机构代码，例如：200669
		ApplyId int64  `json:"applyId"` // 受理编号，例如：548099616395186176
	} `json:"respData"` // 响应业务数据，当retCode为000000时返回
}

func (a *ApplyBindResponse) SuccessOrFail() bool {
	return a.RetCode == "000000"
}
