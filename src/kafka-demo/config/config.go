package config

import "github.com/Shopify/sarama"

var config *sarama.Config

func init() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewManualPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_10_0_0
	config.Consumer.Return.Errors = true
}

func Config() *sarama.Config {
	return config
}
