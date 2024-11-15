// main.go

package main

import (
	"fmt"
	"log"

	"golang/vaultclient"
)

func main() {
	key := "WebKey"

	// Retrieve the specified key directly using GetKey
	value, err := vaultclient.GetKey(key)
	if err != nil {
		log.Fatalf("Failed to retrieve the specified key: %v", err)
	}

	fmt.Printf("Retrieved %s: %s\n", key, value)
}
