package cli

import (
	"fmt"
	orderTypes "github.com/baymax19/cosmos-ibc/modules/ibc/orders/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"reflect"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "orders",
		Short: "IBC Mock for exchange orders b/w chains",
	}

	txCmd.AddCommand(
		MakeOrderCmd(cdc),
		TakerFillOrderTxCmd(cdc),
		ConfirmMakeOrderTxCmd(storeKey, cdc),
	)

	return txCmd
}

func MakeOrderCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make-order [maker-asset] [taker-asset]  ",
		Short: "create make order with base asset amount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock)

			makerAsset, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			takerAsset, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			bytes := orderTypes.NewSignBytes(ctx.GetFromAddress(), makerAsset, takerAsset).Bytes()
			orderHash := orderTypes.CalculateOrderHash(bytes)

			msg := orderTypes.NewMsgCreateMakeOrder(ctx.GetFromAddress(), makerAsset, takerAsset, orderHash, orderTypes.StatusUnFilled)

			return utils.GenerateOrBroadcastMsgs(ctx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(TakerAddress, "", "the receiver account")

	cmd = client.PostCommands(cmd)[0]
	return cmd
}

func TakerFillOrderTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fill-order [taker-fill-amount] [order-hash]",
		Short: "IBC taker order packet transfer to other chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock)

			accGetter := auth.NewAccountRetriever(ctx)

			account, err := accGetter.GetAccount(ctx.GetFromAddress())
			if err != nil {
				return err
			}
			chanID := viper.GetString(ChannelID)

			amount, err := sdk.ParseCoin(args[0])
			if err != nil {
				return err
			}

			hasCoin := account.GetCoins().AmountOf(amount.Denom)
			if hasCoin.IsZero() {
				return fmt.Errorf("insufficient funds")
			}

			msg := orderTypes.NewMsgTakerFillAmount(amount, ctx.GetFromAddress(), args[1], chanID)
			return utils.GenerateOrBroadcastMsgs(ctx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(ChannelID, "", "receiving channel id")
	cmd = client.PostCommands(cmd)[0]
	return cmd
}

func ConfirmMakeOrderTxCmd(orderStoreKey string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-order [order-hash]",
		Short: "confirm order will exchange the token b/w addresses, intra and inter blockchain; to do interblockchain specify the --channe-id flag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc).WithBroadcastMode(flags.BroadcastBlock)

			chanID := viper.GetString(ChannelID)

			data, _, err := ctx.QueryStore(orderTypes.GetMakeOrderKey(args[0]), orderStoreKey)
			if err != nil {
				return err
			}

			var res orderTypes.BaseMakeOrder
			if reflect.DeepEqual(data, res) {
				return fmt.Errorf("no order found")
			}
			cdc.MustUnmarshalBinaryBare(data, &res)

			if !res.MakerAddress.Equals(ctx.GetFromAddress()) {
				return fmt.Errorf("order not associated with signer address")
			}

			msg := orderTypes.NewMsgConfirmMakeOrder(ctx.GetFromAddress(), args[0], chanID)
			return utils.GenerateOrBroadcastMsgs(ctx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(ChannelID, "", "receiving channel id")
	cmd = client.PostCommands(cmd)[0]
	return cmd
}
