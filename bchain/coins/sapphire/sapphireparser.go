package sapphire

import (
	"github.com/decenomy/blockbook/bchain"
	"github.com/decenomy/blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0x583eb213
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	MainNetParams.PubKeyHashAddrID = []byte{63}
	MainNetParams.ScriptHashAddrID = []byte{18}
}

type SapphireParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

func NewSapphireParser(params *chaincfg.Params, c *btc.Configuration) *SapphireParser {
	return &SapphireParser{BitcoinParser: btc.NewBitcoinParser(params, c), baseparser: &bchain.BaseParser{}}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		chaincfg.ResetParams()
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	return &MainNetParams
}

func (p *SapphireParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

func (p *SapphireParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
