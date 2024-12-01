package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0
	textNum := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numDict := map[string]rune{"one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}

	for scanner.Scan() {
		s := scanner.Text()
		str := []rune(s)

		local := ""
		minIdx := int(math.Inf(1))
		maxIdx := int(math.Inf(-1))
		var val string
		_ = val

		for _, num := range textNum {
			idx := strings.Index(s, num)
			if idx != -1 {
				if idx < int(minIdx) {
					minIdx = idx
					val = string(numDict[num])
				}
			}
		}

		for i, char := range str {
			if unicode.IsDigit(char) {
				if i < minIdx {
					minIdx = i
					val = string(char)
				}
				break
			}
		}

		local += val

		for _, num := range textNum {
			idx := strings.LastIndex(s, num)
			if idx != -1 {
				if idx > int(maxIdx) {
					maxIdx = idx
					val = string(numDict[num])
				}
			}
		}

		for i := len(str) - 1; i >= 0; i-- {
			if unicode.IsDigit(str[i]) {
				if i > maxIdx {
					maxIdx = i
					val = string(str[i])
				}
				break
			}
		}

		local += val

		num, err := strconv.Atoi(local)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			ans += num
		}
	}

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
