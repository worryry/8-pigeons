package util

import "fmt"

func PassSign(password, salt string) string {
	return EncodeMD5(fmt.Sprintf("%v.%v", password, salt))
}
