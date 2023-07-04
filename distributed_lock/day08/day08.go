package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
)

type Client struct {
	redis.UniversalClient
}

type Lock struct {
	timeout int
	client  *Client
	tag     string
}

func NewRedisClient() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.212.22.49:6379",
		DB:       0,
		Password: "",
	})

	err := client.Ping().Err()
	if err != nil {
		fmt.Println("client.Ping, ", err)
		return nil, errors.New("redis simple connect fail")
	}
	return &Client{client}, nil
}

// 如果需要过期时间精确到毫秒，可以使用PEXPIRE
var safeLock = redis.NewScript(`
	local key = KEYS[1]
	local r = redis.call("SETNX", key, 1)
	if (r == 0) then
		return 0
	end

	redis.call("EXPIRE", key, ARGV[1])
	return 1
`)

// lock
func (lock *Lock) tryLock() (bool, error) {
	var result int64
	result, err := safeLock.Run(lock.client, []string{lock.key()}, lock.timeout).Int64()
	if err == nil && result == 1 {
		return true, nil
	}
	return false, err
}

func main() {

	println("hello world")
}