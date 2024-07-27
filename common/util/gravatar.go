package util

import (
	"crypto/md5"
	"strings"
	"fmt"
)

func GetGravatarURL(email string) string {
	email = strings.TrimSpace(strings.ToLower(email))
	hash := md5.Sum([]byte(email))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?d=identicon", hash)
}
