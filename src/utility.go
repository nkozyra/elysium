package elysium

import (
	"math/rand"
	"regexp"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func simpleGUID(s string) string {
	strings.Trim(s, " ")
	re := regexp.MustCompile("'")
	s = re.ReplaceAllString(s, "")
	re2 := regexp.MustCompile("[^A-Za-z0-9]+")
	s = re2.ReplaceAllString(s, "-")
	return strings.ToLower(s)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
