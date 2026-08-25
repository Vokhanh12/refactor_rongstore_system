package crypto

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	authsur "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	"golang.org/x/crypto/argon2"
)

var _ authsur.PasswordHasher = (*Argon2idHasher)(nil)

type Argon2idHasher struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func NewArgon2idHasher() *Argon2idHasher {
	return &Argon2idHasher{
		Memory:      64 * 1024, // 64 MiB
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
}

// Verify implements [security.PasswordHasher].
func (a *Argon2idHasher) Verify(hash string, password string) bool {
	params, salt, expectedHash, err := decodeArgon2id(hash)
	if err != nil {
		return false
	}

	actualHash := argon2.IDKey(
		[]byte(password),
		salt,
		params.iterations,
		params.memory,
		params.parallelism,
		uint32(len(expectedHash)),
	)

	return subtle.ConstantTimeCompare(actualHash, expectedHash) == 1
}

// Hash implements [security.PasswordHasher].
func (a *Argon2idHasher) Hash(password string) (string, error) {
	salt := make([]byte, a.SaltLength)

	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate password salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		a.Iterations,
		a.Memory,
		a.Parallelism,
		a.KeyLength,
	)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		a.Memory,
		a.Iterations,
		a.Parallelism,
		encodedSalt,
		encodedHash,
	), nil
}

type argon2idParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
}

func decodeArgon2id(encoded string) (
	argon2idParams,
	[]byte,
	[]byte,
	error,
) {
	var params argon2idParams

	parts := strings.Split(encoded, "$")

	if len(parts) != 6 {
		return params, nil, nil, fmt.Errorf("invalid argon2id hash")
	}

	if parts[1] != "argon2id" {
		return params, nil, nil, fmt.Errorf("unsupported password hash algorithm")
	}

	if parts[2] != "v=19" {
		return params, nil, nil, fmt.Errorf("unsupported argon2 version")
	}

	if _, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&params.memory,
		&params.iterations,
		&params.parallelism,
	); err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 parameters")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 salt")
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return params, nil, nil, fmt.Errorf("invalid argon2 hash")
	}

	return params, salt, expectedHash, nil
}
