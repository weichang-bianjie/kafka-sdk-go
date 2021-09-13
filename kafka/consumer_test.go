package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
	"testing"
)

func TestConsumerRepo_Consume(t *testing.T) {
	topic := []string{"test"}
	address := []string{"127.0.0.1:9092"}

	var (
		offsetCfg types.OffsetConfig
	)

	offsetCfg.Address = address
	offsetCfg.Topic = "test"
	offsetCfg.GroupID = "group-2"
	offsetCfg.Username = ""
	offsetCfg.Password = ""
	offsetCfg.Partition = 0

	offsetRepo, err := NewOffsetManger(offsetCfg)
	defer offsetRepo.Close()
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	g1Cfg := types.Config{}
	g1Cfg.Address = address
	g1Cfg.Topics = topic
	g1Cfg.GroupID = "group-2"
	g1Cfg.Username = ""
	g1Cfg.Password = ""

	offset, err := offsetRepo.GetPartitionNextOffset()

	// 判断应该从头还是从最新（当group为新建，从最旧）
	if offset == -1 {
		g1Cfg.Offset = sarama.OffsetOldest
	} else {
		g1Cfg.Offset = sarama.OffsetNewest
	}

	cRepoG1, err := NewConsumerRepo(g1Cfg)
	if err != nil {
		return
	}
	cRepoG1.Consume(handler)
}
