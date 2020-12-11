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
var pos1, pos2 int
var letter rune
var password string

func is_valid(pos1 int, pos2 int, letter rune, password string) bool {
	c1 := rune(password[pos1 - 1])
	c2 := rune(password[pos2 - 1])
	return (c1 == letter || c2 == letter) && c1 != c2
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
		pos1, converr = strconv.Atoi(splits[0])
		pos2, converr = strconv.Atoi(splits[1])
		if converr != nil {
			log.Println(converr)
		}
		if is_valid(pos1, pos2, letter, password) {
			valid_counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println(valid_counter)
}

