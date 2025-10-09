package lklsdk

import (
	"time"

	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/model"
)

// TradeService 交易服务
type TradeService[T any] struct {
	client *Client[T]
}

// NewTradeService 创建交易服务实例
func NewTradeService[T any](client *Client[T]) *TradeService[T] {
	return &TradeService[T]{
		client: client,
	}
}

// TradeQuery 交易查询
func (t *TradeService[T]) TradeQuery(req *model.TradeQueryReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_TRADE_QUERY_URL

	// 构建BaseModel请求
	baseReq := model.TradeQuery{
		ReqTime:    time.Now().Format("20060102150405"),
		Version:    "3.0",
		OutOrgCode: t.client.config.AppId,
		ReqData:    req,
	}

	// 发送请求
	respBody, err := t.client.doRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

// PreOrder 聚合预下单
func (t *TradeService[T]) PreOrder(req *model.PreorderReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_PRE_ORDER_URL

	// 构建BaseModel请求
	baseReq := model.NewPreorder(req)

	// 发送请求
	respBody, err := t.client.doRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// Refound 退款
func (t *TradeService[T]) Refound(req *model.RefundReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_REFOUND_URL

	// 构建BaseModel请求
	baseReq := model.NewRefund(req)
	// 发送请求
	respBody, err := t.client.doRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
