package encryption_test

import (
	"testing"

	"github.com/hazadus/go-lockbox/internal/encryption"
)

func TestEncryptDecrypt(t *testing.T) {
	stringToEncrypt := "String to encrypt"
	secret16 := "1234567890123456"

	encryptedString, err := encryption.Encrypt(stringToEncrypt, secret16)
	if err != nil {
		t.Fatal(err)
	}

	decryptedString, err := encryption.Decrypt(encryptedString, secret16)
	if err != nil {
		t.Fatal(err)
	}

	if decryptedString != stringToEncrypt {
		t.Errorf("Expected %q, got %q instead.", stringToEncrypt, decryptedString)
	}
}
