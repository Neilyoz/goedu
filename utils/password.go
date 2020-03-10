package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword 加密密码
func HashPassword(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	return string(bytes), err
}

// CheckPasswordHash 比对密码
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
