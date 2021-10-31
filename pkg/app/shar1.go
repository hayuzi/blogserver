package app

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func EncodeSha1(value string) string {
	m := sha1.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func EncodePwd(value string) string {
	prefix := "blogPwd"
	encrypted := EncodeSha1(fmt.Sprintf("%s%s", prefix, value))
	return string([]rune(encrypted)[:32])
}
