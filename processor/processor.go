package processor

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func WordProcessor(id int, words <-chan string, wg *sync.WaitGroup, TotalWords *uint64) {
	defer wg.Done()

	for word := range words {
		fmt.Printf("Processor %d received: %s\n", id, word)
		atomic.AddUint64(TotalWords, 1)
	}
}
