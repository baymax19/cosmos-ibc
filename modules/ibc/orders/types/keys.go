package types

var (
	OrderKey     = []byte{0x01}
	MakeOrderKey = []byte{0x02}
	TakeOrderKey = []byte{0x03}
)

func GetMakeOrderKey(orderHash string) []byte {
	return append(OrderKey, append(MakeOrderKey, []byte(orderHash)...)...)
}

func GetTakeOrderKey(orderHash string) []byte {
	return append(OrderKey, append(TakeOrderKey, []byte(orderHash)...)...)
}
