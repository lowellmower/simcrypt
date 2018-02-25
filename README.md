## SimCrypt
SimCrypt is a tiny, simple, easy way to encrypt and decrypt strings in Go using
a shared key. The tool does one thing and does it quickly making it convenient
for things like storing sensitive data in memory in your applications.

## DISCLAIMER
This package is currently intended for my own use and curiosity. By no stretch of
the imagination should you blindly rely on this (or anything) for serious and 
stringent encryption needs. On a similar note, you should probably not adopt a 
cat without meeting it first. Oh, and, I wouldn't recommend eating soup at a
restaurant if you're a vegetarian and haven't seen an ingredient list. Oh oh 
oh... also... don't marry someone without meeting them before hand. Basically, 
if you just blindly trust something on the Internet to do something incredibly
important and you haven't kicked the tires and looked under the hood, you may
want to start doing that.

#### Usage (see below or `main.go`)
```go
package main

import "github.com/lowellmower/simcrypt/crypt"

func main() {
	key := "Ho03neYb567Ad234G5e45r6DoNTcAR3e" // at least 32 bytes long
	originalText := `My super secret message and password: P@sSw0rD`

	encryptedText, _ := crypt.EncryptString(key, originalText)
	decryptedText, _ := crypt.DecryptString(key, encryptedText)

	// originalText == decryptedText -> true
	// originalText != encryptedText -> false
```

#### Small and Speedy
10,000 Encryptions / Decryptions with 32 bit initialization vector:
```
crypt [master] :> go test -bench=BenchmarkEncryptString10000
BenchmarkEncryptString10000-4   	1000000000	         0.03 ns/op
PASS
ok  	github.com/lowellmower/simcrypt/crypt	0.200s
crypt [master] :> go test -bench=BenchmarkDecryptString10000
BenchmarkDecryptString10000-4   	2000000000	         0.00 ns/op
PASS
ok  	github.com/lowellmower/simcrypt/crypt	0.023s
```

#### Example Use Case:
An application needs to speak to an API to get some credentials or some other
type of sensitive data and then continue to use that information for a period of
time. A shared key can be used between the API and the client application to 
encypt the credentials before sending, allowing the client application to safely
hold the information encrypted in memory and decrypt it only when needed.

#### TODO:
- Make a CLI
- Make flags for cipher types
- Make flags for key byte length
- Support other types
- In package salting
- In package peppering
- Bake a pie
- Go to the gym
