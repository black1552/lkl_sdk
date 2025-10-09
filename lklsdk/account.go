package lklsdk

import (
	"time"

	"github.com/black1552/lkl_sdk/consts"
	"github.com/black1552/lkl_sdk/model"
)

// AccountService 账户服务
type AccountService[T any] struct {
	client *Client[T]
}

// NewAccountService 创建账户服务实例
func NewAccountService[T any](client *Client[T]) *AccountService[T] {
	return &AccountService[T]{
		client: client,
	}
}

// BalanceQuery 账户余额查询
func (a *AccountService[T]) BalanceQuery(req *model.BalanceQueryReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_ACCOUNT_BALANCE_QUERY_URL

	// 构建BaseModel请求
	baseReq := model.BalanceQueryRequest{
		ReqTime: time.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: req,
	}
	// 发送请求
	respBody, err := a.client.doRequest(url, baseReq)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// Withdraw 账户提现
func (a *AccountService[T]) Withdraw(req *model.WithdrawReqData) (*T, error) {
	// 构建请求参数
	url := consts.BASE_URL + consts.LKL_ACCOUNT_WITHDRAW_URL

	// 构建BaseModel请求
	baseReq := model.WithdrawRequest{
		ReqTime: time.Now().Format("20060102150405"),
		Version: "3.0",
		ReqData: req,
	}

	// 发送请求
	respBody, err := a.client.doRequest(url, baseReq)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
