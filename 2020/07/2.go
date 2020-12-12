package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)


type bagNumber struct {
	number int
	colour string
}


func parseRule(rule string) (string, []bagNumber) {
	splits := strings.Split(rule, " bags contain ")
	outer := splits[0]
	splits = strings.Split(splits[1], ", ")
	inner := make([]bagNumber, 0)
	for _, s := range splits {
		words := strings.Split(s, " ")
		number, _ := strconv.Atoi(words[0])
		currentInner := bagNumber{
			number: number,
			colour: words[1] + " " + words[2],
		}
		inner = append(inner, currentInner)
	}
	return outer, inner
}


func countBagsIn(outer string, rules map[string][]bagNumber) int {
	result := 1
	for _, rule := range rules[outer] {
		if rule.number > 0 {
			result += rule.number * countBagsIn(rule.colour, rules)
		}
	}
	return result
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	allRules := make(map[string][]bagNumber)
	for scanner.Scan() {
		line = scanner.Text()
		k, v := parseRule(line)
		allRules[k] = v
	}

	fmt.Println(countBagsIn("shiny gold", allRules) - 1)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

