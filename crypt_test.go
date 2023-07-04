package cage

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		fmt.Println("Error:", err)
		t.Fatal(err)
	}
	err = VerifyPassword(password, hash)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeyPair(t *testing.T) {
	pubKey, privKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pubKey)
	t.Log(privKey)

	// sign message with private key
	signature, err := SignServiceToServiceToken(privKey, []byte("message"))
	if err != nil {
		t.Fatal(err)
	}

	// verify signature
	msg, err := VerifyServiceToServiceToken(pubKey, signature)
	if err != nil {
		t.Fatal(err)
	}
	smg := string(msg)

	t.Log(smg)

}
