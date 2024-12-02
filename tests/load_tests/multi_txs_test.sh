#!/bin/bash

# Source Account
VALIDATOR_NODE=$(pidx-noded keys show node1 -a)
DESTINATION_ADDRESS="pidx17h4zk54wlayhfla7pdxxfqej3ra0rx5j6305hg"

# Tx Params
BASE_PRICE=10
ACCOUNT_SEQUENCE=$(($(echo $(pidx-noded q auth account ${VALIDATOR_NODE} --output json | jq '.sequence') | bc) - 1))
ITERATIONS=1000

for i in $(seq 1 ${ITERATIONS})
do
	SEND_AMT="$((${BASE_PRICE} + $i))upidx"
	SEQN=$((${SEQUENCE} + $i))
	pidx-noded tx bank send ${VALIDATOR_NODE} ${DESTINATION_ADDRESS} ${SEND_AMT} --chain-id pidxnode --sequence ${SEQN} --yes
done