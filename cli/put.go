package cli

import "github.com/ld86/gowallet/wallet"

func PutPassword(args []string) error {
	w := wallet.ReadFromFile(wallet.WalletFilename)
	wallet.SaveToFile(wallet.WalletFilename, w)
	return nil
}
