package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)


func isTerminating(code []string) (bool, int) {
	pointer := 0
	visited := make([]bool, len(code))
	acc := 0
	var instruction, command string
	var increment int

	for true {
		if pointer == len(code) {
			return true, acc
		}
		if visited[pointer] {
			return false, acc
		}
		visited[pointer] = true
		instruction = code[pointer]
		splits := strings.Split(instruction, " ")
		command = splits[0]
		increment, _ = strconv.Atoi(splits[1])
		switch (command) {
			case "nop":
				pointer++
			case "acc":
				acc += increment
				pointer++
			case "jmp":
				pointer += increment
		}
	}
	return true, acc
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	code := make([]string, 0)
	for scanner.Scan() {
		line = scanner.Text()
		code = append(code, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for i := 0; i < len(code); i++ {
		if strings.HasPrefix(code[i], "nop") {
			code[i] = strings.Replace(code[i], "nop", "jmp", 1)
			notLoop, acc := isTerminating(code)
			if notLoop {
				fmt.Println(acc)
				return
			}
			code[i] = strings.Replace(code[i], "jmp", "nop", 1)
		} else if strings.HasPrefix(code[i], "jmp") {
			code[i] = strings.Replace(code[i], "jmp", "nop", 1)
			notLoop, acc := isTerminating(code)
			if notLoop {
				fmt.Println(acc)
				return
			}
			code[i] = strings.Replace(code[i], "nop", "jmp", 1)
		}
	}
}

