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
// the main xPayments network.
var MainnetBootnodes = []string{
	// xPayments Foundation Go Bootnodes
	"enode://b2fa0155c7c4bb0921765de2753c46424f53b01b00ba1fbccff632cfb2a7ed4fe32406e4e176fb58cd92f68c4de2a062f9181558987cd6039a3557b9ef48c8ed@177.74.220.127:30340", // BR
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// the testnet Beryllium network.
var TestnetBootnodes = []string{
	"",
}
