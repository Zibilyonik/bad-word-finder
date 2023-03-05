package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var file_name, sentence string
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
		var final string
		fmt.Scan(&sentence)
		arr := strings.Split(sentence, " ")
		if sentence == "exit" {
			fmt.Println("Bye!")
			return
		}
		for _, word := range arr {
			lowered := strings.ToLower(word)
			if _, ok := word_map[lowered]; ok {
				stars := strings.Repeat("*", len(word))
				final += stars + " "
			} else {
				final += word + " "
			}
		}
		fmt.Println(strings.TrimSuffix(final, " "))

	}
}
