package reward

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/bitsongofficial/go-bitsong/x/reward/keeper"
	"github.com/bitsongofficial/go-bitsong/x/reward/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	RewardPoolKey = types.RewardPoolKey
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	RewardPool   = types.RewardPool
)