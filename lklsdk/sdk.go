package lklsdk

import (
	"context"

	"github.com/black1552/lkl_sdk/lklsdk/common"
	ecApply "github.com/black1552/lkl_sdk/lklsdk/merchant/in_net/ec/apply"
	ecApplyManual "github.com/black1552/lkl_sdk/lklsdk/merchant/in_net/ec/applymanual"
	ecDownload "github.com/black1552/lkl_sdk/lklsdk/merchant/in_net/ec/download"
	ecQmaStatus "github.com/black1552/lkl_sdk/lklsdk/merchant/in_net/ec/qmastatus"
	ecQuery "github.com/black1552/lkl_sdk/lklsdk/merchant/in_net/ec/querystatus"
	"github.com/black1552/lkl_sdk/model"
)

// SDK 拉卡拉SDK主入口
type SDK[T any] struct {
	Client          *common.Client[T]
	SplitLedger     *SplitLedgerService[T]
	Trade           *TradeService[T]
	Account         *AccountService[T]
	UploadFile      *UploadFileService[T]
	MergePre        *MergePreService[T]
	Merchant        *MerService[T]
	EC              *ecApply.Apply
	ECQuery         *ecQuery.QStatus
	ECFileDownload  *ecDownload.Download
	ECPeApplyManual *ecApplyManual.ApplyManual
	ECPeQmaStatus   *ecQmaStatus.QmaStatus
}

// NewSDK 创建拉卡拉SDK实例
func NewSDK[T any](ctx context.Context, cfgJson string) *SDK[T] {
	client := common.NewClient[T](ctx, cfgJson)
	return &SDK[T]{
		Client:          client,
		SplitLedger:     NewSplitLedgerService(client),
		Trade:           NewTradeService(client),
		Account:         NewAccountService(client),
		UploadFile:      NewUploadFileService(client),
		MergePre:        NewMergePreService(client),
		Merchant:        NewMerService(client),
		EC:              ecApply.NewEcApply(common.NewClient[ecApply.ECApplyResponse](ctx, cfgJson)),
		ECQuery:         ecQuery.NewQStatus(common.NewClient[ecQuery.ECQueryStatusResponse](ctx, cfgJson)),
		ECFileDownload:  ecDownload.NewDownload(common.NewClient[ecDownload.ECDownloadResponse](ctx, cfgJson)),
		ECPeApplyManual: ecApplyManual.NewApplyManual(common.NewClient[ecApplyManual.ECApplyManualResponse](ctx, cfgJson)),
		ECPeQmaStatus:   ecQmaStatus.NewQmaStatus(common.NewClient[ecQmaStatus.ECQmaStatusResponse](ctx, cfgJson)),
	}
}

// 以下为便捷方法，直接通过SDK调用各服务的主要功能

// ReconsiderSubmit 商户进件复议提交
func (s *SDK[T]) ReconsiderSubmit(req *model.ReConfSubmitRequestData) (*T, error) {
	return s.Merchant.ReconsiderSubmit(req)
}

// ReconsiderSubmitTest 商户进件复议提交
func (s *SDK[T]) ReconsiderSubmitTest(req *model.ReConfSubmitRequestData) (*T, error) {
	return s.Merchant.ReconsiderSubmitTest(req)
}

// QueryMerchant 商户进件信息查询
func (s *SDK[T]) QueryMerchant(req *model.QueryMerRequestData) (*T, error) {
	return s.Merchant.QueryMer(req)
}

// QueryMerchantTest 商户进件信息查询
func (s *SDK[T]) QueryMerchantTest(req *model.QueryMerRequestData) (*T, error) {
	return s.Merchant.QueryMerTest(req)
}

// MerValidate 商户进件信息校验
func (s *SDK[T]) MerValidate(req *model.MerValidateRequestData) (*T, error) {
	return s.Merchant.MerValidate(req)
}

// MerValidateTest 商户进件信息校验
func (s *SDK[T]) MerValidateTest(req *model.MerValidateRequestData) (*T, error) {
	return s.Merchant.MerValidateTest(req)
}

// AddMer 商户进件
func (s *SDK[T]) AddMer(req *model.MerchantApplyReqData) (*T, error) {
	return s.Merchant.AddMer(req)
}

// AddMerTest 商户进件
func (s *SDK[T]) AddMerTest(req *model.MerchantApplyReqData) (*T, error) {
	return s.Merchant.AddMerTest(req)
}

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

func (s *SDK[T]) ApplyLedgerMerTest(req *model.ApplyLedgerMerReqData) (*T, error) {
	return s.SplitLedger.ApplyLedgerMerTest(req)
}

// QueryLedgerMer 商户分账信息查询
func (s *SDK[T]) QueryLedgerMer(req *model.QueryLedgerMerReqData) (*T, error) {
	return s.SplitLedger.QueryLedgerMer(req)
}

func (s *SDK[T]) QueryLedgerMerTest(req *model.QueryLedgerMerReqData) (*T, error) {
	return s.SplitLedger.QueryLedgerMerTest(req)
}

// ApplyLedgerReceiver 分账接收方创建申请
func (s *SDK[T]) ApplyLedgerReceiver(req *model.ApplyLedgerReceiverReqData) (*T, error) {
	return s.SplitLedger.ApplyLedgerReceiver(req)
}

func (s *SDK[T]) ApplyLedgerReceiverTest(req *model.ApplyLedgerReceiverReqData) (*T, error) {
	return s.SplitLedger.ApplyLedgerReceiverTest(req)
}

// ApplyBind 分账关系绑定申请
func (s *SDK[T]) ApplyBind(req *model.ApplyBindReqData) (*T, error) {
	return s.SplitLedger.ApplyBind(req)
}

func (s *SDK[T]) ApplyBindTest(req *model.ApplyBindReqData) (*T, error) {
	return s.SplitLedger.ApplyBindTest(req)
}

// QuerySplitBalance 可分账金额查询
func (s *SDK[T]) QuerySplitBalance(req *model.SplitBalanceReqData) (*T, error) {
	return s.SplitLedger.QuerySplitBalance(req)
}

func (s *SDK[T]) QuerySplitBalanceTest(req *model.SplitBalanceReqData) (*T, error) {
	return s.SplitLedger.QuerySplitBalanceTest(req)
}

// OrderSplitLedger 订单分账
func (s *SDK[T]) OrderSplitLedger(req *model.OrderSplitLedgerReqData) (*T, error) {
	return s.SplitLedger.OrderSplitLedger(req)
}

func (s *SDK[T]) OrderSplitLedgerTest(req *model.OrderSplitLedgerReqData) (*T, error) {
	return s.SplitLedger.OrderSplitLedgerTest(req)
}

// OrderSplitLedgerResultQuery 订单分账结果查询
func (s *SDK[T]) OrderSplitLedgerResultQuery(req *model.OrderSplitLedgerResultQueryReqData) (*T, error) {
	return s.SplitLedger.OrderSplitLedgerResultQuery(req)
}

// OrderSplitLedgerFallback 订单分账回退
func (s *SDK[T]) OrderSplitLedgerFallback(req *model.OrderSplitLedgerFallbackReqData) (*T, error) {
	return s.SplitLedger.OrderSplitLedgerFallback(req)
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

func (s *SDK[T]) BalanceQueryTest(req *model.BalanceQueryReqData) (*T, error) {
	return s.Account.BalanceQueryTest(req)
}

func (s *SDK[T]) UploadFileQuery(req *model.UploadFileReqData) (*T, error) {
	return s.UploadFile.UploadFileQuery(req)
}
func (s *SDK[T]) UploadFileQueryTest(req *model.UploadFileReqData) (*T, error) {
	return s.UploadFile.UploadFileQueryTest(req)
}

// Withdraw 账户提现
func (s *SDK[T]) Withdraw(req *model.WithdrawReqData) (*T, error) {
	return s.Account.Withdraw(req)
}

func (s *SDK[T]) WithdrawTest(req *model.WithdrawReqData) (*T, error) {
	return s.Account.WithdrawTest(req)
}

// ECApply 电子合同申请
func (s *SDK[T]) ECApply(req *ecApply.ECApplyRequestData) (*ecApply.ECApplyResponse, error) {
	return s.EC.ECApply(req)
}

// ECApplyTest 电子合同申请（测试环境）
func (s *SDK[T]) ECApplyTest(req *ecApply.ECApplyRequestData) (*ecApply.ECApplyResponse, error) {
	return s.EC.ECApplyTest(req)
}

// ECQueryStatus 电子合同查询状态
func (s *SDK[T]) ECQueryStatus(req *ecQuery.ECQueryStatusRequestData) (*ecQuery.ECQueryStatusResponse, error) {
	return s.ECQuery.QueryStatus(req)
}

// ECQueryStatusTest 电子合同查询状态（测试环境）
func (s *SDK[T]) ECQueryStatusTest(req *ecQuery.ECQueryStatusRequestData) (*ecQuery.ECQueryStatusResponse, error) {
	return s.ECQuery.QueryStatusTest(req)
}

// ECDownload 电子合同下载
func (s *SDK[T]) ECDownload(req *ecDownload.ECDownloadRequestData) (*ecDownload.ECDownloadResponse, error) {
	return s.ECFileDownload.ECDownload(req)
}

// ECDownloadTest 电子合同下载（测试环境）
func (s *SDK[T]) ECDownloadTest(req *ecDownload.ECDownloadRequestData) (*ecDownload.ECDownloadResponse, error) {
	return s.ECFileDownload.ECDownloadTest(req)
}

// ECApplyManual 电子合同人工复核申请
func (s *SDK[T]) ECApplyManual(req *ecApplyManual.ECApplyManualRequestData) (*ecApplyManual.ECApplyManualResponse, error) {
	return s.ECPeApplyManual.ECApplyManual(req)
}

// ECApplyManualTest 电子合同人工复核申请（测试环境）
func (s *SDK[T]) ECApplyManualTest(req *ecApplyManual.ECApplyManualRequestData) (*ecApplyManual.ECApplyManualResponse, error) {
	return s.ECPeApplyManual.ECApplyManualTest(req)
}

// ECQmaStatus 电子合同人工复核结果查询
func (s *SDK[T]) ECQmaStatus(req *ecQmaStatus.ECQmaStatusRequestData) (*ecQmaStatus.ECQmaStatusResponse, error) {
	return s.ECPeQmaStatus.ECQmaStatus(req)
}

// ECQmaStatusTest 电子合同人工复核结果查询（测试环境）
func (s *SDK[T]) ECQmaStatusTest(req *ecQmaStatus.ECQmaStatusRequestData) (*ecQmaStatus.ECQmaStatusResponse, error) {
	return s.ECPeQmaStatus.ECQmaStatusTest(req)
}

func (s *SDK[T]) WechatRealNameQuery(req *model.WechatRealNameQueryReqData) (*T, error) {
	return s.Merchant.WechatRealNameQuery(req)
}
