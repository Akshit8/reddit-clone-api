// Package password impls methods for hashing
package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Hasher defines available methods for hashing
type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hashedpassword string) error
}

type nativeHasher struct {
	cost int
}

// NewNativeHasher creates new instance on nativeHasher
func NewNativeHasher() Hasher {
	return &nativeHasher{
		cost: bcrypt.DefaultCost,
	}
}

func (n *nativeHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), n.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func (n *nativeHasher) CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
