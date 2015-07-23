package wallet

import "encoding/json"

type Wallet struct {
	Nonce     string
	Passwords map[string]string
}

func (wallet *Wallet) marshal() ([]byte, error) {
	return json.Marshal(wallet)
}

func unmarshal(data []byte) *Wallet {
	var wallet Wallet
	json.Unmarshal(data, &wallet)
	return &wallet
}

func New() *Wallet {
	return &Wallet{
		Nonce:     "1",
		Passwords: make(map[string]string),
	}
}

func ReadFromFile(filename string) *Wallet {
	return nil
}
