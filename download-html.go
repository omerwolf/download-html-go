package main

import (
	"bufio"
	"fmt"
	"log"
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
		createFile(url)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	urlsFile.Close()
}

func createDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}
}
func createFile(url string) {
	filename := filepath.Join(outputDir, filepath.Base(url))
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating output file %s: %s\n", filename, err)
		return
	}
	defer outputFile.Close()
}
