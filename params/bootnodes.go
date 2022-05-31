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

import "github.com/xpaymentsorg/go-xpayments/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main xPayments network.
var MainnetBootnodes = []string{
	// xPayments Foundation Go Bootnodes
	"enode://b2fa0155c7c4bb0921765de2753c46424f53b01b00ba1fbccff632cfb2a7ed4fe32406e4e176fb58cd92f68c4de2a062f9181558987cd6039a3557b9ef48c8ed@194.195.210.119:30340", // BR
}

// BerylliumBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Beryllium test network.
var BerylliumBootnodes = []string{
	// Beryllium Initiative bootnodes
	"", // BR
}

var V5Bootnodes = []string{
	// xPayments team's bootnode
	"",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case BerylliumGenesisHash:
		net = "beryllium"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
