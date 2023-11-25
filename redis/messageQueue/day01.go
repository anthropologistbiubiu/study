package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// MessageQueue 封装了与 Redis 的交互
type MessageQueue struct {
	client *redis.Client
}

// NewMessageQueue 创建一个新的 MessageQueue 实例
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis 服务器地址
			Password: "",               // Redis 密码，如果没有密码则留空
			DB:       0,                // 默认使用的数据库
		}),
	}
}

// Enqueue 将消息推入队列
func (mq *MessageQueue) Enqueue(message string) {
	mq.client.LPush(context.Background(), "my_queue", message)
}

// Dequeue 从队列中获取消息
func (mq *MessageQueue) Dequeue() string {
	result, err := mq.client.BRPop(context.Background(), 0, "my_queue").Result()
	if err != nil {
		log.Println("Error dequeuing message:", err)
		return ""
	}
	return result[1]
}

func producer(queue *MessageQueue, messages []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, message := range messages {
		queue.Enqueue(message)
		fmt.Printf("Produced: %s\n", message)
		time.Sleep(1 * time.Second)
	}
}

func consumer(queue *MessageQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message := queue.Dequeue()
		if message != "" {
			fmt.Printf("Consumed: %s\n", message)
		} else {
			fmt.Println("No messages to consume.")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	queue := NewMessageQueue()
	var wg sync.WaitGroup
	wg.Add(2)
	// 启动生产者和消费者
	go producer(queue, []string{"Message 1", "Message 2", "Message 3"}, &wg)
	go consumer(queue, &wg)
	// 等待程序终止信号
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// 等待所有 goroutine 完成
	wg.Wait()
}
