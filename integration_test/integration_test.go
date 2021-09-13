package integration_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/weichang-bianjie/kafka-sdk-go"
	"github.com/weichang-bianjie/kafka-sdk-go/types"
	"testing"
)

type IntegrationTestSuite struct {
	suite.Suite
	kafka_sdk_go.ConsumerClient
	kafka_sdk_go.ProducerClient
}

func TestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	topic := []string{"test"}
	groupId := "group-2"
	address := []string{"127.0.0.1:9092"}
	gcfg := types.Config{}
	gcfg.Address = address
	gcfg.Topics = topic
	gcfg.GroupID = groupId
	gcfg.Username = ""
	gcfg.Password = ""

	offsetCfg := types.OffsetConfig{}
	offsetCfg.Address = address
	offsetCfg.Topic = topic[0]
	offsetCfg.GroupID = groupId
	offsetCfg.Username = ""
	offsetCfg.Password = ""
	offsetCfg.Partition = 0
	offsetCfg.ResetOffset = 2
	s.ConsumerClient = kafka_sdk_go.NewConsumerClient(gcfg, offsetCfg)
	s.ProducerClient = kafka_sdk_go.NewProducerClient(gcfg)
	//defer s.ProducerClient.Close()
	//defer s.ConsumerClient.Close()
}
