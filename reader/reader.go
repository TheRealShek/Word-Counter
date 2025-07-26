package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func ReadFileWords(filepath string, wordChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filepath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)

		word = strings.TrimFunc(word, func(r rune) bool {
			return !('a' <= r && r <= 'z' || '0' <= r && r <= '9')
		})

		if word != "" {
			wordChan <- word
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file %s: %v", filepath, err)
	}
}
