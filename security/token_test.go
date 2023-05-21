package security

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	_, err := NewToken(uuid.NewString())
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkNewToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewToken(uuid.NewString())
		if err != nil {
			b.Error(err)
		}
	}
}

func TestParseToken(t *testing.T) {
	id := uuid.NewString()
	token, _ := NewToken(id)

	claims, err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}

	if claims.Subject != id {
		t.Errorf("unexpected token Subject: expected=%v, got=%v", id, claims.Subject)
	}

}

func TestParseTokenWithExpiredToken(t *testing.T) {
	id := uuid.NewString()

	token, _ := newTokenWithClaims(&jwt.StandardClaims{
		Audience:  defaultAudience,
		ExpiresAt: time.Now().Add(time.Second * -1).Unix(),
		Id:        id,
		IssuedAt:  time.Now().Add(defaultExpiration * -1).Unix(),
		Issuer:    defaultIssuer,
		Subject:   id,
	})

	_, err := ParseToken(token)
	fmt.Println(err)
	if err == nil {
		t.Error("unexpected non nil error")
	}

	jwtErr, ok := err.(*jwt.ValidationError)
	if !ok {
		t.Errorf("expected *jwt.ValidationError but received %T", err)
	}

	if jwtErr.Errors != jwt.ValidationErrorExpired {
		t.Errorf("epected expiration error but received %v", jwtErr.Errors)
	}
}
