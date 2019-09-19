package verge_addrdec

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/openwallet/openwallet"
)

const (
	btcAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	XVG_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x1e}, Suffix: nil}
	XVG_testnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x6F}, Suffix: nil}
	XVG_mainnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0x9e}, Suffix: []byte{0x01}}
	XVG_testnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0xEF}, Suffix: []byte{0x01}}
	XVG_mainnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x21}, Suffix: nil}
	XVG_testnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0xC4}, Suffix: nil}
	XVG_mainnetAddressBech32V0      = addressEncoder.AddressType{"bech32", btcAlphabet, "xvg", "h160", 20, nil, nil}
	XVG_testnetAddressBech32V0      = addressEncoder.AddressType{"bech32", btcAlphabet, "tb", "h160", 20, nil, nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	*openwallet.AddressDecoderV2Base
	IsTestNet bool
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := XVG_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = XVG_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := XVG_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = XVG_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)
	return address, nil
}
