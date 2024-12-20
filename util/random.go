package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Random Int dari min - max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random String sepanjang n huruf
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Random Owner name generator

func RandomOwner() string {
	return RandomString(6)
}

// Random Balance generator
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Random currency
func RandomCurrency() string {
	currency := []string{"EUR", "IDR", "USD"}
	n := len(currency)
	return currency[rand.Intn(n)]
}

func RandomEmail() string {
	email := fmt.Sprintf("%s@email.com", RandomString(6))
	return email
}
