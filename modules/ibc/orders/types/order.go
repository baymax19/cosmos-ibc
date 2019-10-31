package types

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

type OrderStatus byte

const (
	StatusUnFilled  OrderStatus = 0x01
	StatusCancelled OrderStatus = 0x02
	StatusFilled    OrderStatus = 0x03
)

type BaseMakeOrder struct {
	MakerAddress types.AccAddress
	TakerAddress types.AccAddress

	MakerAsset types.Coin
	TakerAsset types.Coin

	OrderHash string

	OrderStatus OrderStatus
}

func (order BaseMakeOrder) String() string {
	return fmt.Sprintf(`
MakerAddress: %s,
TakerAddress: %s,
MakerAsset: %s,
TakerAsset: %s,
OrderHash:%s
OrderStatus: %d,
`,
		order.MakerAddress.String(), order.TakerAddress.String(), order.MakerAsset.String(), order.TakerAsset.String(),
		order.OrderHash, order.OrderStatus)
}

type BaseTakeOrder struct {
	TakerFillAmount types.Coin
	TakerAddress    types.AccAddress
	OrderHash       string
}

func (takeOrder BaseTakeOrder) String() string {
	return fmt.Sprintf(`
TakerFillAmount: %s,
TakerAddress: %s,
OrderHash: %s`, takeOrder.TakerFillAmount.String(), takeOrder.TakerAddress.String(),
		takeOrder.OrderHash)
}
func CalculateOrderHash(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil)[0:10])
}

type SignBytes struct {
	MakerAddress types.AccAddress
	MakerAsset   types.Coin
	TakerAsset   types.Coin
}

func NewSignBytes(address types.AccAddress, makerAsset, takerAsset types.Coin) SignBytes {
	return SignBytes{
		MakerAddress: address,
		MakerAsset:   makerAsset,
		TakerAsset:   takerAsset,
	}
}

func (b SignBytes) Bytes() []byte {
	return types.MustSortJSON(ModuleCdc.MustMarshalJSON(b))
}
