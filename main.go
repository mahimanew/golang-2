package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// Path to the injected secret file
	secretPath := os.Getenv("VAULT_SECRET_PATH")

	data, err := os.ReadFile(secretPath)
	if err != nil {
		fmt.Printf("Error reading secret file: %v\n", err)
		return
	}

	fmt.Printf("Secret content: %s\n", string(data))

	fmt.Fprintf(w, "Hello, Dockerized Go App! %s ", string(data))

}

func main() {
	// http.HandleFunc("/", handler)
	secretPath := os.Getenv("VAULT_SECRET_PATH")

	data, err := os.ReadFile(secretPath)
	if err != nil {
		fmt.Printf("Error reading secret file: %v\n", err)
		return
	}

	fmt.Printf("Secret content: %s\n", string(data))
	fmt.Println("heloooooo.....")
	fmt.Println("Server is running on port 9090")
	http.ListenAndServe(":9090", nil)
}
