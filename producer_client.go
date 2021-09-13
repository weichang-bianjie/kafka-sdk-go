package kafka_sdk_go

import (
	"github.com/weichang-bianjie/kafka-sdk-go/kafka"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
)

type ProducerClient interface {
	Produce(topic string, data []byte, partition int32) error
	Close()
}

type clientP struct {
	producer kafka.Producer
}

func (c clientP) Produce(topic string, data []byte, partition int32) error {
	return c.producer.Produce(topic, data, partition)
}

func (c clientP) Close() {
	c.producer.Close()
}

func NewProducerClient(groupCfg types.Config) ProducerClient {
	producer, err := kafka.NewProducerRepo(groupCfg)
	if err != nil {
		panic(err)
	}
	return &clientP{
		producer: producer,
	}
}
