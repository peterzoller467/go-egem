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
"enode://107fac244340062e763fb52d4b2b2232599b0aa7b561176e81601fc2910cba6cffbecb23e30341e9ce8515c89026f751c01df3eb802ec9933a01f3743623224b@[154.20.195.166]:30666",
"enode://021cbc0b5eadb4fb5e424a44cffc1a5cf274307c2dfc8e3a9b32545bbd55b545dea97906acfb51c68bb6016c2a0fafae7df74d0a732fd4b10ce1e21e52d4f433@[199.247.0.245]:30666",
"enode://924bfab6dfde3dd5bf4d544497cc4fa2d4ce36d7c3e9d0db1d1ffe7072daa1923eab03f79503f145159df0eb0d851fa0c2bfcad10df75c66cb32ac5a3c387158@[45.32.43.8]:30666",
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
