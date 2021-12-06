package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func absInt(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var vents [1000][1000]int
	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		startStr := strings.Split(line[0], ",")
		finishStr := strings.Split(line[1], ",")
		startX, _ := strconv.Atoi(startStr[0])
		startY, _ := strconv.Atoi(startStr[1])
		finishX, _ := strconv.Atoi(finishStr[0])
		finishY, _ := strconv.Atoi(finishStr[1])

		if startX == finishX {
			c := absInt(finishY-startY) + min(startY, finishY)
			for i := min(startY, finishY); i <= c; i++ {
				vents[i][startX] += 1
				if vents[i][startX] == 2 {
					score++
				}
			}
		} else if startY == finishY {
			c := absInt(finishX-startX) + min(startX, finishX)
			for i := min(startX, finishX); i <= c; i++ {
				vents[startY][i] += 1
				if vents[startY][i] == 2 {
					score++
				}
			}
		} else {
			if startX < finishX && startY < finishY {
				for i := 0; i <= absInt(finishX-startX); i++ {
					vents[startY+i][startX+i]++
					if vents[startY+i][startX+i] == 2 {
						score++
					}
				}
			} else if startX < finishX && startY > finishY {
				for i := 0; i <= absInt(finishX-startX); i++ {
					vents[startY-i][startX+i]++
					if vents[startY-i][startX+i] == 2 {
						score++
					}
				}
			} else if startX > finishX && startY < finishY {
				for i := 0; i <= absInt(finishX-startX); i++ {
					vents[startY+i][startX-i]++
					if vents[startY+i][startX-i] == 2 {
						score++
					}
				}
			} else if startX > finishX && startY > finishY {
				for i := 0; i <= absInt(finishX-startX); i++ {
					vents[startY-i][startX-i]++
					if vents[startY-i][startX-i] == 2 {
						score++
					}
				}
			}
		}
	}
	fmt.Println(score)
}
