package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)


func is_valid(passport string) bool {
	var required = map[string]bool {
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}

	parts := strings.Split(passport, " ")

	for _, part := range parts {
		field := strings.Split(part, ":")[0]
		_, exists := required[field]
		if exists {
			required[field] = true
		}
	}

	for _, i := range required {
		if !i {
			return false
		}
	}
	return true
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	current_passport := ""
	valid_counter := 0
	for scanner.Scan() {
		line = scanner.Text()
		if line != "" {
			if current_passport != "" {
				current_passport = current_passport + " " + line
			} else {
				current_passport = line
			}
		} else {
			if is_valid(current_passport) {
				valid_counter++
			}
			current_passport = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(valid_counter)

}

