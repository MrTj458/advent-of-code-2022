package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	elves := []int{0}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elves = append(elves, 0)
			continue
		}

		calories, _ := strconv.Atoi(text)
		elves[len(elves)-1] += calories
	}

	sort.Ints(elves)
	max := elves[len(elves)-1]
	second := elves[len(elves)-2]
	third := elves[len(elves)-3]

	fmt.Println(max, second, third)
	fmt.Println("Total:", max+second+third)
}
