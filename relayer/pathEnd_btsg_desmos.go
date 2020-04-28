package relayer

import (
	"fmt"
	"time"

	btsgtypes "github.com/bitsongofficial/go-bitsong/x/ibc/desmos"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcposts "github.com/desmos-labs/desmos/x/ibc/posts"
)

func createSongPacketData(songID string, creationTime time.Time, sender string) ibcposts.PostCreationPacketData {
	return btsgtypes.NewSongCreationData(songID, creationTime, sender)
}

// MsgCreateSongPost creates a new transfer message
func (src *PathEnd) MsgCreateSongPost(
	dst *PathEnd, dstHeight uint64, songID string, creationTime time.Time, postCreator string, signer sdk.AccAddress,
) sdk.Msg {
	data := createSongPacketData(songID, creationTime, postCreator)
	return ibcposts.NewMsgCrossPost(
		src.PortID,
		src.ChannelID,
		dstHeight,
		data,
		signer,
	)
}

// PostCreatePacket creates a new post creation packet
func (src *PathEnd) PostCreatePacket(songID string, creationTime time.Time, sender string) []byte {
	data := createSongPacketData(songID, creationTime, sender).GetBytes()
	fmt.Printf("created packet: %s\n", data)
	return data
}
