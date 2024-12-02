#!/bin/bash

# send upidx from first node to second node
sleep 7
pidx-noded tx bank send node1 $(pidx-noded keys show node2 -a --keyring-backend=test --home=$HOME/.pidx-node/node2) 500000000upidx --keyring-backend=test --home=$HOME/.pidx-node/node1 --chain-id=pidxnode --yes
sleep 7
pidx-noded tx bank send node1 $(pidx-noded keys show node3 -a --keyring-backend=test --home=$HOME/.pidx-node/node3) 400000000upidx --keyring-backend=test --home=$HOME/.pidx-node/node1 --chain-id=pidxnode --yes
