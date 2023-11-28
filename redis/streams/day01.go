package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// 研究一下stream 做消息队列 这个数据结构
// 按照那种简易的写法，来完成这个过程。

var (
	streamKey = "my_stream"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func consumer(client *redis.Client, ctx context.Context, groupName string, consumerName string, streamkey string) {

	for {
		result, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamKey, ">"},
			Block:    0,
			Count:    1,
		}).Result()
		if err != nil {
			fmt.Println("XReadGroup", err)
		}
		for _, message := range result {
			for _, x := range message.Messages {
				for _, msg := range x.Values {
					log.Printf("consumer %s\n", msg)
				}
				client.XAck(ctx, streamkey, groupName, x.ID)
			}

		}
	}
}

func producer(client *redis.Client, ctx context.Context, streamkey string) {
	var i int
	for {
		message := fmt.Sprintf(" message %d", i)
		client.XAdd(ctx, &redis.XAddArgs{
			Stream: streamkey,
			Values: map[string]interface{}{"message": message},
		})
		i++
		log.Printf("producer message%s\n", message)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	client := NewRedisClient()
	ctx := context.Background()
	groupName := "mygroup"
	consumerName := "myconsumer"
	/*
		_, err := client.XGroupCreateMkStream(ctx, streamKey, groupName, "$").Result()
		if err != nil && err.Error() != "BUSY Consumer Group name already exists" {
			// 如果出错并且不是因为已存在，则打印错误信息
			fmt.Println("Error creating consumer group:", err)
			return
		}
	*/
	go consumer(client, ctx, groupName, consumerName, streamKey)
	go producer(client, ctx, streamKey)
	time.Sleep(10 * time.Second)

}

// 现在最重要的开始了，就是对消费者和生产者的协程序控制

// 1。 当遇到信号量，生产者停止写。
// 2。消费者停止读，但是消费者当前的协程数据需要等待被消费完成。
// 3。释放整个主进程。
