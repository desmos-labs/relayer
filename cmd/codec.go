package cmd

import (
	btsgtypes "github.com/bitsongofficial/go-bitsong/x/ibc/desmos"
	"github.com/cosmos/cosmos-sdk/codec"
	dsmstypes "github.com/desmos-labs/desmos/x/posts"
)

func RegisterCodec(codec *codec.Codec) {
	btsgtypes.RegisterCodec(codec)
	dsmstypes.RegisterCodec(codec)
}
