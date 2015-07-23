package wallet

type Wallet struct {
	Nonce     string
	Passwords map[string]string
}

func New() *Wallet {
	return &Wallet{
		Nonce:     "1",
		Passwords: make(map[string]string),
	}
}
