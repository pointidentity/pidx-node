#!/bin/bash
rm -rf $HOME/.pidx-node/


# Make directories for pidx-node config
mkdir $HOME/.pidx-node
mkdir $HOME/.pidx-node/node1
mkdir $HOME/.pidx-node/node2
mkdir $HOME/.pidx-node/node3

# init all three nodes
pidx-noded init --chain-id=pidxnode node1 --home=$HOME/.pidx-node/node1
pidx-noded init --chain-id=pidxnode node2 --home=$HOME/.pidx-node/node2
pidx-noded init --chain-id=pidxnode node3 --home=$HOME/.pidx-node/node3

# Change pidx-node minimum gas prices
pidx-noded configure min-gas-prices 0upidx --home=$HOME/.pidx-node/node1
pidx-noded configure min-gas-prices 0upidx --home=$HOME/.pidx-node/node2
pidx-noded configure min-gas-prices 0upidx --home=$HOME/.pidx-node/node3

# create keys for all three nodes
pidx-noded keys add node1 --keyring-backend=test --home=$HOME/.pidx-node/node1
pidx-noded keys add node2 --keyring-backend=test --home=$HOME/.pidx-node/node2
pidx-noded keys add node3 --keyring-backend=test --home=$HOME/.pidx-node/node3

# change staking denom to upidx
cat $HOME/.pidx-node/node1/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="upidx"' > $HOME/.pidx-node/node1/config/tmp_genesis.json && mv $HOME/.pidx-node/node1/config/tmp_genesis.json $HOME/.pidx-node/node1/config/genesis.json

# create validator node with tokens to transfer to the three other nodes
pidx-noded add-genesis-account $(pidx-noded keys show node1 -a --keyring-backend=test --home=$HOME/.pidx-node/node1) 5000000000000000000upidx --home=$HOME/.pidx-node/node1
pidx-noded gentx node1 5000000000000000000upidx --keyring-backend=test --home=$HOME/.pidx-node/node1 --chain-id=pidxnode
pidx-noded collect-gentxs --home=$HOME/.pidx-node/node1

# update crisis variable to upidx
cat $HOME/.pidx-node/node1/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="upidx"' > $HOME/.pidx-node/node1/config/tmp_genesis.json && mv $HOME/.pidx-node/node1/config/tmp_genesis.json $HOME/.pidx-node/node1/config/genesis.json

# udpate gov genesis
cat $HOME/.pidx-node/node1/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="upidx"' > $HOME/.pidx-node/node1/config/tmp_genesis.json && mv $HOME/.pidx-node/node1/config/tmp_genesis.json $HOME/.pidx-node/node1/config/genesis.json

# update mint genesis
cat $HOME/.pidx-node/node1/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="upidx"' > $HOME/.pidx-node/node1/config/tmp_genesis.json && mv $HOME/.pidx-node/node1/config/tmp_genesis.json $HOME/.pidx-node/node1/config/genesis.json

#update ssi genesis
cat $HOME/.pidx-node/node1/config/genesis.json | jq '.app_state["ssi"]["chainNamespace"]="devnet"' > $HOME/.pidx-node/node1/config/tmp_genesis.json && mv $HOME/.pidx-node/node1/config/tmp_genesis.json $HOME/.pidx-node/node1/config/genesis.json

# change app.toml values

sed -i -E '104s/enable = false/enable = true/' $HOME/.pidx-node/node1/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.pidx-node/node1/config/app.toml

# node2
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:2317|g' $HOME/.pidx-node/node2/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9088|g' $HOME/.pidx-node/node2/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9089|g' $HOME/.pidx-node/node2/config/app.toml
sed -i -E '104s/enable = false/enable = true/' $HOME/.pidx-node/node2/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.pidx-node/node2/config/app.toml

# node3
sed -i -E 's|tcp://0.0.0.0:1317|tcp://0.0.0.0:3317|g' $HOME/.pidx-node/node3/config/app.toml
sed -i -E 's|0.0.0.0:9090|0.0.0.0:9086|g' $HOME/.pidx-node/node3/config/app.toml
sed -i -E 's|0.0.0.0:9091|0.0.0.0:9087|g' $HOME/.pidx-node/node3/config/app.toml
sed -i -E '104s/enable = false/enable = true/' $HOME/.pidx-node/node3/config/app.toml
sed -i -E '107s/swagger = false/swagger = true/' $HOME/.pidx-node/node3/config/app.toml

# change config.toml values

# node1
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.pidx-node/node1/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.pidx-node/node1/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.pidx-node/node1/config/config.toml

# node2
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:36658|g' $HOME/.pidx-node/node2/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:36657|g' $HOME/.pidx-node/node2/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:36656|g' $HOME/.pidx-node/node2/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.pidx-node/node2/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.pidx-node/node2/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.pidx-node/node2/config/config.toml

# node3
sed -i -E 's|tcp://127.0.0.1:26658|tcp://127.0.0.1:46658|g' $HOME/.pidx-node/node3/config/config.toml
sed -i -E 's|tcp://127.0.0.1:26657|tcp://127.0.0.1:46657|g' $HOME/.pidx-node/node3/config/config.toml
sed -i -E 's|tcp://0.0.0.0:26656|tcp://0.0.0.0:46656|g' $HOME/.pidx-node/node3/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.pidx-node/node3/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.pidx-node/node3/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.pidx-node/node3/config/config.toml

# copy node1 genesis file to node2 and node3
cp $HOME/.pidx-node/node1/config/genesis.json $HOME/.pidx-node/node2/config/genesis.json
cp $HOME/.pidx-node/node1/config/genesis.json $HOME/.pidx-node/node3/config/genesis.json


# Copy tendermint node id of node1 to persistent peers of node2 and node3
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(pidx-noded tendermint show-node-id --home=$HOME/.pidx-node/node1)@127.0.0.1:26656\"|g" $HOME/.pidx-node/node2/config/config.toml
sed -i -E "s|persistent_peers = \"\"|persistent_peers = \"$(pidx-noded tendermint show-node-id --home=$HOME/.pidx-node/node1)@127.0.0.1:26656\"|g" $HOME/.pidx-node/node3/config/config.toml



