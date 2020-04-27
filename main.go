package main

import (
	"fmt"
	"hello_kafka_tailf/kafka"
	"hello_kafka_tailf/tailf"
	"time"
)

func run() {
	for {
		select {
		case line := <-tailf.ReadLog():
			kafka.SendToKafka("web_log", line.Text)
			break
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		return
	}
	fmt.Println("kafka init success")
	err = tailf.Init("./my.log")
	if err != nil {
		return
	}
	fmt.Println("kafka init success")
	run()
}
