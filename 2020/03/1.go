package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var line string

func main() {
	tree_map := make([][31]bool, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		var tree_line_bool [31]bool
		for i, char := range line {
			tree_line_bool[i] = (char == rune('#'))
		}
		tree_map = append(tree_map, tree_line_bool)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	tree_counter := 0
	x := 0
	for y := 0; y < len(tree_map); y++ {
		if tree_map[y][x] {
			tree_counter++
		}
		x = (x + 3) % 31
	}
	fmt.Println(tree_counter)
}

