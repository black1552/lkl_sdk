package unifiedreturn

import (
	"context"

	"github.com/black1552/lkl_sdk/lklsdk/common"
	mergeRefund "github.com/black1552/lkl_sdk/lklsdk/unifiedReturn/mergeRefund"
	"github.com/black1552/lkl_sdk/lklsdk/unifiedReturn/refund"
	"github.com/black1552/lkl_sdk/lklsdk/unifiedReturn/refundquery"
)

type Server struct {
	MergeRefound *mergeRefund.MergeRefund
	Refound      *refund.Refund
	RefoundQuery *refundquery.RefundQuery
}

// NewServer 创建拉卡拉统一退货服务实例
func NewServer(ctx context.Context, cfgJson string) *Server {
	return &Server{
		MergeRefound: mergeRefund.NewMergeRefund(common.NewClient[mergeRefund.ResponseMergeRefund](ctx, cfgJson)),
		Refound:      refund.NewRefund(common.NewClient[refund.ResponseRefund](ctx, cfgJson)),
		RefoundQuery: refundquery.NewRefundQuery(common.NewClient[refundquery.ResponseRefundQuery](ctx, cfgJson)),
	}
}

// MergeRefund 合单退货
func (u *Server) MergeRefund(req *mergeRefund.RequestDataMergeRefund) (*mergeRefund.ResponseMergeRefund, error) {
	return u.MergeRefound.MergeRefund(req)
}

// Refund 退货
func (u *Server) Refund(req *refund.RequestDataRefund) (*refund.ResponseRefund, error) {
	return u.Refound.Refund(req)
}

// RefundQuery 退货查询
func (u *Server) RefundQuery(req *refundquery.RequestDataRefundQuery) (*refundquery.ResponseRefundQuery, error) {
	return u.RefoundQuery.RefundQuery(req)
}
