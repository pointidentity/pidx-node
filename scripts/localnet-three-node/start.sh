#!/bin/bash

# Start all three nodes
tmux new -s node1 -d pidx-noded start --home=$HOME/.pidx-node/node1
tmux new -s node2 -d pidx-noded start --home=$HOME/.pidx-node/node2
tmux new -s node3 -d pidx-noded start --home=$HOME/.pidx-node/node3

