package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(`file.txt`)
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()
}
