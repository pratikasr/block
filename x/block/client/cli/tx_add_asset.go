package cli

import (
	"strconv"

	"block/x/block/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-asset [name] [denom] [decimal] [ibc-status]",
		Short: "Broadcast message add-asset",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDenom := args[1]
			argDecimal := args[2]
			//argPrice := args[3]
			//argAppId := args[4]
			argIbcStatus := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddAsset(
				clientCtx.GetFromAddress().String(),
				argName,
				argDenom,
				argDecimal,
				argIbcStatus,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
