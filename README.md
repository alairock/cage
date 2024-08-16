# cage
Golang Cryptography for password hash and verify, and pub/priv key ed2559 sign/verify


## Examples

Hash and verify a password
```go
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
```

Create public/private key, sign with private key, verify with public key
```go
func TestKeyPair(t *testing.T) {
	pubKey, privKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pubKey)
	t.Log(privKey)

	// sign message with private key
	signature, err := SignPayload(privKey, []byte("message"))
	if err != nil {
		t.Fatal(err)
	}

	// verify signature
	msg, err := VerifySignature(pubKey, signature)
	if err != nil {
		t.Fatal(err)
	}
	smg := string(msg)

	t.Log(smg)

}
```
