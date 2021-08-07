package utils

import (
	"crypto/sha256"
	"fmt"
)

func GenSaltedPassword(password, salt string) string {
	sha1 := sha256.New()
	sha1.Write([]byte(password))
	shaPassword := fmt.Sprintf("%x", sha1.Sum(nil))

	sha2 := sha256.New()
	sha2.Write([]byte(shaPassword + salt))
	return fmt.Sprintf("%x", sha2.Sum(nil))
}
