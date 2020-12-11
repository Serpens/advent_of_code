package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"regexp"
)


func is_in_array(value string, array [7]string) bool {
	for _, v := range array {
		if value == v {
			return true
		}
	}
	return false
}


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
		field_value := strings.Split(part, ":")[1]
		_, exists := required[field]
		if exists {
			required[field] = true
		}
		switch (field) {
			case "byr":
				byr, err := strconv.Atoi(field_value)
				if byr < 1920 || byr > 2002 {
					return false
				}
				if err != nil {
					return false
				}
			case "iyr":
				iyr, err := strconv.Atoi(field_value)
				if iyr < 2010 || iyr > 2020 {
					return false
				}
				if err != nil {
					return false
				}
			case "eyr":
				eyr, err := strconv.Atoi(field_value)
				if eyr < 2020 || eyr > 2030 {
					return false
				}
				if err != nil {
					return false
				}
			case "hgt":
				if field_value[len(field_value)-2:len(field_value)] == "cm" {
					hgt, err := strconv.Atoi(field_value[0:3])
					if hgt < 150 || hgt > 193 {
						return false
					}
					if err != nil {
						return false
					}
				} else if field_value[len(field_value)-2:len(field_value)] == "in" {
					hgt, err := strconv.Atoi(field_value[0:2])
					if hgt < 59 || hgt > 76 {
						return false
					}
					if err != nil {
						return false
					}
				} else {
					return false
				}
			case "hcl":
				matched, err := regexp.MatchString("^#[0-9a-f]{6}$", field_value)
				if !matched || err != nil {
					return false
				}
			case "ecl":
				var allowed = [...]string {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				if !is_in_array(field_value, allowed) {
					return false
				}
			case "pid":
				matched, err := regexp.MatchString("^[0-9]{9}$", field_value)
				if !matched || err != nil {
					return false
				}
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
				//fmt.Println(current_passport)
				//fmt.Println("")
			} else {
				//fmt.Println(current_passport)
			}
			current_passport = ""
		}
	}
	if current_passport != "" {
		if is_valid(current_passport) {
			valid_counter++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(valid_counter)

}

