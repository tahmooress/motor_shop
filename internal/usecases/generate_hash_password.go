package usecases

import (
	"crypto/sha1"
	"encoding/hex"
)

func (u *UseCases) generateHashPassword(password string) (string, error) {
	if password == "" {
		return "", nil
	}

	hash := sha1.Sum([]byte(password))

	return hex.EncodeToString(hash[:]), nil
}
