package cli

import (
	"fmt"
	"github.com/baymax19/cosmos-ibc/modules/bank/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "ibcsend",
		Short: "IBC Mock Send Transfer ownership of channel",
	}

	txCmd.AddCommand(
		UpdateUserTxCmd(cdc),
	)

	return txCmd
}

func UpdateUserTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ownership [channe-id] [name]  ",
		Short: "transfer ownership",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock)
			add, name, err := client.GetFromFields(ctx.GetFromName(), false)
			if err != nil {
				return err
			}

			ctx = ctx.WithFromAddress(add)
			ctx = ctx.WithFromName(name)
			fmt.Println("============================", ctx)

			fmt.Println(add, name)
			msg := types.NewMsgUser(args[1], args[0], add)

			return utils.GenerateOrBroadcastMsgs(ctx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}
