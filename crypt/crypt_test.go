package crypt

import (
	"testing"
)

// benchmark for EncryptString
func BenchmarkEncryptString10000(b *testing.B) {
	key := `Ho03neYb567Ad234G5e45r6DoNTcAR3e`
	text := `un:admin,pw:P@zZ\/\/0rD`
	for i := 0; i < 10000; i++ {
		EncryptString(key, text)
	}
}

// benchmark for DecryptString
func BenchmarkDecryptString10000(b *testing.B) {
	key := `Ho03neYb567Ad234G5e45r6DoNTcAR3e`
	text := `un:admin,pw:P@zZ\/\/0rD`
	for i := 0; i < 10000; i++ {
		DecryptString(key, text)
	}
}
