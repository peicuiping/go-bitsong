package keeper

import (
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/track/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Handle all "content" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgTrackAdd:
			return handleMsgTrackAdd(ctx, keeper, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized content message type: %T", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgTrackAdd(ctx sdk.Context, keeper Keeper, msg types.MsgTrackAdd) (*sdk.Result, error) {

	artists := make([]types.Artist, len(msg.Artists))
	for i, name := range msg.Artists {
		artist, err := keeper.GetOrSetArtist(ctx, *types.NewArtist(name, nil, nil, nil))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		artists[i] = *artist
	}

	track, err := types.NewTrack(
		msg.Title, artists, msg.Number, msg.Duration, msg.Explicit, msg.ExternalIds, msg.ExternalUrls, msg.PreviewUrl, msg.Dao,
	)

	if err != nil {
		return nil, err
	}

	cid, err := keeper.Add(ctx, track)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTrackAdded,
			sdk.NewAttribute(types.AttributeKeyTrackCid, cid),
		),
	)

	return &sdk.Result{
		//Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(cid),
		Events: ctx.EventManager().Events().ToABCIEvents(),
	}, nil
}
