package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"kafka-demo/config"
	"time"
)

func Test_producer() {
	config := config.Config()

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"},config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}
	defer producer.AsyncClose()



	var flag bool
	var partition int32 =  0
	var count = 0

	for {
		if flag{
			partition = -1
		} else {
			partition = 0
		}
		flag = !flag
		msg := &sarama.ProducerMessage{
			Topic: "kafka-go-test1",
			Key: sarama.StringEncoder("go_test"),
			Partition: partition,
			Value: sarama.ByteEncoder(fmt.Sprintf("%d", count)),
		}
		fmt.Println(partition)
		count ++
		producer.Input() <- msg

		select {
		case suc := <- producer.Successes():
			fmt.Printf("offset: %d, timestamp:%s", suc.Offset, suc.Timestamp.String())
		case fail := <- producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		default:
			fmt.Printf("Produced message default\n")
		}
		time.Sleep(1 *time.Second)
	}
}
