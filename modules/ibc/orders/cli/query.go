package cli

import (
	"fmt"
	"github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"reflect"
)

func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "orders",
		Short: "Query commands for orders",
	}

	queryCmd.AddCommand(
		GetCmdQueryMakeOrder(queryRoute, cdc),
		GetCmdQueryMakeOrders(queryRoute, cdc),
		GetCmdQueryTakeOrders(queryRoute, cdc),
	)

	return queryCmd
}

func GetCmdQueryMakeOrder(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "make-order [order-hash]",
		Short: "Query Make Order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			data, _, err := ctx.QueryStore(types.GetMakeOrderKey(args[0]), storeName)
			if err != nil {
				return err
			}

			var res types.BaseMakeOrder
			if reflect.DeepEqual(data, res) {
				return fmt.Errorf("no order found")
			}
			cdc.MustUnmarshalBinaryBare(data, &res)

			_ = ctx.PrintOutput(res)
			return nil
		},
	}
}

func GetCmdQueryMakeOrders(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "make-orders ",
		Short: "Query Make Orders",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			data, _, err := ctx.QuerySubspace(append(types.OrderKey, types.MakeOrderKey...), storeName)
			if err != nil {
				return err
			}

			if data == nil {
				return fmt.Errorf("makeorders not found")
			}

			var res types.BaseMakeOrder
			for _, order := range data {
				cdc.MustUnmarshalBinaryBare(order.Value, &res)
				_ = ctx.PrintOutput(res)
			}

			return nil
		},
	}
}

func GetCmdQueryTakeOrders(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "take-orders [order-hash] ",
		Short: "Query TakeOrders for given MakeOrder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			data, _, err := ctx.QueryStore(types.GetTakeOrderKey(args[0]), storeName)
			if err != nil {
				return err
			}
			var res []types.BaseTakeOrder
			if reflect.DeepEqual(data, res) {
				return fmt.Errorf("no order found")
			}
			cdc.MustUnmarshalBinaryBare(data, &res)

			_ = ctx.PrintOutput(res)
			return nil
		},
	}

	return flags.GetCommands(cmd)[0]
}
