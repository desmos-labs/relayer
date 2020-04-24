package relayer

import (
	"time"

	btsgtypes "github.com/bitsongofficial/go-bitsong/x/ibc/desmos"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreateSongPost creates a new transfer message
func (src *PathEnd) MsgCreateSongPost(
	dst *PathEnd, dstHeight uint64, songID string, creationTime time.Time, signer sdk.AccAddress,
) sdk.Msg {
	return btsgtypes.NewMsgCreateSongPost(
		src.PortID,
		src.ChannelID,
		dstHeight,

		songID,
		creationTime,
		signer,
	)
}

// PostCreatePacket creates a new post creation packet
func (src *PathEnd) PostCreatePacket(songID string, creationTime time.Time, sender string) []byte {
	return btsgtypes.NewSongCreationData(songID, creationTime, sender).GetBytes()
}
