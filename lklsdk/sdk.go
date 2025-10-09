package lklsdk

import (
	"context"

	"github.com/black1552/lkl_sdk/model"
)

// SDK 拉卡拉SDK主入口
type SDK[T any] struct {
	Client      *Client[T]
	SplitLedger *SplitLedgerService[T]
	Trade       *TradeService[T]
	Account     *AccountService[T]
	UploadFile  *UploadFileService[T]
	MergePre    *MergePreService[T]
}

// NewSDK 创建拉卡拉SDK实例
func NewSDK[T any](ctx context.Context, cfgJson string) *SDK[T] {
	client := NewClient[T](ctx, cfgJson)
	return &SDK[T]{
		Client:      client,
		SplitLedger: NewSplitLedgerService(client),
		Trade:       NewTradeService(client),
		Account:     NewAccountService(client),
		UploadFile:  NewUploadFileService(client),
		MergePre:    NewMergePreService(client),
	}
}

// 以下为便捷方法，直接通过SDK调用各服务的主要功能

// MergePreOrder 主扫合单交易
func (s *SDK[T]) MergePreOrder(req *model.MergePreorderReqData) (*T, error) {
	return s.MergePre.PreOrder(req)
}

// Refound 退款
func (s *SDK[T]) Refound(req *model.RefundReqData) (*T, error) {
	return s.Trade.Refound(req)
}

// ApplyLedgerMer 商户分账业务开通申请
func (s *SDK[T]) ApplyLedgerMer(req *model.ApplyLedgerMerReqData) (*T, error) {
	return s.SplitLedger.ApplyLedgerMer(req)
}

// QueryLedgerMer 商户分账信息查询
func (s *SDK[T]) QueryLedgerMer(req *model.QueryLedgerMerReqData) (*T, error) {
	return s.SplitLedger.QueryLedgerMer(req)
}

// ApplyLedgerReceiver 分账接收方创建申请
func (s *SDK[T]) ApplyLedgerReceiver(req *model.ApplyLedgerReceiverReqData) (*T, error) {
	return s.SplitLedger.ApplyLedgerReceiver(req)
}

// ApplyBind 分账关系绑定申请
func (s *SDK[T]) ApplyBind(req *model.ApplyBindReqData) (*T, error) {
	return s.SplitLedger.ApplyBind(req)
}

// QuerySplitBalance 可分账金额查询
func (s *SDK[T]) QuerySplitBalance(req *model.SplitBalanceReqData) (*T, error) {
	return s.SplitLedger.QuerySplitBalance(req)
}

// OrderSplitLedger 订单分账
func (s *SDK[T]) OrderSplitLedger(req *model.OrderSplitLedgerReqData) (*T, error) {
	return s.SplitLedger.OrderSplitLedger(req)
}

// TradeQuery 交易查询
func (s *SDK[T]) TradeQuery(req *model.TradeQueryReqData) (*T, error) {
	return s.Trade.TradeQuery(req)
}

// PreOrder 聚合预下单
func (s *SDK[T]) PreOrder(req *model.PreorderReqData) (*T, error) {
	return s.Trade.PreOrder(req)
}

// BalanceQuery 账户余额查询
func (s *SDK[T]) BalanceQuery(req *model.BalanceQueryReqData) (*T, error) {
	return s.Account.BalanceQuery(req)
}

func (s *SDK[T]) UploadFileQuery(req *model.UploadFileReqData) (*T, error) {
	return s.UploadFile.UploadFileQuery(req)
}

// Withdraw 账户提现
func (s *SDK[T]) Withdraw(req *model.WithdrawReqData) (*T, error) {
	return s.Account.Withdraw(req)
}
