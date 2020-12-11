package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"log"
)

const length = 1000
var current string
//var str_range, str_letter, str_min, str_max string
var min, max int
var letter rune
var password string

func is_valid(min int, max int, letter rune, password string) bool {
	counter := 0
	for i, char := range password {
		if char == letter {
			counter++
		}
		_ = i
	}
	return counter >= min && counter <= max
}

func main() {
	valid_counter := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		current = scanner.Text()
		var splits = strings.Split(current, " ")
		letter = rune(splits[1][0])
		password = splits[2]
		splits = strings.Split(splits[0], "-")
		var converr error
		min, converr = strconv.Atoi(splits[0])
		max, converr = strconv.Atoi(splits[1])
		if converr != nil {
			log.Println(converr)
		}
		if is_valid(min, max, letter, password) {
			valid_counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(valid_counter)
}

