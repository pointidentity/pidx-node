#!/bin/bash

# create second validator
sleep 7
pidx-noded tx staking create-validator --amount=500000000upidx --from=node2 --pubkey=$(pidx-noded tendermint show-validator --home=$HOME/.pidx-node/node2) --moniker="node2" --chain-id="pidxnode" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="5000" --keyring-backend=test --home=$HOME/.pidx-node/node2 --yes
# create third validator
sleep 7
pidx-noded tx staking create-validator --amount=400000000upidx --from=node3 --pubkey=$(pidx-noded tendermint show-validator --home=$HOME/.pidx-node/node3) --moniker="node3" --chain-id="pidxnode" --commission-rate="0.1" --commission-max-rate="0.2" --commission-max-change-rate="0.05" --min-self-delegation="4000" --keyring-backend=test --home=$HOME/.pidx-node/node3 --yes


