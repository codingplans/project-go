package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	// 创建消费者
	// consumer, err := sarama.NewConsumer(strings.Split("192.168.0.125:9092", ","), nil)
	consumer, err := sarama.NewConsumer(strings.Split("192.168.0.187:9092", ","), nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	// 设置分区
	// partitionList, err := consumer.Partitions("Deposit")
	partitionList, err := consumer.Partitions("kafeidou3")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}
	fmt.Println(partitionList)
	// 循环分区
	for partition := range partitionList {
		print(122223, partition)

		// pc, err := consumer.Consume
		// Partition("Deposit", int32(partition), sarama.OffsetNewest)
		pc, err := consumer.ConsumePartition("kafeidou3", int32(partition), sarama.OffsetNewest)

		println(pc, 9999999, err)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}

		}(pc)
	}
	// time.Sleep(time.Hour)
	wg.Wait()
	consumer.Close()
}
