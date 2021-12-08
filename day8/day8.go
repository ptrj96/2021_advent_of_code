package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func checkString(dict map[string]bool, input string) bool {
	return false
}

func sortStr(str string) string {
	uStr := strings.Split(str, "")
	sort.Strings(uStr)
	return strings.Join(uStr, "")
}

func getTop(one, seven string) string {
	for i := 0; i < len(one); i++ {
		seven = strings.ReplaceAll(seven, string(one[i]), "")
	}
	return seven
}

func getFour(one, four string) string {
	for i := 0; i < len(one); i++ {
		four = strings.ReplaceAll(four, string(one[i]), "")
	}
	return four
}

func strIn(sub, str string) bool {
	for i := 0; i < len(sub); i++ {
		if !strings.Contains(str, string(sub[i])) {
			return false
		}
	}
	return true
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day8/input.txt")
	if err != nil {
		file, _ = os.Open(path + "/test.txt")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " | ")
		layout := strings.Fields(data[0])

		sort.Slice(layout, func(i, j int) bool {
			return len(layout[i]) < len(layout[j])
		})

		digits := strings.Fields(data[1])
		var out = make(map[string]int)
		var oneStr string
		var four string
		for i := 0; i < len(layout); i++ {
			str := sortStr(layout[i])

			switch len(str) {
			case 2:
				out[str] = 1
				oneStr = str
			case 3:
				out[str] = 7
			case 4:
				out[str] = 4
				four = getFour(oneStr, str)
			case 5:
				//options 2 3 5
				if strIn(oneStr, str) {
					out[str] = 3
				} else if strIn(four, str) {
					out[str] = 5
				} else {
					out[str] = 2
				}
			case 6:
				//options 0 6 9
				if !strIn(four, str) {
					out[str] = 0
				} else if strIn(oneStr, str) {
					out[str] = 9
				} else {
					out[str] = 6
				}
			case 7:
				out[str] = 8
			}
		}
		tempSum := 0
		for i := 0; i < len(digits); i++ {
			tempSum *= 10
			x := sortStr(digits[i])
			tempSum += out[x]
		}
		sum += tempSum
	}

	fmt.Println(sum)
}
