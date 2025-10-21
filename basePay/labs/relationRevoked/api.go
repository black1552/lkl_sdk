package relationRevoked

import (
	"github.com/black1552/lkl_sdk/common"
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelationRevokedServer 聚合扫码撤销API结构体
type RelationRevokedServer struct {
	Client *common.Client[RelationRevokedResponse]
}

// NewRelationRevokedServer 创建聚合扫码撤销API实例
func NewRelationRevokedServer(client *common.Client[RelationRevokedResponse]) *RelationRevokedServer {
	return &RelationRevokedServer{
		Client: client,
	}
}

// RelationRevoked 聚合扫码撤销方法
// - reqData: 撤销请求数据
func (a *RelationRevokedServer) RelationRevoked(reqData *RelationRevokedRequestData) (*RelationRevokedResponse, error) {
	// 构建请求结构体
	request := &RelationRevokedRequest{
		ReqTime: gtime.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: reqData,
	}

	// 调用Client发送请求
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_RELATION_REVOKED_URL, request)
}
