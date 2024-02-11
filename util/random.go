package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var localRand *rand.Rand

func init() {
	localRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// localRandを使用して乱数を生成します。
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
func RandomOwner() string {
	return RandomString(6)
}
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "YEN"}
	n := len(currencies)
	return currencies[localRand.Intn(n)]
}
