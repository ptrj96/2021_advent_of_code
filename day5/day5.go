package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path + "/day5/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		start := strings.Split(line[0], ",")
		finish := strings.Split(line[1], ",")
	}
}
