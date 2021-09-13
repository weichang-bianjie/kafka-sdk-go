package integration_test

import (
	"github.com/stretchr/testify/require"
)

func (s *IntegrationTestSuite) TestProduce() {
	err := s.ProducerClient.Produce("test", []byte("i am a programmer"), 0)
	require.NoError(s.T(), err)
}
