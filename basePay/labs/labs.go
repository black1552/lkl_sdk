package labs

import (
	"context"

	relationclose "github.com/black1552/lkl_sdk/basePay/labs/relationClose"
	"github.com/black1552/lkl_sdk/basePay/labs/relationRefund"
	"github.com/black1552/lkl_sdk/basePay/labs/relationRevoked"
	transmicropay "github.com/black1552/lkl_sdk/basePay/labs/transMicropay"
	"github.com/black1552/lkl_sdk/basePay/labs/transMicropayEncry"
	"github.com/black1552/lkl_sdk/basePay/labs/transPreorder"
	"github.com/black1552/lkl_sdk/basePay/labs/transPreorederEncry"
	transquery "github.com/black1552/lkl_sdk/basePay/labs/transQuery"
	"github.com/black1552/lkl_sdk/common"
)

type Labs struct {
	RelationClose      *relationclose.RelationCloseServer
	RelationRefund     *relationRefund.RelationRefundServer
	RelationRevoked    *relationRevoked.RelationRevokedServer
	TransMicropay      *transmicropay.TransMicropayServer
	TransMicropayEncry *transMicropayEncry.TransMicropayEncryServer
	TransPreorder      *transPreorder.TransPreorderServer
	TransPreorderEncry *transPreorederEncry.TransPreorderEncryServer
	TransQuery         *transquery.TransQueryServer
}

func NewLabs(ctx context.Context, cfgJson string) *Labs {
	return &Labs{
		RelationClose:      relationclose.NewRelationCloseServer(common.NewClient[relationclose.RelationCloseResponse](ctx, cfgJson)),
		RelationRefund:     relationRefund.NewRelationRefundServer(common.NewClient[relationRefund.RelationRefundResponse](ctx, cfgJson)),
		RelationRevoked:    relationRevoked.NewRelationRevokedServer(common.NewClient[relationRevoked.RelationRevokedResponse](ctx, cfgJson)),
		TransMicropay:      transmicropay.NewTransMicropayServer(common.NewClient[transmicropay.TransMicropayResponse](ctx, cfgJson)),
		TransMicropayEncry: transMicropayEncry.NewTransMicropayEncryServer(common.NewClient[transMicropayEncry.TransMicropayEncryResponse](ctx, cfgJson)),
		TransPreorder:      transPreorder.NewTransPreorderServer(common.NewClient[transPreorder.TransPreorderResponse](ctx, cfgJson)),
		TransPreorderEncry: transPreorederEncry.NewTransPreorderEncryServer(common.NewClient[transPreorederEncry.TransPreorederEncryResponse](ctx, cfgJson)),
		TransQuery:         transquery.NewTransQueryServer(common.NewClient[transquery.TransQueryResponse](ctx, cfgJson)),
	}
}
