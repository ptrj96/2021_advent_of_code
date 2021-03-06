package main

import "fmt"

func main() {
	lFish := []int{1, 4, 2, 4, 5, 3, 5, 2, 2, 5, 2, 1, 2, 4, 5, 2, 3, 5, 4, 3, 3, 1, 2, 3, 2, 1, 4, 4, 2, 1, 1, 4, 1, 4, 4, 4, 1, 4, 2, 4, 3, 3, 3, 3, 1, 1, 5, 4, 2, 5, 2, 4, 2, 2, 3, 1, 2, 5, 2, 4, 1, 5, 3, 5, 1, 4, 5, 3, 1, 4, 5, 2, 4, 5, 3, 1, 2, 5, 1, 2, 2, 1, 5, 5, 1, 1, 1, 4, 2, 5, 4, 3, 3, 1, 3, 4, 1, 1, 2, 2, 2, 5, 4, 4, 3, 2, 1, 1, 1, 1, 2, 5, 1, 3, 2, 1, 4, 4, 2, 1, 4, 5, 2, 5, 5, 3, 3, 1, 3, 2, 2, 3, 4, 1, 3, 1, 5, 4, 2, 5, 2, 4, 1, 5, 1, 4, 5, 1, 2, 4, 4, 1, 4, 1, 4, 4, 2, 2, 5, 4, 1, 3, 1, 3, 3, 1, 5, 1, 5, 5, 5, 1, 3, 1, 2, 1, 4, 5, 4, 4, 1, 3, 3, 1, 4, 1, 2, 1, 3, 2, 1, 5, 5, 3, 3, 1, 3, 5, 1, 5, 3, 5, 3, 1, 1, 1, 1, 4, 4, 3, 5, 5, 1, 1, 2, 2, 5, 5, 3, 2, 5, 2, 3, 4, 4, 1, 1, 2, 2, 4, 3, 5, 5, 1, 1, 5, 4, 3, 1, 3, 1, 2, 4, 4, 4, 4, 1, 4, 3, 4, 1, 3, 5, 5, 5, 1, 3, 5, 4, 3, 1, 3, 5, 4, 4, 3, 4, 2, 1, 1, 3, 1, 1, 2, 4, 1, 4, 1, 1, 1, 5, 5, 1, 3, 4, 1, 1, 5, 4, 4, 2, 2, 1, 3, 4, 4, 2, 2, 2, 3}

	var totalFish [9]int

	for i := 0; i < len(lFish); i++ {
		totalFish[lFish[i]]++
	}

	for i := 0; i < 256; i++ {
		newFish := totalFish[0]
		for j := 0; j < 8; j++ {
			totalFish[j] = totalFish[j+1]
		}
		totalFish[6] += newFish
		totalFish[8] = newFish
	}

	sum := 0
	for i := 0; i < len(totalFish); i++ {
		sum += totalFish[i]
	}

	fmt.Println(sum)
}
