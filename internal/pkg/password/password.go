// 2. PASSWORD HELPER
// File: internal/pkg/password/password.go

package password

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
