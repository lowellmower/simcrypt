package main

import (
	"fmt"

	"github.com/lowellmower/simcrypt/crypt"
)

func main() {
	key := "Ho03neYb567Ad234G5e45r6DoNTcAR3e"
	originalText := `My super secret message and password: P@sSw0rD`

	// take a look at the original string
	fmt.Printf("Original string: %s\n\n", originalText)

	// encrypt using the secret key
	cryptoText, err := crypt.EncryptString(key, originalText)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error during EncryptEtring: %s", err.Error()))
	}

	// take a look at the encrypted string
	fmt.Printf("Encrypted string: %s\n\n", cryptoText)

	// decrypt to original value using same key
	decryptedText, err := crypt.DecryptString(key, cryptoText)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error during DecryptString: %s", err.Error()))
	}

	// take a look at the decrypted string
	fmt.Printf("Decrypted string: %s\n\n", decryptedText)

	// are the strings interpreted as equals?
	fmt.Printf("Are the strings interpreted as equal? - %v\n", decryptedText == originalText)
}
