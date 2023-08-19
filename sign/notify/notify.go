package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"context"
	"github.com/Shopify/sarama"
)

var signals = make(chan os.Signal, 1)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Version = sarama.V2_3_0_0 // 指定 Kafka 版本
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建消费者组
	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "test-group", config)
	if err != nil {
		panic(err)
	}
	// 定义消费者组处理函数
	handler := ConsumerGroupHandler{}

	// 开始消费
	go func() {
		fmt.Println("start!")
		for {
			err := consumerGroup.Consume(context.Background(), []string{"test-topic"}, handler)
			if err != nil {
				panic(err)
			}
		}
	}()
	// 处理退出信号
	signal.Notify(signals, os.Interrupt)
	<-signals
	// 关闭消费者组
	consumerGroup.Close()
}

// 定义消费者组处理函数
type ConsumerGroupHandler struct{}

func (h ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		select {
		case <-signals:
			fmt.Println("shut down!")
		default:
			go func() {
				fmt.Printf("Message topic:%q partition:%d offset:%d value:%s\n",
					message.Topic, message.Partition, message.Offset, string(message.Value))
				time.Sleep(10 * time.Second)
				fmt.Println("rest for 10 second")
				session.MarkMessage(message, "")
			}()
		}
	}
	return nil
}
