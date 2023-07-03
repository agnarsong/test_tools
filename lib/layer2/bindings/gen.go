// Copyright 2022, Specular contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bindings

//go:generate ./compile.sh
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L1/messaging/L1StandardBridge.sol/L1StandardBridge.json --pkg bindings --type L1StandardBridge --out L1StandardBridge.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L2/messaging/L2StandardBridge.sol/L2StandardBridge.json --pkg bindings --type L2StandardBridge --out L2StandardBridge.go

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L1/messaging/L1CrossDomainMessenger.sol/L1CrossDomainMessenger.json --pkg bindings --type L1CrossDomainMessenger --out L1CrossDomainMessenger.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L2/messaging/L2CrossDomainMessenger.sol/L2CrossDomainMessenger.json --pkg bindings --type L2CrossDomainMessenger --out L2CrossDomainMessenger.go

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L1/rollup/CanonicalTransactionChain.sol/CanonicalTransactionChain.json --pkg bindings --type CanonicalTransactionChain --out CanonicalTransactionChain.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L1/rollup/StateCommitmentChain.sol/StateCommitmentChain.json --pkg bindings --type StateCommitmentChain --out StateCommitmentChain.go

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/customContracts/abi/contracts/L1CustomERC20.sol/CustomERC20.json --bin ../../../packages/customContracts/data/CustomERC20.bin --pkg bindings --type L1CustomERC20 --out L1CustomERC20.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/customContracts/abi/contracts/L2CustomERC20.sol/L2CustomERC20.json --bin ../../../packages/customContracts/data/L2CustomERC20.bin --pkg bindings --type L2CustomERC20 --out L2CustomERC20.go

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L1/fraud-proof/Rollup.sol/Rollup.json --pkg bindings --type Rollup --out Rollup.go

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi ../../../packages/contracts/abi/contracts/L2/predeploys/TssRewardContract.sol/TssRewardContract.json --pkg bindings --type TssRewardContract --out TssRewardContract.go
