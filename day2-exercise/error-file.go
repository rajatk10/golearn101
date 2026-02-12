package main

import (
	"fmt"
	"os"
)

func OpenNonExistentFile() {
	file, err := os.Open("Nonexistent.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	fmt.Println("Opened file")
}

func main() {
	OpenNonExistentFile()
}
