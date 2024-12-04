package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	p1()
	p2()
}

func p1() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	ans := 0
	for i, line := range input {
		for j, char := range line {
			if char == 'X' {
				// Right check
				if j+3 < len(line) && line[j+1:j+4] == "MAS" {
					ans += 1
				}
				// Left check
				if j-3 >= 0 && line[j-3:j] == "SAM" {
					ans += 1
				}
				if i-3 >= 0 {
					// Up check
					if input[i-3][j] == 'S' && input[i-2][j] == 'A' && input[i-1][j] == 'M' {
						ans += 1
					}
					// Up left check
					if j-3 >= 0 && input[i-3][j-3] == 'S' && input[i-2][j-2] == 'A' && input[i-1][j-1] == 'M' {
						ans += 1
					}
					// Up right check
					if j+3 < len(line) && input[i-3][j+3] == 'S' && input[i-2][j+2] == 'A' && input[i-1][j+1] == 'M' {
						ans += 1
					}
				}
				if i+3 < len(input) {
					// Bottom Check
					if input[i+3][j] == 'S' && input[i+2][j] == 'A' && input[i+1][j] == 'M' {
						ans += 1
					}
					// Bottom left check
					if j-3 >= 0 && input[i+3][j-3] == 'S' && input[i+2][j-2] == 'A' && input[i+1][j-1] == 'M' {
						ans += 1
					}
					// Bottom right check
					if j+3 < len(line) && input[i+3][j+3] == 'S' && input[i+2][j+2] == 'A' && input[i+1][j+1] == 'M' {
						ans += 1
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	ans := 0
	for i, line := range input {
		for j, char := range line {
			if char == 'A' {
				/*
					M		M	M		S	S		S	S		M
						A			A			A			A
					S		S	M		S	M		M	S		M
				*/
				if i-1 >= 0 && i+1 < len(input) && j-1 >= 0 && j+1 < len(line) {
					// Case 1
					if input[i-1][j-1] == 'M' && input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i+1][j+1] == 'S' {
						ans += 1
					}
					// Case 2
					if input[i-1][j-1] == 'M' && input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i+1][j+1] == 'S' {
						ans += 1
					}
					// Case 3
					if input[i-1][j-1] == 'S' && input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M' && input[i+1][j+1] == 'M' {
						ans += 1
					}
					// Case 4
					if input[i-1][j-1] == 'S' && input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S' && input[i+1][j+1] == 'M' {
						ans += 1
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
