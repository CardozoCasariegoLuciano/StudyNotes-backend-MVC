package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Ibcrypt interface {
	Compare(hash string, plane string) error
	HashPassword(plane string) (string, error)
}

type Bcypt struct{}

func (bc *Bcypt) Compare(hash string, plane string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plane))
}

func (bc *Bcypt) HashPassword(plane string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(plane), bcrypt.DefaultCost)
	return string(hashedPass), err
}
