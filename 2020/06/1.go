package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)


func countLetters(group string) int{
	var counter = map[rune]bool {}
	for _, c := range group {
		counter[c] = true
	}
	return(len(counter))
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	current_answers := ""
	result := 0
	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			if current_answers != "" {
				current_answers = current_answers + "" + line
			} else {
				current_answers = line
			}
		} else {
			result += countLetters(current_answers)
			current_answers = ""
		}
	}
	if current_answers != "" {
		result += countLetters(current_answers)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(result)
}

