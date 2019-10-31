## Token Exchange Protocol through Cosmos-IBC

 In this project we are providing very basic Implementation of token exchange protocol, using cosmos-ibc. This project  referenced from the [0x-protocol](https://0x.org/)

 ## Relayer Node
 - Relayer node which contains the list of `makeorders` and maintain the `orderbook` for the unfilled orders
  
### Makeorder
> sudo code snnipet
 
```go
type BaseMakeOrder struct{
	MakerAddress types.AccAddress
	TakerAddress types.AccAddress

	MakerAsset types.Coin
	TakerAsset types.Coin

	OrderHash string

	OrderStatus OrderStatus
}

const (
	StatusUnFilled  OrderStatus = 0x01
	StatusCancelled OrderStatus = 0x02
	StatusFilled    OrderStatus = 0x03
)

```
- `MakerAddress` : Address originating the order
- `TakerAddress` : Address of the reciever, initially it's empty
- `MakerAsset` : BaseAsset of the maker
- `TakerAsser` : QuoteAsset of the maker
- `OrderHash` : Identifing Unique Order 
- `OrderStatus` : Status of the order 

#### Future Improvemets
- Include `sigature` to verify the data
- point to point exchage of tokens when `TakeAddress` is present. (intra or inter blockchain)
- Include NFT to transfer
- `Expirationtime` to order to  cancel order

### TakeOrder
> sudo code snnipet
```go

type BaseTakeOrder struct {
	TakerFillAmount types.Coin
	TakerAddress    types.AccAddress
	OrderHash       string
}
```
- `TakerFillAmount` : BaseAssetAmount of the taker
- `TakerAddress`: Address of the Taker
- `OrderHash` : Unique hash to identify the `makeorder`

For given `makeOrder` we have list of `takeorders`, after receiving list of orders for given  `makeorder`, based on the `TakerFillAmount`, we will exchange the order. It's possible to submit interchain and intrachain `takeorders` for the given makeorder in our implementation.

#### Future Improvements
- sorting like algorithem to get the right `takeorder` for given `makeorder`
- use VDF{vdfProof, vdfIterations} to avoid the randomness in blockchain for selecting `takeorder`.


## Demo 
- `tx`: create makeOrder
- `tx`: submit the takeorder for given makeorder (orderhash)  using IBC
- `tx`: confirm makeOrder to exchage tokens.
- in this I am minting and burning token with respective chains, but we can also use escrow address to transfer tokens instead of burning.
 
### [Testing Demo](api-document-exchange.md)