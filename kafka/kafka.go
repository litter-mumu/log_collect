package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer

func Init(address []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	//连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	return err
}

func SendToKafka(topic,data string) {
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	//发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v", pid, offset)
}
