package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func stackAdd(s []string, str string) []string {
	return append(s, str)
}

func stackPop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}

	ret := s[len(s)-1]
	return ret, s[:len(s)-1]
}

func checkLine(line string) (int, []string) {
	var stack []string
	var x string
	c := strings.Split(line, "")

	for _, k := range c {
		if k == "[" || k == "(" || k == "{" || k == "<" {
			stack = stackAdd(stack, k)
		} else {
			x, stack = stackPop(stack)
			if k == "]" && x != "[" {
				return 57, nil
			} else if k == ")" && x != "(" {
				return 3, nil
			} else if k == "}" && x != "{" {
				return 1197, nil
			} else if k == ">" && x != "<" {
				return 25137, nil
			}
		}
	}

	return 0, stack
}

func finishLine(line []string) int {
	score := 0

	for i := len(line) - 1; i >= 0; i-- {
		score *= 5
		if line[i] == "(" {
			score += 1
		} else if line[i] == "[" {
			score += 2
		} else if line[i] == "{" {
			score += 3
		} else if line[i] == "<" {
			score += 4
		}
	}
	return score
}

func main() {
	score := 0
	var part2 [][]string

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day10/input.txt")
	if err != nil {
		file, _ = os.Open(path + "/test.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		tempScore, tempStack := checkLine(temp)
		if tempScore == 0 {
			part2 = append(part2, tempStack)
		}

		score += tempScore
	}

	var score2 []int
	for _, k := range part2 {
		score2 = append(score2, finishLine(k))
	}

	sort.Ints(score2)

	fmt.Println(score)
	fmt.Println(score2[(len(score2) / 2)])
}
