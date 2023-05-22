package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnoqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUser(n int) string {
	return RandomString(n)
}

func RandomEmail(n int) string {
	return fmt.Sprintf("%s@email.com", RandomUser(n))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomPhoneNumber() string {
	return fmt.Sprintf("0%d%d%d%d%d%d%d%d", rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))
}
