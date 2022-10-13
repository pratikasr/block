package v2

import (
	v2types "block/x/block/migrations/v2/types"
	blocktypes "block/x/block/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuftypes "github.com/gogo/protobuf/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	err := migrateValuesAsset(store, cdc)
	if err != nil {
		return err
	}

	return nil
}
func migrateValuesAsset(store sdk.KVStore, cdc codec.BinaryCodec) error {

	key := blocktypes.AssetIDKey
	oldVal := store.Get(key)
	var id protobuftypes.UInt64Value
	cdc.MustUnmarshal(oldVal, &id)
	for i := uint64(1); i <= id.GetValue(); i++ {
		newVal := migrateValueAsset(cdc, oldVal)
		store.Delete(key)
		store.Set(key, newVal)
	}

	return nil
}

func migrateValueAsset(cdc codec.BinaryCodec, oldVal []byte) (newVal []byte) {

	// convert oldVal into lend type of previous version
	// use oldVal to create new lend of updated struct

	var asset v2types.Asset
	cdc.MustUnmarshal(oldVal, &asset)

	newAsset := blocktypes.Asset{
		Id:        asset.Id,
		Name:      asset.Name,
		Denom:     asset.Denom,
		Decimal:   asset.Decimal,
		IbcStatus: asset.IbcStatus,
	}

	newVal = cdc.MustMarshal(&newAsset)
	return newVal
}
