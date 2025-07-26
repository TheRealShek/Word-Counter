package main

import (
	"fmt"
	"log"
	"sync"
	"wordpipeline/filescanner"
	"wordpipeline/processor"
	"wordpipeline/reader"
)

func main() {
	root := "TestFiles"
	maxProcessors := 5

	txtFiles, err := filescanner.FindTxtFiles(root)
	if err != nil {
		log.Fatalf("Error walking directory: %v", err)
	}

	if len(txtFiles) == 0 {
		fmt.Println("No .txt files found to process.")
		return
	}

	fmt.Printf("Found %d .txt files.\n", len(txtFiles))
	for _, file := range txtFiles {
		fmt.Println(file)
	}

	wordChan := make(chan string, 100)

	var fileReadWg sync.WaitGroup
	var wordProcessWg sync.WaitGroup

	for _, file := range txtFiles {
		fileReadWg.Add(1)
		go reader.ReadFileWords(file, wordChan, &fileReadWg)
	}

	var TotalWordsPerFile uint64
	for i := 1; i <= maxProcessors; i++ {
		wordProcessWg.Add(1)
		go processor.WordProcessor(i, wordChan, &wordProcessWg, &TotalWordsPerFile)
	}

	go func() {
		fileReadWg.Wait()
		close(wordChan)
		fmt.Println("All file readers finished and word channel closed.")
	}()

	wordProcessWg.Wait()
	fmt.Printf("All word processors finished. Total words processed: %d\n", TotalWordsPerFile)
}
