package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
	"os"
)

type Consumer interface {
	Consume(cb func(value []byte) error)
	ConsumeMessage(cb func(value *sarama.ConsumerMessage) error)
	Close()
}

type consumerRepo struct {
	consumer *cluster.Consumer
}

func NewConsumerRepo(cfg types.Config) (Consumer, error) {
	cRepo := &consumerRepo{}
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	if cfg.Password != "" && cfg.Username != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.User = cfg.Username
		config.Net.SASL.Password = cfg.Password
	}
	if cfg.Offset == 0 {
		cfg.Offset = sarama.OffsetNewest
	}
	config.Consumer.Offsets.Initial = cfg.Offset
	consumer, err := cluster.NewConsumer(cfg.Address, cfg.GroupID, cfg.Topics, config)
	if err != nil {
		return nil, err
	}

	cRepo.consumer = consumer
	return cRepo, nil
}

func (cRepo *consumerRepo) Close() {
	cRepo.consumer.Close()
}

func (cRepo *consumerRepo) Consume(callback func(value []byte) error) {
	// consume errors
	go func() {
		for err := range cRepo.consumer.Errors() {
			fmt.Println(err.Error())
		}
	}()
	// consume notifications
	go func() {
		for ntf := range cRepo.consumer.Notifications() {
			fmt.Println(fmt.Sprint("Rebalanced: ", ntf))
		}
	}()

	for {
		select {
		case msg, ok := <-cRepo.consumer.Messages():
			if ok {
				//fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				err := callback(msg.Value)
				fmt.Println(fmt.Sprint("current offset is: ", msg.Offset))
				if err != nil {
					fmt.Println(err.Error())
				} else {
					cRepo.consumer.MarkOffset(msg, "") // mark message as processed
				}

			}
		}
	}
}

func (cRepo *consumerRepo) ConsumeMessage(cb func(value *sarama.ConsumerMessage) error) {
	// consume errors
	go func() {
		for err := range cRepo.consumer.Errors() {
			fmt.Println(err.Error())
		}
	}()
	// consume notifications
	go func() {
		for ntf := range cRepo.consumer.Notifications() {
			fmt.Println(fmt.Sprint("Rebalanced: ", ntf))
		}
	}()

	for {
		select {
		case msg, ok := <-cRepo.consumer.Messages():
			if ok {

				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				err := cb(msg)
				if err != nil {
					return
				} else {
					cRepo.consumer.MarkOffset(msg, "") // mark message as processed
				}
			}
		}
	}
}
