package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)


var seatMap [][]rune
const empty rune = 'L'
const occupied rune = '#'
const floor rune = '.'


func countNeighbours(x int, y int, seatMap [][]rune) int {
	maxX := len(seatMap) - 1
	maxY := len(seatMap[0]) - 1
	result := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x + i > maxX || x + i < 0 || y + j > maxY || y + j < 0 {
				continue
			}
			if seatMap[x+i][y+j] == occupied {
				result++
			}
		}
	}
	return result
}


func copySlice(source [][]rune) [][]rune {
	target := make([][]rune, 0)
	for i := 0; i < len(source); i++ {
		line := make([]rune, len(source[0]))
		target = append(target, line)
	}
	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[0]); j++ {
			target[i][j] = source[i][j]
		}
	}
	return target
}



func equalSlices(s1, s2 [][]rune) bool {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[0]); j++ {
			if s1[i][j] != s2[i][j] {
				return false
			}
		}
	}
	return true
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	seatMap := make([][]rune, 0)
	seatMapPrevious := make([][]rune, 0)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		runes := make([]rune, len(line))
		for i, c := range line {
			runes[i] = c
		}
		seatMap = append(seatMap, runes)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	x := len(seatMap)
	y := len(seatMap[0])
	neighbourCount := make([][]int, x)
	for i := 0; i < x; i++ {
		neighbourCount[i] = make([]int, y)
	}
	for true {
		seatMapPrevious = copySlice(seatMap)
		// count neighbours
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				neighbourCount[i][j] = countNeighbours(i, j, seatMap)
			}
		}
		// apply rules
		for i := 0; i < x; i++ {
			for j := 0; j < y; j++ {
				if seatMap[i][j] == empty && neighbourCount[i][j] == 0 {
					seatMap[i][j] = occupied
				} else if seatMap[i][j] == occupied && neighbourCount[i][j] >= 4 {
					seatMap[i][j] = empty
				}
			}
		}
		if equalSlices(seatMap, seatMapPrevious) {
			break
		}
	}

	result := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if seatMap[i][j] == occupied {
				result++
			}
		}
	}
	fmt.Println(result)
}

