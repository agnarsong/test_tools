#!/bin/bash
export NODE_OPTIONS="--max-old-space-size=9000"
# cd ../../../packages/contracts/ && npx hardhat compile
# cd -
cd ../../../packages/customContracts/ && npx hardhat compile && npx hardhat export-bytecode --no-compile