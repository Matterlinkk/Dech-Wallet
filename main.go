package main

import (
	"encoding/json"
	"fmt"
	"github.com/Matterlinkk/Dech-Wallet/keys"
	"math/big"
)

func main() {
	// Пример приватного ключа (ваша логика может использовать другие значения)
	pK := big.NewInt(10)

	// Получаем ключи на основе приватного ключа
	keyPair := keys.GetKeys(pK)

	// Маршалинг ключевой пары в JSON
	jsonKeys, err := json.Marshal(keyPair)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Выводим JSON как строку
	fmt.Println("JSON Representation:", string(jsonKeys))
}
