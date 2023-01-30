package cipher

import "golang.org/x/crypto/bcrypt"

func Hash(text string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(text), 9)
}

func Compare(hash []byte, text string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(text))
}
