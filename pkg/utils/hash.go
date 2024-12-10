package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func EncodeHash256(value string) string {
	h := sha256.New()
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

func HmacSha256(value, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(value))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
	//return hex.EncodeToString(h.Sum(nil))
}
