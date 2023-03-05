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
	word_map := map[string]struct{}{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word_map[scanner.Text()] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for true {
		fmt.Scan(&word)
		lowered := strings.ToLower(word)
		if lowered == "exit" {
			fmt.Println("Bye!")
			return
		}
		if _, ok := word_map[lowered]; ok {
			stars := strings.Repeat("*", len(word))
			fmt.Println(stars)
		} else {
			fmt.Println(word)
		}
	}
}
