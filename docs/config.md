## Pre-Configuration

### Dependencies
```bash
go get github.com/baymax19/cosmos-ibc

make install
```

### Environment Setup

```bash
mkdir $HOME/ibc-testnets

cd $HOME/ibc-testnets

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
  $(gaiacli --home ibc1/n0/gaiacli/ query ibc client consensus-state -o json) \
  --from n0 -y -o text  


# client for chain ibc0 on chain ibc1
gaiacli --home ibc1/n0/gaiacli \
  tx ibc client create c1 \
  $(gaiacli --home ibc0/n0/gaiacli q ibc client consensus-state -o json) \
  --from n1 -y -o text
```

To query clients 
```bash
gaiacli --home ibc0/n0/gaiacli q ibc client client c0 --indent -o json
gaiacli --home ibc1/n0/gaiacli q ibc client client c1 --indent -o json
```

#### Connection Creation

> Note : 7 transactions between both chains

```bash
gaiacli \
  --home ibc0/n0/gaiacli \
  tx ibc connection handshake \
  conn0 c0 $(gaiacli --home ibc1/n0/gaiacli q ibc client path -o json) \
  conn1 c1 $(gaiacli --home ibc0/n0/gaiacli q ibc client path -o json) \
  --chain-id2 ibc1 \
  --from1 n0 --from2 n1 \
  --node1 tcp://localhost:26657 \
  --node2 tcp://localhost:26557
```

Once connection is established, you can query it:
```bash
gaiacli --home ibc0/n0/gaiacli q ibc connection connection conn0 --indent -o json
gaiacli --home ibc1/n0/gaiacli q ibc connection connection conn1 --indent -o json
```
