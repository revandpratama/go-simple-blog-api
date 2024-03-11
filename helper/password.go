package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(encryptedPassword), err
}

func VerifyPassword(encryptedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))

	return err
}
