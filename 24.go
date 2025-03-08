package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	// 1. Генерация seed-фразы (BIP39) с 256-битной энтропией (24 слова)
	entropy, err := bip39.NewEntropy(256) // Увеличено до 256 бит
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seed phrase (24 words):", mnemonic)

	// 2. Создание HD-кошелька из seed-фразы
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Деривация первого Ethereum-адреса (m/44'/60'/0'/0/0)
	path := accounts.DefaultBaseDerivationPath
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Получение приватного ключа
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatal(err)
	}

	// 5. Генерация Ethereum-адреса
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	// 6. Вывод результата
	fmt.Println("Ethereum Address:", address)
	fmt.Println("Private Key:", fmt.Sprintf("%x", privateKey.D))
}
