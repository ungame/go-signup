package security

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/ungame/go-signup/security/internal"
	"time"
)

const (
	defaultAudience   = "auth"
	defaultIssuer     = "go_signup"
	defaultExpiration = time.Minute * 10
)

func NewToken(id string) (string, error) {

	claims := &jwt.StandardClaims{
		Audience:  defaultAudience,
		ExpiresAt: time.Now().Add(defaultExpiration).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    defaultIssuer,
		Subject:   id,
	}

	return newTokenWithClaims(claims)
}

func newTokenWithClaims(claims *jwt.StandardClaims) (string, error) {
	privateKey, err := internal.GetPrivateKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(privateKey)
}

func onParseToken(token *jwt.Token) (interface{}, error) {
	publicKey, err := internal.GetPublicKey()
	if err != nil {
		return nil, err
	}

	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("invalid jwt algorithm method: %s", token.Header["alg"])
	}

	return publicKey, nil
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {

	claims := new(jwt.StandardClaims)

	token, err := jwt.ParseWithClaims(tokenString, claims, onParseToken)
	if err != nil {
		return claims, err
	}

	if !token.Valid || claims == nil {
		return claims, jwt.NewValidationError("invalid token", jwt.ValidationErrorMalformed)
	}

	return claims, nil
}
