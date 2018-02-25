package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// EncryptString takes a string (key) and a string (text)
// and encrypts the text using the key as the block size
// for an aes cipher.
func EncryptString(key, text string) (string, error) {
	if err := checkKeyLength(key); err != nil {
		return "", err
	}

	btext := []byte(text)
	bkey := []byte(key)[:32]

	block, err := aes.NewCipher(bkey)
	if err != nil {
		return text, err
	}

	cipherText := make([]byte, aes.BlockSize+len(btext))
	initVector := cipherText[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, initVector)
	if err != nil {
		return text, err
	}

	stream := cipher.NewCFBEncrypter(block, initVector)
	stream.XORKeyStream(cipherText[aes.BlockSize:], btext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// DecryptString takes a string (key) and a string (cryptoText)
// and returns a decrypted string and an error, the latter of
// which will be nil upon success.
func DecryptString(key, cryptoText string) (string, error) {
	if err := checkKeyLength(key); err != nil {
		return "", err
	}

	bkey := []byte(key)[:32]

	cipherText, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return cryptoText, err
	}

	block, err := aes.NewCipher(bkey)
	if err != nil {
		return cryptoText, err
	}

	if len(cipherText) < aes.BlockSize {
		err = fmt.Errorf("%s as cipher text too short", string(cryptoText))
		return cryptoText, err
	}
	initVector := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, initVector)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

// checkKeyLength ensures that a key being provided is at least 32
// bytes as aes.NewCipher will only produce a valid block from a key
// of 32 bytes or greater
func checkKeyLength(k string) error {
	bLen := len([]byte(k))
	if bLen < 32 {
		err := fmt.Errorf("Provided key was %d bytes, must be >= 32", bLen)
		return err
	}

	return nil
}
