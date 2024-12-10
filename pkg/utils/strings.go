package util

import (
	"crypto/rand"
	"math/big"
)

// GetRandomStr 获取随机字符串
func GetRandomStr(length int) string {
	baseStr := "abcdefghijklmnopqrstuvwxyz0123456789"
	var randStr string
	max := len(baseStr)
	for i := 0; i < length; i++ {
		randStr = randStr + string(baseStr[GetRandom(int64(max-1))]) // fmt.Sprintf("%v%v", baseStr[randNum%max], randStr)
	}
	return randStr
}

// GetRandom 获取随机数
func GetRandom(length int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(length))
	return result.Int64()
}
