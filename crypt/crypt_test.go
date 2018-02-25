package crypt

import (
	"testing"
)

var (
	key     = `Ho03neYb567Ad234G5e45r6DoNTcAR3e`
	text    = `un:admin,pw:P@zZ\/\/0rD`
	keyText = `
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.6
Comment: Hostname: pgp.mit.edu

mQGiBDhrtKIRBADKe46LzhEDLEAklMJn1fb0dP/zk7G/M9WRG+4SQbpvijSChC63Jpph8Ufg
5lWzr5VdsZCh9B15OQ3PVmG0+gcGplxQuGIN8WYlU0saIERs/h9ejg0yWvCo412FCP6xxAC6
fgxPSVbn3OWi/uFk1pGFzovoweTKiWn7XYVilGOk1QCg/yBC6/6rs6YcLr6Y9C0wzhF+7dkD
/27qyS7EO647OxXZS667d5RRQSjKFPnaymM/In0Bv6FMyZID/UAzZ0M6yGdUrXerh3bcLuIK
d98k4dGtJjqXB6+nKyZH64nCQRoGo6XYYoN/98q2ZcSfy3PwYLV3jLPSobzYtbsUmBOMq+Hu
7COMeS6/fO3p1h6keLQ9BNTUShgkBACTKMeVSj3a05ERDlAzkXHz2jqjyCcCQ+cIkXQ3gxMQ
uXQirpRPg6wT7S3ytGqmrwc6j1SzaXy/x7xzZ9iioat8ul1B7ZCo18iTDaMpXsD2Erm9G+0N
q1B6ZshYbHIOvUKmWU59Q9eJ3WWKaXzCkQvWWjhNsIMyNqHdmmIZ1o416bQoSmFzb24gRXVn
ZW5lIG1vd2VyIDxqZW1vd2VyQGFsbHdlc3QubmV0PohUBBARAgAUBQI4a7SiBQkAAqMABAsD
AgECGQEACgkQs7r3GQ7Ark8KhwCgoo5O/fdBT9FX8BkRTC7/J0/go/wAn20ofPtLENfoS3AM
JrYr1J1gpXRruQINBDhrtKIQCAD2Qle3CH8IF3KiutapQvMF6PlTETlPtvFuuUs4INoBp1aj
FOmPQFXz0AfGy0OplK33TGSGSfgMg71l6RfUodNQ+PVZX9x2Uk89PY3bzpnhV5JZzf24rnRP
xfx2vIPFRzBhznzJZv8V+bv9kV7HAarTW56NoKVyOtQa8L9GAFgr5fSI/VhOSdvNILSd5JEH
NmszbDgNRR0PfIizHHxbLY7288kjwEPwpVsYjY67VYy4XTjTNP18F1dDox0YbN4zISy1Kv88
4bEpQBgRjXyEpwpy1obEAxnIByl6ypUM2Zafq9AKUJsCRtMIPWakXUGfnHy9iUsiGSa6q6Je
w1XpMgs7AAICCADKEWvBdKiBNx7Sp5/Xnvjx8K4pDkwNSMnhh9fqBcQI1J1kwW2y9RdUtdhY
q7u1P92PqO8+YAB1pRsptcJfGX5YFRY0/TLAUg1kMaBitN1PyiuVrJveKsYpeneZkLCxp4+Z
mcFYsfKcu/qgi4Wo4RzCx/08ZA2qlY77GXwWL8tXdBAhlIVhJ6iNiuv9NMfuG5uYSwojna2a
AwQzoD4OiNMtmb9phkNwYtgNDPLovYmTgotBSuEAJ3LXcgvec4VJBmoNdOb0rxSneHHZiZx0
URtsWknlSTcPdp55u5RFEHJ+KseR6OSUJHh8ZYiEoyP9zuu8Kj/hpSHuwJRmofjB6YifiEwE
GBECAAwFAjhrtKIFCQACowAACgkQs7r3GQ7Ark+5+wCfblZ86Q5d4xEN8E+ENg1yJEREA7oA
n1OQItnPeQoAmz2n552q8bCK703R
=o6z8
-----END PGP PUBLIC KEY BLOCK-----`
)

// test that we can encrypt a string
func TestEncryptString(t *testing.T) {
	eString, err := EncryptString(key, keyText)
	if err != nil {
		t.Errorf("TestEncryptString failed with err: %s", err.Error())
		t.Fail()
	}

	if eString == keyText {
		t.Error("EncryptString failed to encrypt the string")
		t.Fail()
	}
}

// test that we can decrypt an encrypted string
func TestDecryptString(t *testing.T) {
	eString, _ := EncryptString(key, keyText)
	dString, err := DecryptString(key, eString)
	if err != nil {
		t.Errorf("TestDecryptString failed with err: %s", err.Error())
		t.Fail()
	}

	if dString != keyText {
		t.Error("DecryptString failed to match the original input")
		t.Fail()
	}
}

// test a key smaller than 32 bytes
// test a key greater than 32 bytes
// test a key exactly 32 bytes

// benchmark for EncryptString
func BenchmarkEncryptString10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		EncryptString(key, text)
	}
}

// benchmark for DecryptString
func BenchmarkDecryptString10000(b *testing.B) {
	for i := 0; i < 10000; i++ {
		DecryptString(key, text)
	}
}
