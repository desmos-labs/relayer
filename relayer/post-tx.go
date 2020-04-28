package relayer

import (
	"fmt"
	"time"

	"github.com/avast/retry-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chanTypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	tmclient "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
)

// SendTransferBothSides sends a packet asking to create a post from src to dst.
// The post owner will be the src address.
func (src *Chain) SendPostBothSides(dst *Chain, songID string) error {
	dstHeader, err := dst.UpdateLiteWithHeader()
	if err != nil {
		return err
	}

	creationTime := time.Now()

	timeoutHeight := dstHeader.GetHeight() + uint64(defaultPacketTimeout)

	// Properly render the address string
	done := dst.UseSDKContext()
	dstAddrString := src.MustGetAddress().String()
	done()

	// MsgCreateSongPost will call SendPacket on src chain
	txs := RelayMsgs{
		Src: []sdk.Msg{
			src.PathEnd.MsgCreateSongPost(
				dst.PathEnd,
				dstHeader.GetHeight(),
				songID,
				creationTime,
				dstAddrString,
				src.MustGetAddress(),
			),
		},
		Dst: []sdk.Msg{},
	}

	if txs.Send(src, dst); !txs.Success() {
		return fmt.Errorf("failed to send first transaction")
	}

	// Working on SRC chain :point_up:
	// Working on DST chain :point_down:

	var (
		hs           map[string]*tmclient.Header
		seqRecv      chanTypes.RecvResponse
		seqSend      uint64
		srcCommitRes CommitmentResponse
	)

	if err = retry.Do(func() error {
		hs, err = UpdatesWithHeaders(src, dst)
		if err != nil {
			return err
		}

		seqRecv, err = dst.QueryNextSeqRecv(hs[dst.ChainID].Height)
		if err != nil {
			return err
		}

		seqSend, err = src.QueryNextSeqSend(hs[src.ChainID].Height)
		if err != nil {
			return err
		}

		srcCommitRes, err = src.QueryPacketCommitment(hs[src.ChainID].Height-1, int64(seqSend-1))
		if err != nil {
			return err
		}

		if srcCommitRes.Proof.Proof == nil {
			return fmt.Errorf("nil proof, retrying")
		}
		return nil
	}); err != nil {
		return err
	}

	// Properly render the source and destination address strings
	done = dst.UseSDKContext()
	srcAddr := src.MustGetAddress().String()
	done()

	// reconstructing packet data here instead of retrieving from an indexed node
	packet := src.PathEnd.PostCreatePacket(
		songID,
		creationTime,
		srcAddr,
	)

	// Debugging by simply passing in the packet information that we know was sent earlier in the SendPacket
	// part of the command. In a real relayer, this would be a separate command that retrieved the packet
	// information from an indexing node
	txs = RelayMsgs{
		Dst: []sdk.Msg{
			dst.PathEnd.UpdateClient(hs[src.ChainID], dst.MustGetAddress()),
			dst.PathEnd.MsgRecvPacket(
				src.PathEnd,
				seqRecv.NextSequenceRecv,
				timeoutHeight,
				defaultPacketTimeoutStamp(),
				packet,
				srcCommitRes.Proof,
				srcCommitRes.ProofHeight,
				dst.MustGetAddress(),
			),
		},
		Src: []sdk.Msg{},
	}

	txs.Send(src, dst)
	return nil
}
