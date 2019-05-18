/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package verge

import "github.com/blocktree/go-owcdrivers/btcTransaction"

const (
	//币种
	Symbol    = "XVG"
	MasterKey = "verge seed"
)

var (
	MainNetAddressPrefix = btcTransaction.AddressPrefix{[]byte{0x1e}, []byte{0x21}, "xvg"}
	TestNetAddressPrefix = btcTransaction.BTCTestnetAddressPrefix
)
