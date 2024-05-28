package util

import (
	"crypto/rand"
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
)

// NewRandInt64 生成范围内整数
func NewRandInt64(max int64) (i int64, err error) {
	var bi *big.Int
	bi, err = rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return
	}
	i = bi.Int64()
	return
}

// NewRandInt 生成范围内整数
func NewRandInt(max int) (int, error) {
	i64, err := NewRandInt64(int64(max))
	return int(i64), err
}

// NewRandStringByByte 在指定 []byte 中随机抽取指定长度的字符串
func NewRandStringByByte(length int, chars []byte) (s string, err error) {
	charsLen := int64(len(chars))
	var k int64
	for i := 0; i < length; i++ {
		k, err = NewRandInt64(charsLen)
		if err != nil {
			return
		}
		s += string(chars[k])
	}
	return
}

// NewRandNumber 生成指定长度的随机字符串(只有数字)
func NewRandNumber(len int) (string, error) {
	b := []byte("0123456789")
	return NewRandStringByByte(len, b)
}

// NewRandString 生成指定长度的随机字符串(含有特殊字符)
func NewRandString(len int) (string, error) {
	b := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")
	return NewRandStringByByte(len, b)
}

// NewRandStringNoSpecial 生成指定长度的随机字符串(只有字母和数字)
func NewRandStringNoSpecial(len int) (string, error) {
	b := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	return NewRandStringByByte(len, b)
}

func RealRand(min, max int) int {
	rand2.Seed(int64(time.Now().Nanosecond()))
	//生成5000到10000之间的int32类型随机数
	//参数10000减5000保证函数生成的随机数在0到5000区间，
	//生成的随机数再加5000则落在5000到10000区间
	return rand2.Intn(max-min) + min
}

func GetFloatArr(coin float64, num, dNum int) (addCoin float64, extraC int64) {
	coinStr := decimal.NewFromFloat(coin).String()
	coinArr := strings.Split(coinStr, ".")
	if len(coinArr) <= 1 {
		addCoin = coin
	} else { //
		if len(coinArr[1]) <= num {
			addCoin = coin
			return
		} else {
			n := coinArr[1][:num]
			addD, _ := decimal.NewFromString(fmt.Sprintf("%s.%s", coinArr[0], n))
			addCoin, _ = addD.Float64()

			bNum := dNum - len(coinArr[1])
			extraStr := coinArr[1][num:]
			extraCTmp, _ := strconv.Atoi(extraStr)

			extraC = int64(extraCTmp)
			if bNum >= 1 {
				extraC = extraC * int64(math.Pow10(bNum))
			}
		}
	}
	return
}
