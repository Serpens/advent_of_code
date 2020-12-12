package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)


func parseRule(rule string) map[string][]string {
	result := make(map[string][]string)
	splits := strings.Split(rule, " bags contain ")
	outer := splits[0]
	splits = strings.Split(splits[1], ", ")
	for _, inner := range splits {
		words := strings.Split(inner, " ")
		colour := words[1] + " " + words[2]
		result[colour] = append(result[colour], outer)
	}
	return(result)
}


func stringInSlice(str string, slice []string) bool {
	for _, s := range(slice) {
		if str == s {
			return true
		}
	}
	return false
}


func mergeSlices(s1 []string, s2 []string) []string {
	var result []string
	for _, v := range s1 {
		result = append(result, v)
	}
	for _, v1 := range s2 {
		if !stringInSlice(v1, result) {
			result = append(result, v1)
		}
	}
	return result
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	allRules := make(map[string][]string)
	for scanner.Scan() {
		line = scanner.Text()
		currentRules := parseRule(line)
		for k, v := range currentRules {
			allRules[k] = mergeSlices(allRules[k], v)
		}
	}

	currentLayer := allRules["shiny gold"]
	allColours := make([]string, 0)
	for len(currentLayer) > 0 {
		for _, bag := range currentLayer {
			if !stringInSlice(bag, allColours) {
				allColours = append(allColours, bag)
			}
		}
		nextLayer := make([]string, 0)
		for _, bag := range currentLayer {
			nextLayer = mergeSlices(nextLayer, allRules[bag])
		}
		currentLayer = nextLayer
	}
	fmt.Println(len(allColours))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}

