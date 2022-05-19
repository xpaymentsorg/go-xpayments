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
	// Ethereum Foundation Go Bootnodes
	"enode://b2fa0155c7c4bb0921765de2753c46424f53b01b00ba1fbccff632cfb2a7ed4fe32406e4e176fb58cd92f68c4de2a062f9181558987cd6039a3557b9ef48c8ed@170.82.241.117:30350",  // BR
	"enode://0279e35016791b272a0cdf721ec20db77827b806682df5b7dbb29810098b5511764cca849ecbe2c6dcf63222782a78d8f49787f65c811162f6992a6607fc8826@194.195.210.119:30350", // USA
}

// BerylliumBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Beryllium test network.
var BerylliumBootnodes = []string{
	"", // BR
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://b2fa0155c7c4bb0921765de2753c46424f53b01b00ba1fbccff632cfb2a7ed4fe32406e4e176fb58cd92f68c4de2a062f9181558987cd6039a3557b9ef48c8ed@170.82.241.117:30350",  // BR
	"enode://0279e35016791b272a0cdf721ec20db77827b806682df5b7dbb29810098b5511764cca849ecbe2c6dcf63222782a78d8f49787f65c811162f6992a6607fc8826@194.195.210.119:30350", // USA
}
