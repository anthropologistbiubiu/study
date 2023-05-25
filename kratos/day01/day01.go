package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println(err)
	}
	val, err := rdb.Del("sun").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("val:%+v type:%+T\n", val, val)
	fmt.Println()
	arr := make([]int, 0)
	arr = append(arr, 1)
}
