package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"kafka-demo/config"
)

func Test_Consumer(){
	config := config.Config()

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	x, err := consumer.Partitions("kafka-go-test1")
	fmt.Println(x)

	partition_consumer, err := consumer.ConsumePartition("kafka-go-test1", x[1], sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg:=<-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <- partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Err.Error())
		}
	}
}