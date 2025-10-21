package relationRefund

import (
	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelationRefundServer 聚合扫码退款API结构体
type RelationRefundServer struct {
	Client *common.Client[RelationRefundResponse]
}

// NewRelationRefundServer 创建聚合扫码退款API实例
func NewRelationRefundServer(client *common.Client[RelationRefundResponse]) *RelationRefundServer {
	return &RelationRefundServer{
		Client: client,
	}
}

// RelationRefund 聚合扫码退款方法
// - reqData: 退款请求数据
func (a *RelationRefundServer) RelationRefund(reqData *RelationRefundRequestData) (*RelationRefundResponse, error) {
	// 构建请求结构体
	request := &RelationRefundRequest{
		ReqTime: gtime.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: reqData,
	}

	// 调用Client发送请求
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_RELATION_REFUND_URL, request)
}
