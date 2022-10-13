package keeper

import (
	"context"

	"block/x/block/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAsset(goCtx context.Context, msg *types.MsgAddAsset) (*types.MsgAddAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddAssetResponse{}, nil
}
