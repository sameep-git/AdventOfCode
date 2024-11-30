package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	p1()
}

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	engine := []string{}

	for scanner.Scan() {
		row := scanner.Text()
		engine = append(engine, row)
	}

	isNumber := false
	isPart := false
	currNumber := 0
	ans := 0
	for i, row := range engine {
		for j, char := range row {
			if unicode.IsDigit(char) {
				currNumber *= 10
				currNumber += int(char - '0')
				isNumber = true
				if (i - 1) >= 0 {
					if engine[i-1][j] != '.' && !unicode.IsDigit(rune(engine[i-1][j])) {
						isPart = true
					}
					if (j - 1) >= 0 {
						if engine[i-1][j-1] != '.' && !unicode.IsDigit(rune(engine[i-1][j-1])) {
							isPart = true
						}
					}
					if (j + 1) < len(engine[0]) {
						if engine[i-1][j+1] != '.' && !unicode.IsDigit(rune(engine[i-1][j+1])) {
							isPart = true
						}
					}
				}
				if (i + 1) < len(engine) {
					if engine[i+1][j] != '.' && !unicode.IsDigit(rune(engine[i+1][j])) {
						isPart = true
					}
					if (j - 1) >= 0 {
						if engine[i+1][j-1] != '.' && !unicode.IsDigit(rune(engine[i+1][j-1])) {
							isPart = true
						}
					}
					if (j + 1) < len(engine[0]) {
						if engine[i+1][j+1] != '.' && !unicode.IsDigit(rune(engine[i+1][j+1])) {
							isPart = true
						}
					}
				}
				if (j - 1) >= 0 {
					if engine[i][j-1] != '.' && !unicode.IsDigit(rune(engine[i][j-1])) {
						isPart = true
					}
				}
				if (j + 1) < len(engine[0]) {
					if engine[i][j+1] != '.' && !unicode.IsDigit(rune(engine[i][j+1])) {
						isPart = true
					}
				}
			}
			if char == '.' || !unicode.IsDigit(char) || (j == len(engine[0])-1) {
				if isPart && isNumber {
					fmt.Println(currNumber)
					ans += currNumber
				}
				isNumber = false
				isPart = false
				currNumber = 0
			}
		}
	}
	if isPart && isNumber {
		fmt.Println(currNumber)
		ans += currNumber
	}
	fmt.Println(ans)
}
