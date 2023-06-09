package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sge-network/sge/x/house/types"
)

// Keeper of the house store
type Keeper struct {
	storeKey    sdk.StoreKey
	cdc         codec.BinaryCodec
	paramstore  paramtypes.Subspace
	authzKeeper types.AuthzKeeper
	srKeeper    types.SRKeeper
	ovmKeeper   types.OVMKeeper
}

// SdkExpectedKeepers contains expected keepers parameter needed by NewKeeper
type SdkExpectedKeepers struct {
	AuthzKeeper types.AuthzKeeper
}

// NewKeeper returns an instance of the housekeeper
func NewKeeper(
	cdc codec.BinaryCodec,
	key sdk.StoreKey,
	srKeeper types.SRKeeper,
	ovmKeeper types.OVMKeeper,
	ps paramtypes.Subspace,
	expectedKeepers SdkExpectedKeepers,
) *Keeper {
	// set KeyTable if it is not already set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		storeKey:    key,
		cdc:         cdc,
		srKeeper:    srKeeper,
		ovmKeeper:   ovmKeeper,
		paramstore:  ps,
		authzKeeper: expectedKeepers.AuthzKeeper,
	}
}
