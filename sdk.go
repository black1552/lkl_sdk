package lklSDK

import (
	"context"

	"github.com/black1552/lkl_sdk/basePay/labs"
	"github.com/black1552/lkl_sdk/unifiedreturn"
)

type SDK struct {
	labs          *labs.Labs
	unifiedReturn *unifiedreturn.Server
}

func NewSDK(ctx context.Context, cfgJson string) *SDK {
	return &SDK{
		labs:          labs.NewLabs(ctx, cfgJson),
		unifiedReturn: unifiedreturn.NewServer(ctx, cfgJson),
	}
}
