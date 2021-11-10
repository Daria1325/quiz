package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fileName := flag.String("file", "default.json", "file with questions in JSON")
	flag.Parse()
	_ = fileName

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Faild to open JSON file: %s", *fileName)
		os.Exit(1)
	}
	defer file.Close()
	_ = file
}
