package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

