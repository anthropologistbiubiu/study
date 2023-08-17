package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"sync"
)

func consumer() {
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{"172.20.3.13:30901"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	partitionList, err := consumer.Partitions("test0") //获得该topic所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test0", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { //为每个分区开一个go协程去取值
			for msg := range pc.Messages() { //阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
}
func main1() {
	consumer()
}

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"127.0.0.1:9092"} // Kafka brokers
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		fmt.Println("errrrr", err)
	}
	defer consumer.Close()
	topic := "order-topic"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("NNNNNNNNNNN", err)
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println("Error:", err.Err)
		case <-signals:
			break ConsumerLoop
		}
	}
}
