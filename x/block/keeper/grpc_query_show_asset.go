package keeper

import (
	"context"

	"block/x/block/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q Keeper) ShowAsset(c context.Context, req *types.QueryShowAssetRequest) (*types.QueryShowAssetResponse, error) {

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be empty")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetAsset(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "asset does not exist for id %d", req.Id)
	}

	return &types.QueryShowAssetResponse{
		Asset: item,
	}, nil

}
