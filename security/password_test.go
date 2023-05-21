package security

import (
	"errors"
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {

	_, err := GeneratePassword("test")
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkGeneratePassword(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := GeneratePassword(fmt.Sprint(i))
		if err != nil {
			b.Error(err)
		}
	}

}

func TestDecodePassword(t *testing.T) {

	password, _ := GeneratePassword("test")

	_, _, _, err := DecodePassword(password)
	if err != nil {
		t.Error(err)
	}

}

func TestVerifyPassword(t *testing.T) {

	hashed, _ := GeneratePassword("test")

	err := VerifyPassword("test", hashed)
	if err != nil {
		t.Error(err)
	}

	err = VerifyPassword("", hashed)
	if !errors.Is(err, ErrInvalidPassword) {
		t.Errorf("expected error=%s, got=%s", ErrInvalidPassword, err)
	}
}
