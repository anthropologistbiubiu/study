package service

import (
	"crypto/md5"
	"encoding/hex"
)

// 我是想抽象出一个接口

type Md5Sign struct {
}

func (r *Md5Sign) GetServiceSign(data []byte) string {
	hash := md5.New()
	// 将数据写入哈希对象
	hash.Write(data)
	// 计算MD5哈希值
	hashValue := hash.Sum(nil)
	// 将哈希值转换为16进制字符串
	hashString := hex.EncodeToString(hashValue)
	return hashString
}
