package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func strToIntArr(s string) []int {
	strs := strings.Split(s, "")
	res := make([]int, len(strs))
	for i := range res {
		res[i], _ = strconv.Atoi(strs[i])
	}
	return res
}

func minOfArr(a []int) int {
	min := a[0]
	for i := range a {
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}

func checkAllSame(a []int) bool {
	c := make(map[int]bool)
	for i := range a {
		c[a[i]] = true
	}

	return len(c) == 1
}

func getScore(a []int) int {
	sum := len(a)
	for i := range a {
		sum += a[i]
	}
	return sum
}

func main() {
	var floor [][]int

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day9/input.txt")
	if err != nil {
		file, _ = os.Open(path + "/input.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tempFloor := strToIntArr(scanner.Text())

		floor = append(floor, tempFloor)
	}

	var lowPoints []int
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			var n []int

			n = append(n, floor[i][j])
			if i != 0 {
				n = append(n, floor[i-1][j])
			}
			if i != len(floor)-1 {
				n = append(n, floor[i+1][j])
			}
			if j != 0 {
				n = append(n, floor[i][j-1])
			}
			if j != len(floor[i])-1 {
				n = append(n, floor[i][j+1])
			}

			if minOfArr(n) == floor[i][j] && !checkAllSame(n) {
				lowPoints = append(lowPoints, minOfArr(n))
			}
		}
	}

	fmt.Println(getScore(lowPoints))
}
