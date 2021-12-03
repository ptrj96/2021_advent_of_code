package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/ptrj96/go/src/AOC2021/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for scanner.Scan() {
		for i, c := range strings.Split(scanner.Text(), "") {
			x, _ := strconv.Atoi(c)
			res[i] += convBinary(x)
		}
	}

	var gFinal string = ""
	var eFinal string = ""
	for i := 0; i < 12; i++ {
		if res[i] > 0 {
			gFinal += "1"
			eFinal += "0"
		} else {
			gFinal += "0"
			eFinal += "1"
		}
	}

	gVal, _ := strconv.ParseUint(gFinal, 2, 32)
	eVal, _ := strconv.ParseUint(eFinal, 2, 32)
	fmt.Println(gVal * eVal)
}

func convBinary(a int) int {
	if a == 0 {
		return -1
	}
	return a
}
