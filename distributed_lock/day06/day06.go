package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
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

// 分布式锁key
func (lock *Lock) key() string {
	return "lock_" + lock.tag
}

// 设置带超时的锁
func (lock *Lock) LockWithTimeout(client *Client, key string, timeout int) (ok bool, err error) {
	if timeout < 0 {
		timeout = 0
	}
	lock.tag = key
	lock.timeout = timeout
	lock.client = client
	ok, err = lock.tryLock()
	if !ok || err != nil {
		return
	}
	return
}

// 循环等待的锁，获取分布式锁失败，那就一直循环等待，一直循环等待maxWaitTime==-1
func (lock *Lock) LockWithWait(client *Client, key string, timeout, maxWaitTime int) (ok bool, err error) {
	now := time.Now()
	for {
		ok, err = lock.LockWithTimeout(client, key, timeout)
		if err != nil {
			fmt.Println("LockWithWait error, ", key, timeout, maxWaitTime)
			return
		}
		if ok == true {
			return
		}
		fmt.Println("LockWithWait fail, ", key, timeout, maxWaitTime)
		if maxWaitTime != -1 && int(time.Since(now).Seconds()) > maxWaitTime {
			fmt.Println("maxWaitTime reach, ", key, timeout, maxWaitTime)
			return
		}
		time.Sleep(time.Second * 1)
	}
}

// unlock
func (lock *Lock) Unlock() (err error) {
	err = lock.client.Del(lock.key()).Err()
	return
}

// lock
func (lock *Lock) Lock(client *Client, key string) (ok bool, err error) {
	return lock.LockWithTimeout(client, key, 1)
}
