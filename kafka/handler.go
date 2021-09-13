package kafka

import (
	"fmt"
	"os"
)

func handler(data []byte) error {

	if len(data) == 0 {
		return nil
	}
	fmt.Fprintf(os.Stdout, "%s\n", data)
	return nil
}
