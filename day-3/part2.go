package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	charMap := make(map[string]int)
	for i, char := range strings.Split(letters, "") {
		charMap[char] = i + 1
	}

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		match := ""
		for _, char := range strings.Split(line1, "") {
			if strings.Contains(line2, char) && strings.Contains(line3, char) {
				match = char
				break
			}
		}

		total += charMap[match]
	}

	fmt.Println(total)
}
