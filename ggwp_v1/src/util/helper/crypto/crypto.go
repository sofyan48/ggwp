package crypto

import "golang.org/x/crypto/bcrypt"

// HashPassword params
// @password: string
// return string, error
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash params
// @password: string
// @hash: string
// return bool
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
