package main

import "github.com/redis/go-redis/v9"

// 研究一下stream 做消息队列 这个数据结构
// 按照那种简易的写法，来完成这个过程。

func NewRedisClient() *redis.Client {
	return redis.NewClient(*redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})
}
func main() {

}
