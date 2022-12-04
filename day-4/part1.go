package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		range1 := strings.Split(strings.Split(line, ",")[0], "-")
		range2 := strings.Split(strings.Split(line, ",")[1], "-")

		var section1 []int
		for _, num := range range1 {
			i, _ := strconv.Atoi(num)
			section1 = append(section1, i)
		}

		var section2 []int
		for _, num := range range2 {
			i, _ := strconv.Atoi(num)
			section2 = append(section2, i)
		}

		containing :=
			(section1[0] >= section2[0] &&
				section1[0] <= section2[1] &&
				section1[1] >= section2[0] &&
				section1[1] <= section2[1]) ||
				(section2[0] >= section1[0] &&
					section2[0] <= section1[1] &&
					section2[1] >= section1[0] &&
					section2[1] <= section1[1])

		if containing {
			total += 1
		}

		fmt.Println(section1, section2, containing)
	}

	fmt.Println(total)
}
