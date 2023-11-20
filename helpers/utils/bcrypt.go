package utils

import (
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type Ibcrypt interface {
	Compare(hash string, plane string) error
	HashPassword(plane string) (string, error)
}

type Bcypt struct{}

var bc *Bcypt
var onceBc sync.Once

func NewBcrypy() Ibcrypt {
	onceBc.Do(func() {
		bc = &Bcypt{}
	})
	return bc
}

func (bc *Bcypt) Compare(hash string, plane string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plane))
}

func (bc *Bcypt) HashPassword(plane string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(plane), bcrypt.DefaultCost)
	return string(hashedPass), err
}
