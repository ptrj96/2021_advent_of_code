package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CheckCards(val int, cards *[][]map[int]bool) int {
	for i := 0; i < len(*cards); i++ {
		for j := 0; j < len((*cards)[i]); j++ {
			if (*cards)[i][j][val] {
				delete((*cards)[i][j], val)
			}
			if len((*cards)[i][j]) == 0 {
				if len((*cards)) == 1 {
					return ScoreCard(val, (*cards)[i])
				} else {
					*cards = append((*cards)[:i], (*cards)[i+1:]...)
					i -= 1
					break
				}
			}
		}
	}
	return 0
}

func ScoreCard(val int, card []map[int]bool) int {
	remainder := make(map[int]bool)

	for i := 0; i < len(card); i++ {
		for k := range card[i] {
			if k != val {
				remainder[k] = true
			}
		}
	}

	sum := 0
	for k := range remainder {
		sum += k
	}
	return sum * val
}

func main() {
	vals := []int{87, 12, 53, 23, 31, 70, 37, 79, 95, 16, 72, 9, 98, 92, 5, 74, 17, 60, 96, 80, 75, 11, 73, 33, 3, 84, 81, 2, 97, 93, 59, 13, 77, 52, 69, 83, 51, 64, 48, 82, 7, 49, 20, 8, 36, 66, 19, 0, 99, 41, 91, 78, 42, 40, 62, 63, 57, 39, 55, 47, 29, 50, 58, 34, 27, 43, 30, 35, 22, 28, 4, 14, 26, 32, 10, 88, 46, 65, 90, 76, 38, 6, 71, 67, 44, 68, 86, 25, 21, 24, 56, 94, 18, 89, 61, 15, 1, 45, 54, 85}
	// vals := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards [][]map[int]bool = [][]map[int]bool{{make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool)}}

	i := 0
	boardNum := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 0 {
			cards = append(cards, []map[int]bool{make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool), make(map[int]bool)})
			boardNum++
			i = 0
			continue
		}
		for j := 0; j < len(line); j++ {
			temp, _ := strconv.Atoi(line[j])
			cards[boardNum][j][temp] = true
			cards[boardNum][i+5][temp] = true
		}
		i++
	}

	for _, val := range vals {
		score := CheckCards(val, &cards)
		if score > 0 {
			fmt.Printf("Winning value is %d\n", score)
			fmt.Printf("%d was drawn last\n", val)
			return
		}
	}
}
