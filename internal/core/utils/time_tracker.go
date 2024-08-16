package utils

import (
	"fmt"
	"time"
)

func TimeTrack(operationName string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s took %s\n", operationName, elapsed)
	}
}
