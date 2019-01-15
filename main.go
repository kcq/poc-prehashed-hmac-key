package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("Pre-hashed HMAC Key PoC...")

	fmt.Println("SHA256 Block Size:", sha256.BlockSize)
	data := []byte("data to protect")
	//The key must be bigger than the hash function block size
	key := []byte("secret.key:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Printf("key -> %v (%v)\n", string(key), len(key))

	h := sha256.New()
	h.Write(key)
	keyPH := make([]byte, sha256.BlockSize)
	copy(keyPH, h.Sum(nil))
	fmt.Printf("keyPH -> %x (%v)\n", keyPH, len(keyPH))

	hasher := hmac.New(sha256.New, key)
	hasher.Write(data)
	result := hasher.Sum(nil)
	fmt.Printf("HMAC(key,data)   -> %x\n", result)

	hasherPH := hmac.New(sha256.New, keyPH)
	hasherPH.Write(data)
	resultPH := hasherPH.Sum(nil)
	fmt.Printf("HMAC(keyPH,data) -> %x\n", resultPH)

	if hmac.Equal(result, resultPH) {
		fmt.Println("HMAC Match!")
	}

	fmt.Println("Pre-hashed HMAC Key PoC: done!")
}
