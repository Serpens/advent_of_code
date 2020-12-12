package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)


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

	pointer := 0
	visited := make([]bool, len(code))
	acc := 0
	var instruction, command string
	var increment int

	for true {
		if visited[pointer] {
			fmt.Println(acc)
			return
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
}

