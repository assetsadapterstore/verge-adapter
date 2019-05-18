package openwtester

import (
	"github.com/assetsadapterstore/verge-adapter/verge"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(verge.Symbol, verge.NewWalletManager())
}
