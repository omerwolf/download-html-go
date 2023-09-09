package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("ListOfAsciiSiteUrl.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	fmt.Println(string(content))
}
