package cage

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateKeyPair() (string, string, error) {
	// this is used for functions signServiceToServiceToken and verifyServiceToServiceSignature
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	pubKeyURLSafe := base64.URLEncoding.EncodeToString(pubKey)
	privKeyURLSafe := base64.URLEncoding.EncodeToString(privKey)
	return pubKeyURLSafe, privKeyURLSafe, nil
}

func SignPayload(privKeyURLSafe string, message []byte) (string, error) {
	privKey, err := base64.URLEncoding.DecodeString(privKeyURLSafe)
	if err != nil {
		return "", err
	}
	privateKey := ed25519.PrivateKey(privKey)
	signature := ed25519.Sign(privateKey, message)

	// Concatenate the message and signature, and encode as URL-safe Base64
	combined := append(message, signature...)
	token := base64.URLEncoding.EncodeToString(combined)
	return token, nil
}

func VerifySignature(pubKeyURLSafe string, token string) ([]byte, error) {
	pubKey, err := base64.URLEncoding.DecodeString(pubKeyURLSafe)
	if err != nil {
		return nil, err
	}
	publicKey := ed25519.PublicKey(pubKey)

	// Decode the URL-safe Base64 encoded combined string
	combined, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}

	// Separate the message and signature
	message := combined[:len(combined)-ed25519.SignatureSize]
	signature := combined[len(combined)-ed25519.SignatureSize:]

	if ed25519.Verify(publicKey, message, signature) {
		return message, nil
	}
	return nil, fmt.Errorf("signature verification failed")
}
