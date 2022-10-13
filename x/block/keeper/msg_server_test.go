package keeper_test

import (
	"context"
	"testing"

	keepertest "block/testutil/keeper"
	"block/x/block/keeper"
	"block/x/block/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlockKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
