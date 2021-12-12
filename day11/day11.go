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

func flashOctopus(arr [][]int, j, k int) ([][]int, int) {
	if j < 0 || k < 0 || j >= len(arr) || k >= len(arr[0]) {
		return arr, 0
	}

	if arr[j][k] == -1 {
		return arr, 0
	}

	numFlashes := 0
	tempFlashes := 0
	arr[j][k]++
	if arr[j][k] > 9 {
		arr[j][k] = -1
		numFlashes = 1
		arr, tempFlashes = flashOctopus(arr, j+1, k)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j+1, k+1)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j+1, k-1)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j-1, k)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j-1, k+1)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j-1, k-1)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j, k+1)
		numFlashes += tempFlashes
		arr, tempFlashes = flashOctopus(arr, j, k-1)
		numFlashes += tempFlashes
	}

	return arr, numFlashes
}

func main() {
	var octopuses [][]int

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day11/input.txt")
	if err != nil {
		file, _ = os.Open(path + "/test.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		octopuses = append(octopuses, strToIntArr(scanner.Text()))
	}

	flashes := 0
	totalFlashes := 0
	for i := 0; i < 5000; i++ {
		roundFlashes := 0
		for j := 0; j < len(octopuses); j++ {
			for k := 0; k < len(octopuses[0]); k++ {
				octopuses, flashes = flashOctopus(octopuses, j, k)
				totalFlashes += flashes
				roundFlashes += flashes
			}
		}
		if roundFlashes == 100 {
			fmt.Println(i + 1)
			return
		}
		for j := 0; j < len(octopuses); j++ {
			for k := 0; k < len(octopuses[0]); k++ {
				if octopuses[j][k] == -1 {
					octopuses[j][k] = 0
				}
			}
		}
		if i == 99 {
			fmt.Println(totalFlashes)
		}
	}

}
