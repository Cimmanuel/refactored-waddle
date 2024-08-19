package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

func main() {
	// Shared secret key
	secretKey := []byte("my_secret_key")

	// Request data
	method := "GET"
	url := "https://example.com/api/data"
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body := []byte(`{"key": "value"}`)

	// Create a canonical request
	canonicalRequest := fmt.Sprintf("%s %s %s", method, url, headers["Content-Type"])

	// Sign the canonical request
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(canonicalRequest))
	signature := mac.Sum(nil)

	// Include the HMAC signature in the request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("X-HMAC-Signature", hex.EncodeToString(signature))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
