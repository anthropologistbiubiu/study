package sha

import (
	"crypto/sha256"
	"encoding/hex"
)

type shaSign struct{}

func (r *shaSign) GetSign(data []byte) string {
	// 创建SHA-256哈希对象
	hash := sha256.New()
	// 将数据写入哈希对象
	hash.Write(data)
	hashValue := hash.Sum(nil)
	// 将哈希值转换为16进制字符串
	hashString := hex.EncodeToString(hashValue)
	return hashString
}
