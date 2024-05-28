package enkrip

import "golang.org/x/crypto/bcrypt"

type HashInterface interface {
	Compare(hashed, input string) error
	HashPassword(string) (string, error)
}

type Hash struct{}

func New() HashInterface {
	return &Hash{}
}

func (h *Hash) Compare(hashed, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}

func (h *Hash) HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}
