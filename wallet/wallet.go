package wallet

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"syscall"
)

var WalletFilename = os.Getenv("HOME") + "/.wallet"

type Wallet struct {
	Nonce     string
	Passwords map[string]string
}

func (wallet *Wallet) marshal() ([]byte, error) {
	return json.Marshal(wallet)
}

func unmarshal(data []byte) (*Wallet, error) {
	var wallet Wallet
	err := json.Unmarshal(data, &wallet)
	return &wallet, err
}

func New() *Wallet {
	return &Wallet{
		Nonce:     "1",
		Passwords: make(map[string]string),
	}
}

func ReadFromFile(filename string) *Wallet {
	content, readErr := ioutil.ReadFile(filename)
	if e, ok := readErr.(*os.PathError); ok && e.Err == syscall.ENOENT {
		return New()
	}
	w, unmarshalErr := unmarshal(content)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
	return w
}

func SaveToFile(filename string, w *Wallet) {
	content, marshalErr := w.marshal()
	if marshalErr != nil {
		panic(marshalErr)
	}
	writeErr := ioutil.WriteFile(filename, content, 0600)
	if writeErr != nil {
		panic(writeErr)
	}
}
