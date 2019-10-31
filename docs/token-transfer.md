# Token transfer using IBC

Before starting transfer transactions, we need to setup ibc-chains, configurations and creation of client, connections. Follow this document to setup [IBC Instructions to Setup](config.md)

- refer the above document for up to connection handshake

#### Channel

> Note : 7 transactions between both chains

```bash
gaiacli \
  --home ibc0/n0/gaiacli \
  tx ibc channel handshake \
  send chan0 conn0 \
  receive chan1 conn1 \
  --node1 tcp://localhost:26657 \
  --node2 tcp://localhost:26557 \
  --chain-id2 ibc1 \
  --from1 n0 --from2 n1
```

#### Send Packet
> Note : this example shows transfer coins between two chains ibc0, ibc1

```bash
gaiacli --home ibc0/n0/gaiacli tx ibcsend transfer [to-address] [amount] [channel-id] --from n0
```

```bash
# example
gaiacli --home ibc0/n0/gaiacli tx ibcsend transfer cosmos15eegjmdn8hgtvma86srd0ugax9tmxqrajt2x2a 1000stake chan0 --from n0
```

After the successful transaction check the account balance in chain1 ibc0
```bash
gaiacli --home ibc0/n0/gaiacli q account [account-address] --chain-id [chain-id]
```
```bash
# example
gaiacli --home ibc0/n0/gaiacli q account cosmos15eegjmdn8hgtvma86srd0ugax9tmxqrajt2x2a   --chain-id ibc0
```
Query the account in chain ibc1, before Receiving Packet
> If you query the address in chain ibc1, It will says account not exist with given address.
```bash
gaiacli --home ibc1/n0/gaiacli q account cosmos15eegjmdn8hgtvma86srd0ugax9tmxqrajt2x2a --chain-id ibc1
```

#### Receiving Packet
- This will pull the packets from chain ibc0, using respective channel

```bash
gaiacli \
 --home ibc1/n0/gaiacli \
 tx ibc channel pull \
  [channel-type]  [channel-id] \
  --node1 tcp://0.0.0.0:26557 \
  --node2 tcp://0.0.0.0:26657 \
  --chain-id2 ibc0 \
   --from n1 
```

```bash
# example
gaiacli \
  --home ibc1/n0/gaiacli \
  tx ibc channel pull \
  receive chan1 \
  --node1 tcp://0.0.0.0:26557 \
  --node2 tcp://0.0.0.0:26657 \
  --chain-id2 ibc0 \
  --from n1 
```

After sucessful transaction check the balance of the given address in respctive to chain ibc1
```bash
gaiacli --home ibc1/n0/gaiacli q account [account-address] --chain-id [chain-id]
```
```bash
# example
gaiacli --home ibc1/n0/gaiacli q account cosmos15eegjmdn8hgtvma86srd0ugax9tmxqrajt2x2a --chain-id ibc1
```
