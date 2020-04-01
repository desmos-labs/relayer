package cmd

import (
	"fmt"
	"time"

	"github.com/avast/retry-go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chanTypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	tmclient "github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
	"github.com/iqlusioninc/relayer/relayer"
	"github.com/spf13/cobra"
)

func postCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "song-post [src-chain-id] [dst-chain-id] [song-id]",
		Short: "spost",
		Long:  "This creates a new post to a Desmos-based chain",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			src, dst := args[0], args[1]
			c, err := config.Chains.Gets(src, dst)
			if err != nil {
				return err
			}

			pth, err := cmd.Flags().GetString(flagPath)
			if err != nil {
				return err
			}

			if _, err = setPathsFromArgs(c[src], c[dst], pth); err != nil {
				return err
			}

			dstHeader, err := c[dst].UpdateLiteWithHeader()
			if err != nil {
				return err
			}

			songID := args[2]
			creationTime := time.Now()

			// MsgCreateSongPost will call SendPacket on src chain
			txs := relayer.RelayMsgs{
				Src: []sdk.Msg{
					c[src].PathEnd.MsgCreateSongPost(
						c[dst].PathEnd, dstHeader.GetHeight(),
						songID, creationTime, c[src].MustGetAddress(),
					),
				},
				Dst: []sdk.Msg{},
			}

			if txs.Send(c[src], c[dst]); !txs.Success() {
				return fmt.Errorf("failed to send first transaction")
			}

			// Working on SRC chain :point_up:
			// Working on DST chain :point_down:

			var (
				hs           map[string]*tmclient.Header
				seqRecv      chanTypes.RecvResponse
				seqSend      uint64
				srcCommitRes relayer.CommitmentResponse
			)

			if err = retry.Do(func() error {
				hs, err = relayer.UpdatesWithHeaders(c[src], c[dst])
				if err != nil {
					return err
				}

				seqRecv, err = c[dst].QueryNextSeqRecv(hs[dst].Height)
				if err != nil {
					return err
				}

				seqSend, err = c[src].QueryNextSeqSend(hs[src].Height)
				if err != nil {
					return err
				}

				srcCommitRes, err = c[src].QueryPacketCommitment(hs[src].Height-1, int64(seqSend-1))
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

			// reconstructing packet data here instead of retrieving from an indexed node
			packet := c[src].PathEnd.PostCreatePacket(
				songID, creationTime, c[src].MustGetAddress(),
				dstHeader.GetHeight()+1000,
			)

			// Debugging by simply passing in the packet information that we know was sent earlier in the SendPacket
			// part of the command. In a real relayer, this would be a separate command that retrieved the packet
			// information from an indexing node
			txs = relayer.RelayMsgs{
				Dst: []sdk.Msg{
					c[dst].PathEnd.UpdateClient(hs[src], c[dst].MustGetAddress()),
					c[dst].PathEnd.MsgRecvPacket(
						c[src].PathEnd,
						seqRecv.NextSequenceRecv,
						packet,
						chanTypes.NewPacketResponse(
							c[src].PathEnd.PortID,
							c[src].PathEnd.ChannelID,
							seqSend-1,
							c[src].PathEnd.NewPacket(
								c[dst].PathEnd,
								seqSend-1,
								packet,
							),
							srcCommitRes.Proof.Proof,
							int64(srcCommitRes.ProofHeight),
						),
						c[dst].MustGetAddress(),
					),
				},
				Src: []sdk.Msg{},
			}

			txs.Send(c[src], c[dst])
			return nil
		},
	}
	return pathFlag(cmd)
}
