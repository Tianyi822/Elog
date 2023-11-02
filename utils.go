package easy_go_log

import (
	"crypto/sha256"
	"encoding/hex"
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
