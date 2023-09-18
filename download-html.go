package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	inputFile  = "ListOfAsciiSiteUrl.txt"
	outputDir  = "html_files"
	numWorkers = 5
	bufferSize = 10
)

func main() {
	urlsFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer urlsFile.Close()

	createDir(outputDir)

	urls := make(chan string, bufferSize)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urls {
				downloadHTML(url, outputDir)
			}
		}()
	}

	scanner := bufio.NewScanner(urlsFile)
	for scanner.Scan() {
		url := scanner.Text()
		urls <- url
	}
	close(urls)

	wg.Wait()
}

func createDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
	}
}

func downloadHTML(url, outputDir string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", url, err)
		return
	}
	defer response.Body.Close()

	filename := filepath.Join(outputDir, filepath.Base(url))
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating output file %s: %s\n", filename, err)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Printf("Error writing to output file %s: %s\n", filename, err)
		return
	}

	fmt.Printf("Downloaded %s to %s\n", url, filename)
}
