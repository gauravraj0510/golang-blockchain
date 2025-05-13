// wallets.go - Fixed version
package wallet

import (
	"bytes"
	"crypto/x509"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const walletFile = "./tmp/wallets.data"

// SerializedWallet is a serializable version of the Wallet struct
type SerializedWallet struct {
	PrivateKeyBytes []byte
	PublicKey       []byte
}

// Wallets holds a map of address to wallet
type Wallets struct {
	Wallets map[string]*Wallet
}

func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFile()
	// If the file doesn't exist, that's fine for a new wallet system
	if err != nil && os.IsNotExist(err) {
		return &wallets, nil
	}

	return &wallets, err
}

func (ws *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := string(wallet.Address())

	ws.Wallets[address] = wallet

	return address
}

func (ws *Wallets) GetAllAddresses() []string {
	var addresses []string

	for address := range ws.Wallets {
		addresses = append(addresses, address)
	}

	return addresses
}

func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

// Convert Wallet to SerializedWallet for encoding
func serializeWallet(w *Wallet) SerializedWallet {
	// Convert private key to DER-encoded format
	privateKeyBytes, err := x509.MarshalECPrivateKey(&w.PrivateKey)
	if err != nil {
		log.Panic(err)
	}
    
	return SerializedWallet{
		PrivateKeyBytes: privateKeyBytes,
		PublicKey:       w.PublicKey,
	}
}

// Convert SerializedWallet back to Wallet
func deserializeWallet(sw SerializedWallet) *Wallet {
	// Parse DER-encoded private key
	privateKey, err := x509.ParseECPrivateKey(sw.PrivateKeyBytes)
	if err != nil {
		log.Panic(err)
	}

	return &Wallet{
		PrivateKey: *privateKey,
		PublicKey:  sw.PublicKey,
	}
}

func (ws *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		return err
	}

	var serializedWallets map[string]SerializedWallet
	
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&serializedWallets)
	if err != nil {
		return err
	}

	// Convert SerializedWallet back to normal Wallet
	wallets := make(map[string]*Wallet)
	for address, serializedWallet := range serializedWallets {
		wallets[address] = deserializeWallet(serializedWallet)
	}
	
	ws.Wallets = wallets

	return nil
}

func (ws *Wallets) SaveFile() {
	// Create directory if it doesn't exist
	dir := filepath.Dir(walletFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Panic(err)
		}
	}

	// Convert Wallet map to SerializedWallet map for encoding
	serializedWallets := make(map[string]SerializedWallet)
	for address, wallet := range ws.Wallets {
		serializedWallets[address] = serializeWallet(wallet)
	}

	var content bytes.Buffer
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(serializedWallets)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}