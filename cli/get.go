package cli

import (
	"fmt"

	"github.com/ld86/gowallet/wallet"
)

func GetPassword(args []string) error {
	wallet := wallet.New()
	fmt.Println(wallet)
	return nil
}
