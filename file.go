package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteSting("123456")
	file.Close()
}