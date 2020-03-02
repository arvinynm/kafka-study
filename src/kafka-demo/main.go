package main

import (
	"kafka-demo/consumer"
	"kafka-demo/producer"
)

func main() {
	//ch := make(chan int)
	go consumer.Test_Consumer()
	producer.Test_producer()
	//<-ch
}
