package internal

import (
	"crypto/rsa"
	_ "embed"
	"github.com/golang-jwt/jwt"
)

var (
	//go:embed jwtrsa.pem
	privateKeyBytes []byte

	//go:embed jwtrsa.pem.pub
	publicKeyBytes []byte
)

func GetPrivateKey() (*rsa.PrivateKey, error) {
	return jwt.ParseRSAPrivateKeyFromPEM(GetPrivateKeyBytes())
}

func GetPrivateKeyBytes() []byte {
	return cpy(privateKeyBytes)
}

func GetPublicKey() (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(GetPublicKeyBytes())
}

func GetPublicKeyBytes() []byte {
	return cpy(publicKeyBytes)
}

func cpy(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
