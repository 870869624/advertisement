package utils

import (
	"crypto/sha256"
	"fmt"
)

const (
	salt = "advertisement"
)

func Encrypt(input string) string {
	sum := sha256.Sum256([]byte(input))
	retult := sha256.Sum256([]byte(fmt.Sprintf("%x", sum) + salt))
	return fmt.Sprintf("%x", retult)
}
