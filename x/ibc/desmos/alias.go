package desmos

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/bitsongofficial/go-bitsong/x/ibc/desmos/keeper"
	"github.com/bitsongofficial/go-bitsong/x/ibc/desmos/types"
)

const (
	DefaultPacketTimeout  = keeper.DefaultPacketTimeout
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
	RouterKey             = types.RouterKey
	QuerierRoute          = types.QuerierRoute
	DesmosBitsongSubspace = types.DesmosBitsongSubspace
	DesmosSongIDAttribute = types.DesmosSongIDAttribute
)

var (
	// functions aliases
	NewKeeper            = keeper.NewKeeper
	NewMsgCreateSongPost = types.NewMsgCreateSongPost
	NewSongCreationData  = types.NewSongCreationData
	RegisterCodec        = types.RegisterCodec

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper            = keeper.Keeper
	MsgCreateSongPost = types.MsgCreateSongPost
)
