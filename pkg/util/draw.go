package util

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"strings"
	"time"
)

func Lottery(probability float64, base int64) bool {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(base) + 1
	dP := decimal.NewFromFloat(probability)
	dB := decimal.NewFromInt(base)
	limit := dP.Mul(dB).IntPart()
	if randNum <= limit {
		return true
	}
	return false
}

// [0,n)
func RandIndex(base int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(base)
	return randNum
}

func RandRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min+1) + min
	return randNum
}

func ArrayToString(arr []string) string {
	// 使用strings.Join函数直接将整数转换为逗号隔开的字符串
	return strings.Join(strings.Fields(fmt.Sprint(arr)), ",")
}
