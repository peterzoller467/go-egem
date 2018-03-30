// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
// Egem Go Bootnodes
"enode://91db08c0148d571a75e4975ecf9e2f0c0edc028bd0576108ef0487f86ae1fbf4f7cad4a53ae90a106fb8b0f19247c078dd66f5019da5c9ae7ee2290c06d75ecf@[154.20.195.166]:30666",
"enode://b991c3b233e1c535fa89cf690ab71327618eba701ad0f5e45704100ee4dbbb869b0e7243958c9a9e56d3f4c94b2135274836ea4ee4871423697fa6443ae7dfd5@[104.207.149.147]:30666",
"enode://30addb02c9d61f88f67384e03cadc98e95beddc45d7bad05c6bb6e166d52250529cb8fea3e34c4406b2e71e85acc07ba4d3163562a0917888e4580ad60d6744f@[199.247.0.245]:30666",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// RoGem test network.
var TestnetBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

}
