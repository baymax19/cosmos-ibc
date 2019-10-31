package cli

import (
	"fmt"
	types2 "github.com/baymax19/cosmos-ibc/modules/ibc/bank/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "ibcsend",
		Short: "Quering Commands for ibcsend module",
	}

	queryCmd.AddCommand(
		GetCmdQueryName(queryRoute, cdc),
	)
	return queryCmd
}

func GetCmdQueryName(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "ownership [channel-id]",
		Short: "Query the User for channel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			val, _, err := ctx.QueryStore(types2.UserKey(args[0]), storeName)
			if err != nil {
				return err
			}

			var res string
			if val == nil {
				return fmt.Errorf("no user with this id", args[0])
			}

			cdc.MustUnmarshalBinaryBare(val, &res)

			fmt.Println(res)

			return nil
		},
	}
}
