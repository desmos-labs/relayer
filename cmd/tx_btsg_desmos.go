package cmd

import (
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

			return c[src].SendPostBothSides(c[dst], args[2])
		},
	}
	return pathFlag(cmd)
}
