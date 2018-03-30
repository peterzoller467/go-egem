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
"enode://68e928b4980ff70395cd1cf540d3f60efb5476f91f4d02f9057b25af042de7125c834d6fa06fe7c92bd573b871e4a7ec0f20e421d64e08fc317f5381849c8fd8@[154.20.195.166]:30666",
"enode://b991c3b233e1c535fa89cf690ab71327618eba701ad0f5e45704100ee4dbbb869b0e7243958c9a9e56d3f4c94b2135274836ea4ee4871423697fa6443ae7dfd5@[104.207.149.147]:30666",
"enode://2fcd625bdcc1001bd2cbcd8b22f8f4e08e8a2daa9b311e100ad99c3d26c99642cc9f57696b9e97d4be2052bafbbb330a5ddb92475e7331da433fca164f3801b7@[199.247.0.245]:30666",
"enode://41b7e5b7bf90218ab335245a84e2a33584bc7d6e8aaafffbe7a479ecca534870303f7e67a6aa8241fef4ec0dd6b6f7e2582d4c772371b04368578e96758c8548@[173.199.119.64]:30666",
"enode://58b67fb79911dabc4579190eeca64e5346cc57bdf28da2eac5fd92948e3e4b0081cc40add2b663d03cd4416a8a9e5f2e95226769abee7ccd665f1f83521c4ad1@[45.32.24.153]:30666",
"enode://56d32343155e869df3da89b07b3e3f990afc16fa561b09b09b1fccf8e65a0011dd090543f0b4a26c6239e2b77286b689d07b2a8fbd04e8a26684bfe3d223c3ba@[45.63.43.108]:30666",
"enode://31d526a27f1af85ecbf6f6629c21ae0c0ca7dd423f1cb57aaf2434896507c0fcc83ea42ac52656d693592c566836afc0e40d7a71d70c28807b2dd63a050d1558@[45.32.43.8]:30666",
"enode://e2f2471622b308217e513951ce75bfce412fe6a541ba49f4fa87784bad2602201f0bd32838089d9469bc0325237261b45b7e47a59fbf570d915186d03e891de9@[46.105.132.9]:30303",
"enode://316d2c8c55dd6100ede975696dd7a5fe94fc3310b88fecb98bf74385ac5447c55d199322d9db74f1cee9d301cf97715e7f3bc0642309db07f52f6b3edd95b549@[45.76.63.206]:30666",
"enode://69ebb51eb3ec7ca212be8a75a98ebf6cd8728e07e149d55e3ef48adb68c3179cdaf18457aa82556fd68f4cacd51b6b717df2c4c49f03466ef3f52c904088a8d3@[207.148.117.17]:30666",
"enode://f94ca913308d28d623fe9f79992c9baa0b4846dac4bff2bd2002b16394d2d380a8d3839409996725d595831b1484bb458795936523cb0660f0664d11192ddea1@[18.221.125.18]:30666",
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
