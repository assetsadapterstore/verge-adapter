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

import (
	"encoding/hex"
	"github.com/assetsadapterstore/verge-adapter/verge_addrdec"
	"testing"
)

func TestAddressDecoder_AddressEncode(t *testing.T) {
	verge_addrdec.Default.IsTestNet = false

	p2pk, _ := hex.DecodeString("42fc23346faf777fa52dffb946f76e4120bd08a4")
	p2pkAddr, _ := verge_addrdec.Default.AddressEncode(p2pk)
	t.Logf("p2pkAddr: %s", p2pkAddr)

	p2sh, _ := hex.DecodeString("57b8d87fdb55ede830c99ff1c457f8ac8a0ac176")
	p2shAddr, _ := verge_addrdec.Default.AddressEncode(p2sh, verge_addrdec.XVG_mainnetAddressP2SH)
	t.Logf("p2shAddr: %s", p2shAddr)
}

func TestAddressDecoder_AddressDecode(t *testing.T) {

	verge_addrdec.Default.IsTestNet = false

	p2pkAddr := "DBFHBjwbSM8KwMzEDk3AVVHrKqnw1V3zLG"
	p2pkHash, _ := verge_addrdec.Default.AddressDecode(p2pkAddr)
	t.Logf("p2pkHash: %s", hex.EncodeToString(p2pkHash))

	p2shAddr := "ER9jiaULQT39X25BC2PR2V6d6kjhtkXBKH"

	p2shHash, _ := verge_addrdec.Default.AddressDecode(p2shAddr, verge_addrdec.XVG_mainnetAddressP2SH)
	t.Logf("p2shHash: %s", hex.EncodeToString(p2shHash))
}

func TestAddressDecoder_ScriptPubKeyToBech32Address(t *testing.T) {

	scriptPubKey, _ := hex.DecodeString("002079db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1")

	addr, err := tw.Decoder.ScriptPubKeyToBech32Address(scriptPubKey)
	if err != nil {
		t.Errorf("ScriptPubKeyToBech32Address failed unexpected error: %v\n", err)
		return
	}
	t.Logf("addr: %s", addr)


	t.Logf("addr: %s", addr)
}