package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
	"time"
)

type IOffsetManagerRepo interface {
	ResetPartitionOffset(offset int64)
	GetPartitionNextOffset() (offset int64, err error)
	Close()
}

type OffsetRepo struct {
	manager   sarama.OffsetManager
	topic     string
	partition int32
}

func NewOffsetManger(cfg types.OffsetConfig) (*OffsetRepo, error) {
	repo := &OffsetRepo{}
	cfgCli := sarama.NewConfig()
	cfgCli.Consumer.Offsets.CommitInterval = 1 * time.Second
	cfgCli.Version = sarama.V2_1_0_0
	if cfg.Password != "" && cfg.Username != "" {
		cfgCli.Net.SASL.Enable = true
		cfgCli.Net.SASL.User = cfg.Username
		cfgCli.Net.SASL.Password = cfg.Password
	}
	cli, err := sarama.NewClient(cfg.Address, cfgCli)
	if err != nil {
		return nil, err
	}
	offsetManager, err := sarama.NewOffsetManagerFromClient(cfg.GroupID, cli)
	if err != nil {
		fmt.Println("offsetManager: " + err.Error())
		return nil, err
	}
	repo.manager = offsetManager
	repo.partition = cfg.Partition
	repo.topic = cfg.Topic
	return repo, nil

}
func (repo *OffsetRepo) Close() {
	repo.manager.Close()
}

func (repo *OffsetRepo) ResetPartitionOffset(offset int64) {
	partitionOffsetManager, err := repo.manager.ManagePartition(repo.topic, repo.partition)
	defer partitionOffsetManager.Close()
	if err != nil {
		fmt.Println("partitionOffsetManager: " + err.Error())
		return
	}
	partitionOffsetManager.ResetOffset(offset, "")
}

func (repo *OffsetRepo) GetPartitionNextOffset() (int64, error) {
	partitionOffsetManager, err := repo.manager.ManagePartition(repo.topic, repo.partition)
	defer partitionOffsetManager.Close()
	if err != nil {
		return -1, err
	}
	offset, _ := partitionOffsetManager.NextOffset()
	return offset, nil
}
