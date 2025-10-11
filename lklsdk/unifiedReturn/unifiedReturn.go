package unifiedreturn

import (
	"context"

	"github.com/black1552/lkl_sdk/lklsdk/common"
	"github.com/black1552/lkl_sdk/lklsdk/unifiedreturn/mergerefund"
	"github.com/black1552/lkl_sdk/lklsdk/unifiedreturn/refund"
	"github.com/black1552/lkl_sdk/lklsdk/unifiedreturn/refundquery"
)

type Server[T any] struct {
	Client       *common.Client[T]
	MergeRefound *mergerefund.MergeRefund[T]
	Refound      *refund.Refund[T]
	RefoundQuery *refundquery.RefundQuery[T]
}

// NewServer 创建拉卡拉统一退货服务实例
func NewServer[T any](ctx context.Context, cfgJson string) *Server[T] {
	client := common.NewClient[T](ctx, cfgJson)
	return &Server[T]{
		Client:       client,
		MergeRefound: mergerefund.NewMergeRefund[T](client),
		Refound:      refund.NewRefund[T](client),
		RefoundQuery: refundquery.NewRefundQuery[T](client),
	}
}

// MergeRefund 合单退货
func (u *Server[T]) MergeRefund(req *mergerefund.RequestDataMergeRefund) (*T, error) {
	return u.MergeRefound.MergeRefund(req)
}

// Refund 退货
func (u *Server[T]) Refund(req *refund.RequestDataRefund) (*T, error) {
	return u.Refound.Refund(req)
}

// RefundQuery 退货查询
func (u *Server[T]) RefundQuery(req *refundquery.RequestDataRefundQuery) (*T, error) {
	return u.RefoundQuery.RefundQuery(req)
}
