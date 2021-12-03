package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	num  int
	zero *Node
	one  *Node
}

func Insert(s string, n *Node) {
	if len(s) == 0 {
		return
	}
	if s[0:1] == "0" {
		if n.zero != nil {
			n.zero.num++
		} else {
			n.zero = &Node{num: 1}
		}
		Insert(s[1:], n.zero)
	} else {
		if n.one != nil {
			n.one.num++
		} else {
			n.one = &Node{num: 1}
		}
		Insert(s[1:], n.one)
	}
}

func GetMost(n Node) string {
	if n.one == nil && n.zero == nil {
		return ""
	}

	if (n.one != nil && n.zero != nil && n.one.num >= n.zero.num) || (n.one != nil && n.zero == nil) {
		return "1" + GetMost(*n.one)
	} else {
		return "0" + GetMost(*n.zero)
	}
}

func GetLeast(n Node) string {
	if n.one == nil && n.zero == nil {
		return ""
	}

	if (n.one != nil && n.zero != nil && n.one.num < n.zero.num) || (n.one != nil && n.zero == nil) {
		return "1" + GetLeast(*n.one)
	} else {
		return "0" + GetLeast(*n.zero)
	}
}

func main() {
	file, err := os.Open("/home/ptrj96/go/src/AOC2021/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	root := Node{num: 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Insert(scanner.Text(), &root)
	}

	gVal, _ := strconv.ParseUint(GetMost(root), 2, 32)
	lVal, _ := strconv.ParseUint(GetLeast(root), 2, 32)

	fmt.Println(gVal * lVal)
}
