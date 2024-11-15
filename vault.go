package vaultclient

import (
	"log"

	"github.com/hashicorp/vault/api"
)

func NewVaultClient() (*api.Client, error) {
	address := "http://127.0.0.1:8200" // Host URL
	token := "XXXXXXXXXXXXXXX"         // Vault token

	// Initialize Vault client configuration
	config := api.DefaultConfig()
	config.Address = address

	// Create a new client instance
	client, err := api.NewClient(config)
	if err != nil {
		log.Printf("Failed to create Vault client: %v", err)
		return nil, err
	}

	// Set the token for client authentication
	client.SetToken(token)

	// Return the client instance
	return client, nil
}

func GetKey(key string) (string, error) {
	client, err := NewVaultClient()

	if err != nil {
		log.Fatalf("Failed to initialize Vault client: %v", err)
	}

	// Secret path
	secretPath := "kv/data/Keys"

	// Read the secret from the specified path
	secret, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Printf("Failed to read secret: %v", err)
		return "", err
	}

	if secret == nil || secret.Data["data"] == nil {
		log.Println("No secret found at specified path or unexpected format")
		return "", nil
	}

	// Retrieve the specified key from the nested data field
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Println("Data field in unexpected format")
		return "", nil
	}

	// Extract the specific key
	value, ok := data[key].(string)
	if !ok {
		log.Printf("Key %s not found or is not a string", key)
		return "", nil
	}

	// Successfully retrieved the key
	return value, nil
}
