package postgres

import (
	"crypto/sha1"
	"fmt"
	"os"
)


func hash(s string) string {
	sum := sha1.Sum([]byte(s + os.Getenv("HASH_SALT")))
	return fmt.Sprintf("%x", sum)
}