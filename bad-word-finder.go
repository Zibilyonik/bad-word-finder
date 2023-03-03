package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var file_name string
	fmt.Scan(&file_name)
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
