package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var floor [][]int

type Point struct {
	x int
	y int
}

func strToIntArr(s string) []int {
	strs := strings.Split(s, "")
	res := make([]int, len(strs))
	for i := range res {
		res[i], _ = strconv.Atoi(strs[i])
	}
	return res
}

func findWell(x, y int) (bool, map[Point]bool) {
	tempWell := make(map[Point]bool)
	if x == 0 && y == 0 {
		// Top Left
		if floor[x+1][y] > floor[x][y] && floor[x][y+1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y+1, tempWell)
			tempWell = helpWell(x+1, y, tempWell)

			return true, tempWell
		}
	} else if x == 0 && y == len(floor[0])-1 {
		// Top Right
		if floor[x+1][y] > floor[x][y] && floor[x][y-1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y-1, tempWell)
			tempWell = helpWell(x+1, y, tempWell)

			return true, tempWell
		}
	} else if x == len(floor)-1 && y == 0 {
		// Bottom Left
		if floor[x-1][y] > floor[x][y] && floor[x][y+1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y+1, tempWell)
			tempWell = helpWell(x-1, y, tempWell)

			return true, tempWell
		}
	} else if x == len(floor)-1 && y == len(floor[0])-1 {
		// Bottom Right
		if floor[x-1][y] > floor[x][y] && floor[x][y-1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y-1, tempWell)
			tempWell = helpWell(x-1, y, tempWell)

			return true, tempWell
		}
	} else if x == 0 {
		// Top Row
		if floor[x+1][y] > floor[x][y] && floor[x][y-1] > floor[x][y] && floor[x][y+1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y+1, tempWell)
			tempWell = helpWell(x, y-1, tempWell)
			tempWell = helpWell(x+1, y, tempWell)

			return true, tempWell
		}
	} else if x == len(floor)-1 {
		// Bottom Row
		if floor[x-1][y] > floor[x][y] && floor[x][y-1] > floor[x][y] && floor[x][y+1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y+1, tempWell)
			tempWell = helpWell(x, y-1, tempWell)
			tempWell = helpWell(x-1, y, tempWell)

			return true, tempWell
		}
	} else if y == 0 {
		// Left Collumn
		if floor[x-1][y] > floor[x][y] && floor[x+1][y] > floor[x][y] && floor[x][y+1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y+1, tempWell)
			tempWell = helpWell(x+1, y, tempWell)
			tempWell = helpWell(x-1, y, tempWell)

			return true, tempWell
		}
	} else if y == len(floor[0])-1 {
		// Right Collumn
		if floor[x-1][y] > floor[x][y] && floor[x+1][y] > floor[x][y] && floor[x][y-1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y-1, tempWell)
			tempWell = helpWell(x+1, y, tempWell)
			tempWell = helpWell(x-1, y, tempWell)

			return true, tempWell
		}
	} else {
		// Middle
		if floor[x-1][y] > floor[x][y] && floor[x+1][y] > floor[x][y] && floor[x][y-1] > floor[x][y] && floor[x][y+1] > floor[x][y] {
			tempWell[Point{x: x, y: y}] = true
			tempWell = helpWell(x, y+1, tempWell)
			tempWell = helpWell(x, y-1, tempWell)
			tempWell = helpWell(x+1, y, tempWell)
			tempWell = helpWell(x-1, y, tempWell)

			return true, tempWell
		}
	}
	return false, nil
}

func helpWell(x, y int, w map[Point]bool) map[Point]bool {
	if x < 0 || y < 0 || x >= len(floor) || y >= len(floor[0]) {
		return w
	} else if w[Point{x: x, y: y}] {
		return w
	} else if floor[x][y] == 9 {
		return w
	}

	w[Point{x: x, y: y}] = true

	w = helpWell(x, y+1, w)
	w = helpWell(x, y-1, w)
	w = helpWell(x+1, y, w)
	w = helpWell(x-1, y, w)

	return w
}

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day9/input.txt")
	if err != nil {
		file, _ = os.Open(path + "/test.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tempFloor := strToIntArr(scanner.Text())

		floor = append(floor, tempFloor)
	}

	var wells []map[Point]bool
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			if floor[i][j] != 9 {
				if isWell, well := findWell(i, j); isWell {
					wells = append(wells, well)
				}
			}
		}
	}

	sort.Slice(wells, func(i, j int) bool {
		return len(wells[i]) > len(wells[j])
	})
	total := 1
	for i := 0; i < 3; i++ {
		total *= len(wells[i])
	}

	fmt.Println(total)
}
