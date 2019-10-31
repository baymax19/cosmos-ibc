# Token exchange using IBC

Before starting order transactions, we need to setup ibc-chains, configurations and creation of client, connections. Follow this document to setup [IBC Instructions to Setup](config.md)

- refer the above document for up to connection handshake


#### Channel Creation

> Note : 7 transactions between both chains

```bash
gaiacli \
  --home ibc1/n0/gaiacli \
  tx ibc channel handshake \
  orders chan1 conn1 \
  ordersrecv chan0 conn0 \
  --node1 tcp://localhost:26557 \
  --node2 tcp://localhost:26657 \
  --chain-id2 ibc0 \
  --from1 n1 --from2 n0
```

```bash

gaiacli \
  --home ibc0/n0/gaiacli \
  tx ibc channel handshake \
  orders chan0 conn0 \
  ordersrecv chan1 conn1 \
  --node1 tcp://localhost:26657 \
  --node2 tcp://localhost:26557 \
  --chain-id2 ibc1 \
  --from1 n0 --from2 n1

```

##### Query Channels

```bash
gaiacli --home ibc0/n0/gaiacli query ibc channel channel ordersrecv chan0 --indent -o json
gaiacli --home ibc1/n0/gaiacli query ibc channel channel orders chan1 --indent -o json
```

##### Query account
```bash
gaiacli --home ibc0/n0/gaiacli keys list
gaiacli --home ibc0/n0/gaiacli query account [account-address]
```


#### Create MakeOrder
> Note :  intra chain transaction.
```bash
gaiacli --home ibc0/n0/gaiacli/ tx \
     orders make-order 1stake 2usd --from n0
```
##### Query MakeOrder
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders make-order [orderhash]  -o json
```
>Example
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders make-order 8b35d253dbc866a19ef7  -o json
```

To get all `makeorders`
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders make-orders -o json 
```


#### TakeOrder 

To Query `takeOrders` for given `makeOrder`
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders take-orders [orderhash] -o json 
```
>Example
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders take-orders 8b35d253dbc866a19ef7 -o json
```
##### Inter chain transaction to ibc0 from ibc1

```bash
gaiacli --home ibc1/n0/gaiacli/ tx orders fill-order [take-fill-amount] [orderhash] --channel-id [channel-id] --from n1
```
It will send the `PacketTakeOrder` to ibc0 chain.

> Example:
```bash
gaiacli --home ibc1/n0/gaiacli/ tx orders fill-order 5usd 8b35d253dbc866a19ef7 --channel-id chan1 --from n1
```


##### Recieve TakeOrder Packet in ibc0
```bash
gaiacli --home ibc0/n0/gaiacli tx ibc channel pull ordersrecv chan0 --from n0 --node1 tcp://0.0.0.0:26657 --node2 tcp://0.0.0.0:26557 --chain-id2 ibc1
```
It will pull the Packet and added to `takeOrders` list for given `makeorder`
>Note : we will get the takeorder from ibc1 here

To Query
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders take-orders [order-hash] 
```
>Example
```bash
gaiacli --home ibc0/n0/gaiacli/ q orders take-orders 8b35d253dbc866a19ef7 -o json
```

#### Confirm MakeOrder
- We will take the top one of `takeorder` which is sorted based on the AssetAmount.
>Note : Here we are confirming IBC takeOrder 
>
>if channel-id flag not specified then it considered as intra blockchain transaction.
```bash
gaiacli --home ibc0/n0/gaiacli/ tx orders confirm-order [orderhas] --from n0 --channel-id [channel-id]
```

Example
```bash
gaiacli --home ibc0/n0/gaiacli/ tx orders confirm-order 8b35d253dbc866a19ef7 --from n0 --channel-id chan0
```

##### Receive exchange packet in ibc1 chain
```bash
gaiacli --home ibc1/n0/gaiacli tx ibc channel pull ordersrecv chan1 --from n1 --node1 tcp://0.0.0.0:26557 --node2 tcp://0.0.0.0:26657 --chain-id2 ibc0
```

##### Query Account in both chains
```bash
gaiacli --home ibc0/n0/gaiacli keys list
gaiacli --home ibc0/n0/gaiacli query account [account-address]

gaiacli --home ibc1/n0/gaiacli keys list
gaiacli --home ibc1/n0/gaiacli query account [account-address]
```

