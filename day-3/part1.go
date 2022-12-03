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
		line := scanner.Text()
		chars := strings.Split(line, "")

		comp1 := chars[:len(chars)/2]
		comp2 := chars[len(chars)/2:]

		match := ""
		for _, char1 := range comp1 {
			for _, char2 := range comp2 {
				if char1 == char2 {
					match = char1
					break
				}
			}
			if match != "" {
				break
			}
		}

		total += charMap[match]
	}

	fmt.Println(total)
}
