package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 研究一下stream 做消息队列 这个数据结构
// 按照那种简易的写法，来完成这个过程。

var (
	streamKey = "my_stream"
	helper    = make(chan struct{})
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
					//log.Printf("consumer %s\n", msg)
					select {
					case <-helper:
						log.Println("consumer quit")
						return
					default:
						go func() {
							time.Sleep(time.Second * 2)
							log.Printf("consumer %s\n", msg)
							client.XAck(ctx, streamKey, groupName, x.ID)
						}()
					}
				}
			}

		}
	}
}

func producer(client *redis.Client, ctx context.Context, streamkey string) {
	var i int = 1
	for {
		select {
		case <-helper:
			log.Println("producer quit")
			return
		default:
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
}

func SiganlInit() {

}
func main() {
	signalChan := make(chan os.Signal, 0)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

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
	<-signalChan
	close(helper)
	time.Sleep(5 * time.Second)
	// 实践证明，只有主进程结束,协程序才结束。
}
