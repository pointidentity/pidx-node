#!/bin/bash

echo "================================================"
echo "   Setup & Starting Pointidentity 3 Node Network    "
echo "================================================"

./setup.sh
./start.sh 
./fund.sh 
./createValidators.sh

echo "================================================"
echo "      PointIdentity 3 Node Network Setup End        " 
echo "================================================"




