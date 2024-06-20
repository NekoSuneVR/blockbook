package kyanite

import (
	"github.com/decenomy/blockbook/bchain"
	"github.com/decenomy/blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xf9a9f6a6
	TestnetMagic wire.BitcoinNet = 0x9f6f9a6a
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	MainNetParams.PubKeyHashAddrID = []byte{46}
	MainNetParams.ScriptHashAddrID = []byte{16}

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic

	// Address encoding magics
	TestNetParams.PubKeyHashAddrID = []byte{107}
	TestNetParams.ScriptHashAddrID = []byte{19}
}

type KyaniteParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

func NewKyaniteParser(params *chaincfg.Params, c *btc.Configuration) *KyaniteParser {
	return &KyaniteParser{BitcoinParser: btc.NewBitcoinParser(params, c), baseparser: &bchain.BaseParser{}}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		chaincfg.ResetParams()
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}

func (p *KyaniteParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

func (p *KyaniteParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
