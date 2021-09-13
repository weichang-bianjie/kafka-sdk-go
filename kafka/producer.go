package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
	"time"
)

type Producer interface {
	Produce(topic string, data []byte, partition int32) error
	Close()
}

type producerRepo struct {
	producer sarama.SyncProducer
}

func NewProducerRepo(cfg types.Config) (Producer, error) {
	config := sarama.NewConfig()
	// kafka 认证 start
	if cfg.Password != "" && cfg.Username != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = cfg.Username
		config.Net.SASL.Password = cfg.Password
	}
	// kafka 认证 end
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(cfg.Address, config)
	if err != nil {
		fmt.Println("NewSyncProducer err, message= \n", err)
		return nil, err
	}
	pRepo := &producerRepo{}
	pRepo.producer = p
	return pRepo, nil
}

func (pRepo *producerRepo) Produce(topic string, data []byte, partition int32) (err error) {
	pMsg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(data),
		Partition: partition,
	}
	part, offset, err := pRepo.producer.SendMessage(pMsg)
	if err != nil {
		return
	}
	msg := fmt.Sprintf("send succcess: partition=%d, offset=%d \n", part, offset)
	fmt.Println(msg)
	return
}

func (pRepo *producerRepo) Close() {
	pRepo.producer.Close()
}
