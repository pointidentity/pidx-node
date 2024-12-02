#!/bin/bash

set -e pipefail

# Set the binary name
BINARY=pidx-noded

# Check if the binary is installed
${BINARY} &> /dev/null

CHAINID=pidxnode
CHAIN_NAMESAPCE=devnet

RET_VAL=$?
if [ ${RET_VAL} -ne 0 ]; then
    echo "pidx-noded binary is not installed in your system."
    exit 1
fi

# Setting up config files
rm -rf $HOME/.pidx-node/

# Make directories for pidx-node config
mkdir $HOME/.pidx-node

# Init node
pidx-noded init --chain-id=$CHAINID node1 --home=$HOME/.pidx-node

# Create key for the node or recover existing key from mnemonic
if [[ -z "$MNEMONIC" ]]; then
    pidx-noded keys add node1 --keyring-backend=test --home=$HOME/.pidx-node
else
    echo $MNEMONIC | pidx-noded keys add node1 --keyring-backend=test --home=$HOME/.pidx-node --recover
fi

# change staking denom to upidx
cat $HOME/.pidx-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="upidx"' > $HOME/.pidx-node/config/tmp_genesis.json && mv $HOME/.pidx-node/config/tmp_genesis.json $HOME/.pidx-node/config/genesis.json

# update crisis variable to upidx
cat $HOME/.pidx-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="upidx"' > $HOME/.pidx-node/config/tmp_genesis.json && mv $HOME/.pidx-node/config/tmp_genesis.json $HOME/.pidx-node/config/genesis.json

# update gov genesis
cat $HOME/.pidx-node/config/genesis.json | jq '.app_state["gov"]["params"]["min_deposit"][0]["denom"]="upidx"' > $HOME/.pidx-node/config/tmp_genesis.json && mv $HOME/.pidx-node/config/tmp_genesis.json $HOME/.pidx-node/config/genesis.json
cat $HOME/.pidx-node/config/genesis.json | jq '.app_state["gov"]["params"]["voting_period"]="50s"' > $HOME/.pidx-node/config/tmp_genesis.json && mv $HOME/.pidx-node/config/tmp_genesis.json $HOME/.pidx-node/config/genesis.json

# update ssi genesis
cat $HOME/.pidx-node/config/genesis.json | jq '.app_state["ssi"]["chainNamespace"]="'$CHAIN_NAMESAPCE'"' > $HOME/.pidx-node/config/tmp_genesis.json && mv $HOME/.pidx-node/config/tmp_genesis.json $HOME/.pidx-node/config/genesis.json

# update mint genesis
cat $HOME/.pidx-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="upidx"' > $HOME/.pidx-node/config/tmp_genesis.json && mv $HOME/.pidx-node/config/tmp_genesis.json $HOME/.pidx-node/config/genesis.json

# create validator node with tokens
pidx-noded add-genesis-account $(pidx-noded keys show node1 -a --keyring-backend=test --home=$HOME/.pidx-node) 500000000000000000upidx --home=$HOME/.pidx-node --keyring-backend test
pidx-noded gentx node1 50000000000000000upidx --keyring-backend=test --home=$HOME/.pidx-node --chain-id=$CHAINID
pidx-noded collect-gentxs --home=$HOME/.pidx-node

# change app.toml values
sed -i -E '119s/enable = false/enable = true/' $HOME/.pidx-node/config/app.toml
sed -i -E '122s/swagger = false/swagger = true/' $HOME/.pidx-node/config/app.toml
sed -i -E 's|tcp://localhost:1317|tcp://0.0.0.0:1317|g' $HOME/.pidx-node/config/app.toml
sed -i -E 's|localhost:9090|0.0.0.0:9090|g' $HOME/.pidx-node/config/app.toml
sed -i -E '140s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $HOME/.pidx-node/config/app.toml

# change config.toml values
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' $HOME/.pidx-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' $HOME/.pidx-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' $HOME/.pidx-node/config/config.toml

echo -e "\nConfiguration set up is done, you are ready to run pidx-noded now!"

echo -e "\nPlease note the important chain configurations below:"

echo -e "\nRPC server address: http://localhost:26657"
echo -e "API server address: http://localhost:1317"
echo -e "DID Namespace: $CHAIN_NAMESAPCE"