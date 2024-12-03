#!/bin/bash

# Setup chain
echo "Setting up chain"
../../scripts/localnet-single-node/setup.sh
echo "Setup done"
echo ""

# Run the chain
echo "Running pidx-node"
echo ""
tmux new -s pidxnode -d pidx-noded start
sleep 5
if [[ -n $(pidx-noded status) ]]; then
  echo "pidx-noded daemon is now running"
  echo ""
else
  echo "pidx-noded daemon failed to start, exiting...."
  exit 1
fi