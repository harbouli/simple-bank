package utils

import (
	"math/rand"
	"strings"
	"time"
)

type Currency int

const (
	EUR Currency = iota // Euro
	USD                 // US Dollar
	MAD                 // Moroccan Dirham
)
const alphabet = "abcdefghijklmnopqrstezyuvxw"

var localRand *rand.Rand

func init() {
	localRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + localRand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[localRand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwnerName() string {

	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(100, 1000)
}
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "MAD"}
	n := len(currencies)
	return currencies[localRand.Intn(n)]

}
