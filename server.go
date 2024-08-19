package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net/http"
)

func main() {
	// Shared secret key
	secretKey := []byte("my_secret_key")

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		// Extract the HMAC signature
		signature := r.Header.Get("X-HMAC-Signature")

		// Recreate the canonical request
		canonicalRequest := fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Header.Get("Content-Type"))

		// Verify the HMAC signature
		mac := hmac.New(sha256.New, secretKey)
		mac.Write([]byte(canonicalRequest))
		expectedSignature := mac.Sum(nil)

		if !hmac.Equal([]byte(signature), expectedSignature) {
			http.Error(w, "Invalid HMAC signature", http.StatusUnauthorized)
			return
		}

		// Authenticate the request
		w.Write([]byte("Hello, authenticated user!"))
	})

	http.ListenAndServe(":8080", nil)
}
