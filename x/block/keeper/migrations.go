package keeper

import (
	v2 "block/x/block/migrations/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) MigrateToV2(ctx sdk.Context) error {

	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
