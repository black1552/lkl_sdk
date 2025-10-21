package unifiedreturn

import (
	"context"

	"github.com/black1552/lkl_sdk/common"
	mergeRefund "github.com/black1552/lkl_sdk/unifiedReturn/mergeRefund"
	"github.com/black1552/lkl_sdk/unifiedReturn/refund"
	"github.com/black1552/lkl_sdk/unifiedReturn/refundquery"
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
