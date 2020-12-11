package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)


func intersection(s1 string, s2 string) string {
	var counter = map[rune]bool {}
	for _, c1 := range s1 {
		for _, c2 := range s2 {
			if c1 == c2 {
				counter[c1] = true
			}
		}
	}
	result := ""
	for i, _ := range counter {
		result += string(i)
	}
	return(result)
}


func countEveryone(group string) int {
	splits := strings.Split(group, " ")
	all := splits[0]
	if len(splits) == 1 {
		return(len(group))
	}
	for i := 1; i < len(splits); i++ {
		all = intersection(all, splits[i])
	}
	return(len(all))
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
				current_answers = current_answers + " " + line
			} else {
				current_answers = line
			}
		} else {
			result += countEveryone(current_answers)
			current_answers = ""
		}
	}
	if current_answers != "" {
		result += countEveryone(current_answers)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(result)
}

