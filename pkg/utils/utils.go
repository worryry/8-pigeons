package util

import (
	"fmt"
	"time"
)

func StrToTimestamp(timeStr string) (int64, error) {
	t, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// ToString 转换成字符串
func ToString(mi interface{}) string {
	switch mi.(type) {

	case string:
		return mi.(string)
	case int, int32, int64, uint32, uint64:
		return fmt.Sprintf("%d", mi)
	}
	return fmt.Sprintf("%v", mi)
}

func GetInvCodeByUIDUniqueNew(uid int32, l int) string {
	const (
		PRIME1 = 3         // 与字符集长度 62 互质
		PRIME2 = 5         // 与邀请码长度 6 互质
		SALT   = 123456789 // 随意取一个数值
	)

	var AlphanumericSet = []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}

	// 放大 + 加盐。
	uid = uid*PRIME1 + SALT

	var code []rune
	slIdx := make([]byte, l)

	// 扩散。
	for i := 0; i < l; i++ {
		slIdx[i] = byte(uid % int32(len(AlphanumericSet)))                    // 获取 62 进制的每一位值
		slIdx[i] = (slIdx[i] + byte(i)*slIdx[0]) % byte(len(AlphanumericSet)) // 其他位与个位加和再取余（让个位的变化影响到所有位）
		uid = uid / int32(len(AlphanumericSet))                               // 相当于右移一位（62进制）
	}

	// 混淆。
	for i := 0; i < l; i++ {
		idx := (byte(i) * PRIME2) % byte(l)
		code = append(code, AlphanumericSet[slIdx[idx]])
	}
	return string(code)
}
