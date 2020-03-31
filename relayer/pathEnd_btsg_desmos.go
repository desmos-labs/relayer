package relayer

import (
	"time"

	btsgtypes "github.com/bitsongofficial/go-bitsong/x/ibc_desmos"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chanState "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	dsmtypes "github.com/desmos-labs/desmos/x/posts"
)

// MsgCreateSongPost creates a new transfer message
func (src *PathEnd) MsgCreateSongPost(dst *PathEnd, dstHeight uint64, signer sdk.AccAddress) sdk.Msg {
	return btsgtypes.NewMsgCreateSongPost(
		src.PortID,
		src.ChannelID,
		dstHeight,
		signer,
	)
}

// PostCreatePacket creates a new post creation packet
func (src *PathEnd) PostCreatePacket(sender sdk.AccAddress, timeout uint64) chanState.PacketDataI {

	// TODO: Specify the post data
	// The message, subspace and song_id should be probably be specified inside the MsgCreateSongPost and later
	// retrieved here to be used properly.

	return dsmtypes.NewCreatePostPacketData(
		dsmtypes.NewPostCreationData(
			"New song post",
			dsmtypes.PostID(0),
			true,
			"a31be8a1946fb15200d7081163bf3c41eae3b8b745e8bbf7d96e04e57c9ddf9b",
			map[string]string{
				"song_id": "songId",
			},
			sender,
			time.Now(),
			nil,
			nil,
		),
		timeout,
	)
}
