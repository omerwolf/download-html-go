package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	inputFile = "ListOfAsciiSiteUrl.txt"
	outputDir = "html_files"
)

func main() {
	urlsFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	createDir(outputDir)

	scanner := bufio.NewScanner(urlsFile)
	for scanner.Scan() {
		url := scanner.Text()
		downloadHTML(url, outputDir)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = urlsFile.Close()
	if err != nil {
		return
	}
}

func createDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}
}

func downloadHTML(url, outputDir string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", url, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	filename := filepath.Join(outputDir, filepath.Base(url))
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating output file %s: %s\n", filename, err)
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			return
		}
	}(outputFile)

	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Printf("Error writing to output file %s: %s\n", filename, err)
		return
	}

	fmt.Printf("Downloaded %s to %s\n", url, filename)
}
