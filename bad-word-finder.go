package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var file_name, word string
	fmt.Scan(&file_name)
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Scan(&word)
	scanner := bufio.NewScanner(file)

	// first step of the project
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == strings.ToLower(word) {
			fmt.Println("true")
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("false")
}
