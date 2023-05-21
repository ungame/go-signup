package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

var (
	ErrInvalidHash            = errors.New("invalid password hash")
	ErrInvalidPasswordVersion = errors.New("invalid password version")
	ErrInvalidPassword        = errors.New("invalid password")
)

const (
	defaultPasswordMemory      = 64 * 1024
	defaultPasswordIterations  = 3
	defaultPasswordParallelism = 2
	defaultPasswordSaltLength  = 16
	defaultPasswordKeyLength   = 32
)

type PasswordConfig struct {
	Version     uint32
	Memory      uint32
	Iterations  uint32
	Parallelism uint32
	SaltLength  uint32
	KeyLength   uint32
}

func GeneratePassword(password string) (string, error) {
	salt, err := RandomBytes(defaultPasswordSaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultPasswordIterations,
		defaultPasswordMemory,
		defaultPasswordParallelism,
		defaultPasswordKeyLength,
	)

	return EncodePassword(salt, hash), nil
}

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func EncodePassword(salt, hash []byte) string {
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultPasswordMemory,
		defaultPasswordIterations,
		defaultPasswordParallelism,
		b64Salt,
		b64Hash,
	)
}

func DecodePassword(encodedHash string) (*PasswordConfig, []byte, []byte, error) {

	throw := func(err error) (*PasswordConfig, []byte, []byte, error) {
		return nil, nil, nil, err
	}

	values := strings.Split(encodedHash, "$")

	if len(values) != 6 {
		return throw(ErrInvalidHash)
	}

	var cfg PasswordConfig

	_, err := fmt.Sscanf(values[2], "v=%d", &cfg.Version)
	if err != nil {
		return throw(err)
	}

	if cfg.Version != argon2.Version {
		return throw(ErrInvalidPasswordVersion)
	}

	_, err = fmt.Sscanf(
		values[3],
		"m=%d,t=%d,p=%d",
		&cfg.Memory,
		&cfg.Iterations,
		&cfg.Parallelism,
	)
	if err != nil {
		return throw(err)
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return throw(err)
	}

	cfg.SaltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return throw(err)
	}

	cfg.KeyLength = uint32(len(hash))

	return &cfg, salt, hash, nil
}

func VerifyPassword(password, encodedPassword string) error {
	cfg, salt, hash, err := DecodePassword(encodedPassword)
	if err != nil {
		return err
	}

	maybeTheSameHash := argon2.IDKey(
		[]byte(password),
		salt,
		cfg.Iterations,
		defaultPasswordMemory,
		defaultPasswordParallelism,
		defaultPasswordKeyLength,
	)

	if subtle.ConstantTimeCompare(hash, maybeTheSameHash) == 1 {
		return nil
	}

	return ErrInvalidPassword
}
