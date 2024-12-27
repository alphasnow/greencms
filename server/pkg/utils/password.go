// @author AlphaSnow

package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(p string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hp), nil
}
func PasswordVerify(p string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	if err != nil {
		return false
	}
	return true
}
