### simcrypt
simcrypt is a tiny, simple, easy way to encrypt and decrypt strings in Go using
a shared key.

#### Usage (see below or `main.go`)
```go
package main

import "github.com/lowellmower/simcrypt/crypt"

func main() {
	key := "Ho03neYb567Ad234G5e45r6DoNTcAR3e"
	originalText := `My super secret message and password: P@sSw0rD`

	encryptedText, _ := crypt.EncryptString(key, originalText)
	decryptedText, _ := crypt.DecryptString(key, encryptedText)

	// originalText == decryptedText -> true
	// originalText != encryptedText -> false
```
