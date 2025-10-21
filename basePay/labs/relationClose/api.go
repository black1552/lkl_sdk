package relationclose

import (
	"github.com/black1552/lkl_sdk/common"
	"github.com/black1552/lkl_sdk/consts"
	"github.com/gogf/gf/v2/os/gtime"
)

// RelationCloseServer 聚合扫码关单API结构体
type RelationCloseServer struct {
	Client *common.Client[RelationCloseResponse]
}

// NewAPI 创建聚合扫码关单API实例
func NewRelationCloseServer(client *common.Client[RelationCloseResponse]) *RelationCloseServer {
	return &RelationCloseServer{
		Client: client,
	}
}

// RelationClose 执行聚合扫码关单请求
// - merchantNo: 商户号
// - termNo: 终端号
// - 原交易标识：originOutTradeNo、originTradeNo、(originOutOrderSource+originOutOrderNo) 三者必选其一
// - locationInfo: 地址位置信息（风控要求必送）
func (a *RelationCloseServer) RelationClose(reqData *RelationCloseRequestData) (*RelationCloseResponse, error) {
	// 构建请求结构体
	request := &RelationCloseRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: reqData,
	}
	// 发送请求 - 使用consts中定义的聚合扫码关单URL
	return a.Client.DoRequest(consts.BASE_URL+consts.LKL_BASE_URL_RELATION_CLOSE_URL, request)
}

// RelationCloseTest 执行聚合扫码关单测试请求
// - merchantNo: 商户号
// - termNo: 终端号
// - 原交易标识：originOutTradeNo、originTradeNo、(originOutOrderSource+originOutOrderNo) 三者必选其一
// - locationInfo: 地址位置信息（风控要求必送）
func (a *RelationCloseServer) RelationCloseTest(reqData *RelationCloseRequestData) (*RelationCloseResponse, error) {
	// 构建请求结构体
	request := &RelationCloseRequest{
		ReqTime: gtime.Now().Format("YmdHis"),
		Version: "3.0",
		ReqData: reqData,
	}
	// 发送请求 - 使用consts中定义的聚合扫码关单URL
	return a.Client.DoRequest(consts.BASE_TEST_URL+consts.LKL_BASE_URL_RELATION_CLOSE_URL, request)
}
