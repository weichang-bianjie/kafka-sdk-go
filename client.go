package kafka_sdk_go

import (
	"github.com/weichang-bianjie/kafka-sdk-go/kafka"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
)

type ConsumerClient interface {
	Start(handler Handler)
	Close()
}

type Handler func(data []byte) error
type client struct {
	consumer kafka.Consumer
}

func (c client) Close() {
	c.consumer.Close()
}

func (c client) Start(handler Handler) {
	c.consumer.Consume(handler)
}

func NewConsumerClient(groupCfg types.Config, offsetCfg types.OffsetConfig) ConsumerClient {
	consumer := kafka.InitConsumer(groupCfg, offsetCfg)
	return &client{
		consumer: consumer,
	}
}
