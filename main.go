package main

import (
	"fmt"
)

func KSA(key []byte) []byte {
	s := make([]byte, 256)
	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}
	j := 0
	for i := 0; i < 256; i++ {
		j = (j + int(key[i%len(key)]) + int(s[i])) % 256
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func PRGA(s []byte, plaintext []byte) []byte {
	i, j := 0, 0
	key := make([]byte, len(plaintext))
	for k := 0; k < len(plaintext); k++ {
		i = (i + 1) % 256
		j = (j + int(s[i])) % 256
		s[i], s[j] = s[j], s[i]
		key[k] = s[(int(s[i])+int(s[j]))%256]
	}
	return key
}

func RC4(key []byte, plaintext []byte) []byte {
	s := KSA(key)
	keyStream := PRGA(s, plaintext)
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ keyStream[i]
	}
	return ciphertext
}

func main() {
	key := []byte("MySecretKey")
	plaintext := []byte("Hello, RC4!")

	// Encryption
	ciphertext := RC4(key, plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Decryption (using the same key)
	decrypted := RC4(key, ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
