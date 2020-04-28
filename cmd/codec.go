package cmd

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ibcposts "github.com/desmos-labs/desmos/x/ibc/posts"
)

func RegisterCodec(codec *codec.Codec) {
	ibcposts.RegisterCodec(codec)
}
