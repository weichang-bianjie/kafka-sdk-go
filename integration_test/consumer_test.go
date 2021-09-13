package integration_test

import (
	"fmt"
	"os"
)

func (s *IntegrationTestSuite) TestStart() {
	s.ConsumerClient.Start(func(data []byte) error {
		if len(data) == 0 {
			return nil
		}
		fmt.Fprintf(os.Stdout, "%s\n", data)
		return nil
	})
}
