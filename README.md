## Cosmos-IBC 

- [Cosmos-Referenece](https://github.com/cosmos/cosmos-sdk/tree/joon/ibc-sdk-interface-2) 

### Dependencies
```bash
ProjectDir : $GOPATH/src/github.com/baymax19/cosmos-ibc

git clone https://gitlab.com/baymax19/cosmos-ibc.git

git checkout transfer-token-implementation

make install
```

### Environment Setup

```bash
mkdir $HOME/ibc-testnets

gaiad testnet -o ibc0 --v 1 --chain-id ibc0 --node-dir-prefix n

gaiad testnet -o ibc1 --v 1 --chain-id ibc1 --node-dir-prefix n
```

### Setup `gaiad` and `gaiacli` Configuration

```bash
# Configure the proper database backend for each node and different listening ports
sed -i 's/"leveldb"/"goleveldb"/g' ibc0/n0/gaiad/config/config.toml
sed -i 's/"leveldb"/"goleveldb"/g' ibc1/n0/gaiad/config/config.toml
sed -i 's#"tcp://0.0.0.0:26656"#"tcp://0.0.0.0:26556"#g' ibc1/n0/gaiad/config/config.toml
sed -i 's#"tcp://0.0.0.0:26657"#"tcp://0.0.0.0:26557"#g' ibc1/n0/gaiad/config/config.toml
sed -i 's#"localhost:6060"#"localhost:6061"#g' ibc1/n0/gaiad/config/config.toml
sed -i 's#"tcp://127.0.0.1:26658"#"tcp://127.0.0.1:26558"#g' ibc1/n0/gaiad/config/config.toml

gaiacli config --home ibc0/n0/gaiacli/ chain-id ibc0
gaiacli config --home ibc1/n0/gaiacli/ chain-id ibc1
gaiacli config --home ibc0/n0/gaiacli/ output json
gaiacli config --home ibc1/n0/gaiacli/ output json
gaiacli config --home ibc0/n0/gaiacli/ node http://localhost:26657
gaiacli config --home ibc1/n0/gaiacli/ node http://localhost:26557
```
Add key from `ibc0/n0/gaiacli/key_seed.json` with `n0` name in both chains ibc0 and ibc1

Add key from `ibc1/n0/gaiacli/key_seed.json` with `n1` name in both chains ibc0 and ibc1

> Note : delete the existing keys in ibc0, ibc1 chains, <br/>
> Add  above keys using `gaiacli keys add --recover`

```bash

# Remove the key n0 on ibc1
gaiacli --home ibc1/n0/gaiacli keys delete n0

# seed from ibc0/n0/gaiacli/key_seed.json -> ibc0/n0
gaiacli --home ibc0/n0/gaiacli keys add n0 --recover

# seed from ibc1/n0/gaiacli/key_seed.json -> ibc0/n1
gaiacli --home ibc0/n0/gaiacli keys add n1 --recover

# seed from ibc0/n0/gaiacli/key_seed.json -> ibc1/n0
gaiacli --home ibc1/n0/gaiacli keys add n0 --recover

# seed from ibc1/n0/gaiacli/key_seed.json -> ibc1/n1
gaiacli --home ibc1/n0/gaiacli keys add n1 --recover


# check keys match
gaiacli --home ibc0/n0/gaiacli keys list | jq '.[].address'
gaiacli --home ibc1/n0/gaiacli keys list | jq '.[].address'
```

After configuration is complete, start `gaiad` in both home directories

```bash
gaiad --home ibc0/n0/gaiacli start

gaiad --home ibc1/n0/gaiacli start
```

### IBC Command Sequence

#### Client Creation

```bash
# client for chain ibc1 on chain ibc0
gaiacli --home ~/ibc-testnets/ibc0/n0/gaiacli/ \
  tx ibc client create c0 \
  $(gaiacli --home ibc1/n0/gaiacli/ query ibc client consensus-state) \
  --from n0 -y -o text  


# client for chain ibc0 on chain ibc1
gaiacli --home ibc1/n0/gaiacli \
  tx ibc client create c1 \
  $(gaiacli --home ibc0/n0/gaiacli q ibc client consensus-state) \
  --from n1 -y -o text
```

To query clients 
```bash
gaiacli --home ibc0/n0/gaiacli q ibc client client c0 --indent
gaiacli --home ibc1/n0/gaiacli q ibc client client c1 --indent
```

#### Connection Creation

> Note : 7 transactions between both chains

```bash
gaiacli \
  --home ibc0/n0/gaiacli \
  tx ibc connection handshake \
  conn0 c0 $(gaiacli --home ibc1/n0/gaiacli q ibc client path) \
  conn1 c1 $(gaiacli --home ibc0/n0/gaiacli q ibc client path) \
  --chain-id2 ibc1 \
  --from1 n0 --from2 n1 \
  --node1 tcp://localhost:26657 \
  --node2 tcp://localhost:26557
```

Once connection is established, you can query it:
```bash
gaiacli --home ibc0/n0/gaiacli q ibc connection connection conn0 --indent --trust-node
gaiacli --home ibc1/n0/gaiacli q ibc connection connection conn1 --indent --trust-node
```

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
