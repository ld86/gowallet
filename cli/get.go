package cli

import (
	"fmt"

	"github.com/ld86/gowallet/wallet"
)

func GetPassword(args []string) error {
	w := wallet.New()
	fmt.Println(w)
	return nil
}
