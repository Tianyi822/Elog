package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenHash(str string) string {
	// 创建SHA256哈希对象
	hash := sha256.New()

	// 将字符串转换为字节数组并写入哈希对象
	hash.Write([]byte(str))

	// 计算哈希值
	hashBytes := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashStr := hex.EncodeToString(hashBytes)

	return hashStr
}

// DefaultTimeFormat 默认时间格式
func DefaultTimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// CustomTimeFormat 自定义时间格式
func CustomTimeFormat(t time.Time, format string) string {
	return t.Format(format)
}
