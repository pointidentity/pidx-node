#!/bin/bash

# Use the already built binary or use a different one
BINARY=$1
if [ -z ${BINARY} ]; then
	BINARY=pidx-noded
fi

# Check if the binary is installed
${BINARY} &> /dev/null

RET_VAL=$?
if [ ${RET_VAL} -ne 0 ]; then
    echo "${BINARY} binary not found"
    exit 1
fi

CHAIN_NAMESAPCE=devnet

# Setting up config files
rm -rf /root/.pidx-node/

# Make directories for pidx-node config
mkdir /root/.pidx-node

# Init node
pidx-noded init --chain-id=pidxnode node1 --home=/root/.pidx-node

# Create key for the node
pidx-noded keys add node1 --keyring-backend=test --home=/root/.pidx-node

# change staking denom to upidx
cat /root/.pidx-node/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="upidx"' > /root/.pidx-node/config/tmp_genesis.json && mv /root/.pidx-node/config/tmp_genesis.json /root/.pidx-node/config/genesis.json

# update crisis variable to upidx
cat /root/.pidx-node/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="upidx"' > /root/.pidx-node/config/tmp_genesis.json && mv /root/.pidx-node/config/tmp_genesis.json /root/.pidx-node/config/genesis.json

# update gov genesis
cat /root/.pidx-node/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="upidx"' > /root/.pidx-node/config/tmp_genesis.json && mv /root/.pidx-node/config/tmp_genesis.json /root/.pidx-node/config/genesis.json
cat /root/.pidx-node/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="500s"' > /root/.pidx-node/config/tmp_genesis.json && mv /root/.pidx-node/config/tmp_genesis.json /root/.pidx-node/config/genesis.json

# update ssi genesis
cat /root/.pidx-node/config/genesis.json | jq '.app_state["ssi"]["chainNamespace"]="'$CHAIN_NAMESPACE'"' > /root/.pidx-node/config/tmp_genesis.json && mv /root/.pidx-node/config/tmp_genesis.json /root/.pidx-node/config/genesis.json

# update mint genesis
cat /root/.pidx-node/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="upidx"' > /root/.pidx-node/config/tmp_genesis.json && mv /root/.pidx-node/config/tmp_genesis.json /root/.pidx-node/config/genesis.json

# create validator node with tokens
pidx-noded add-genesis-account $(pidx-noded keys show node1 -a --keyring-backend=test --home=/root/.pidx-node) 110000000000upidx --home=/root/.pidx-node
pidx-noded gentx node1 100000000000upidx --keyring-backend=test --home=/root/.pidx-node --chain-id=pidxnode
pidx-noded collect-gentxs --home=/root/.pidx-node

# change app.toml values
sed -i -E '119s/enable = false/enable = true/' /root/.pidx-node/config/app.toml
sed -i -E '122s/swagger = false/swagger = true/' /root/.pidx-node/config/app.toml
sed -i -E '140s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' /root/.pidx-node/config/app.toml
sed -i -E 's|tcp://localhost:1317|tcp://0.0.0.0:1317|g' /root/.pidx-node/config/app.toml
sed -i -E 's|localhost:9090|0.0.0.0:9090|g' /root/.pidx-node/config/app.toml

# change config.toml values
sed -i -E 's|tcp://127.0.0.1:26657|tcp://0.0.0.0:26657|g' /root/.pidx-node/config/config.toml
sed -i -E 's|allow_duplicate_ip = false|allow_duplicate_ip = true|g' /root/.pidx-node/config/config.toml
sed -i -E 's|addr_book_strict = true|addr_book_strict = false|g' /root/.pidx-node/config/config.toml
sed -i -E 's|cors_allowed_origins = \[\]|cors_allowed_origins = \[\"\*\"\]|g' /root/.pidx-node/config/config.toml
