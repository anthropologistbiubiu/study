package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"log"
)

// 这里主要是暴露一个服务
func main1() {
	r := gin.New()
	r.POST("/payment/create")
	r.POST("/payment/info")
	r.POST("/cashout/create")
	r.POST("/cashout/info")
	r.POST("/refound/create")
	r.POST("/refound/info")
	r.POST("/payment/notify")
	r.POST("/cashout/notify") // 将回调用请求转发到对应的服务
	r.Run(":5555")
}

func main() {
	// Kafka 服务器地址
	brokerList := []string{"localhost:9092"}

	// 创建配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	// 创建生产者
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	// 主题名称
	topic := "test-topic"

	// 要发送的消息
	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("key"),
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	// 发送消息
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
	} else {
		fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	}
}
