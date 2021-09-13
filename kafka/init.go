package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
)

func InitConsumer(groupCfg types.Config, offsetCfg types.OffsetConfig) Consumer {
	offsetRepo, err := NewOffsetManger(offsetCfg)
	defer func() {
		if offsetRepo != nil {
			offsetRepo.Close()
		}
	}()
	if err != nil {
		panic(err.Error())
	}

	offset, err := offsetRepo.GetPartitionNextOffset()
	if err != nil {
		panic(err.Error())
	}
	// 设置读取位置
	if offsetCfg.ResetOffset > 0 {
		offsetRepo.ResetPartitionOffset(offsetCfg.ResetOffset)
	}
	// 判断应该从头还是从最新（当group为新建，从最旧）
	if offset == -1 {
		groupCfg.Offset = sarama.OffsetOldest
	} else {
		groupCfg.Offset = sarama.OffsetNewest
	}

	consumer, err := NewConsumerRepo(groupCfg)
	if err != nil {
		panic(err.Error())
	}
	return consumer
}

//func Start(consumer *consumerRepo) {
//	go consumer.Consume(handler)
//}
