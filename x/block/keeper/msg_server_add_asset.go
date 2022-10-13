package keeper

import (
	"block/x/block/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m msgServer) AddAsset(goCtx context.Context, msg *types.MsgAddAsset) (*types.MsgAddAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := m.Keeper.AddAsset(ctx, *msg); err != nil {
		return nil, err
	}
	return &types.MsgAddAssetResponse{}, nil
}
