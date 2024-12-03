package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	p1()
	p2()
}

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	r := regexp.MustCompile("(?i)mul\\([0-9]+,[0-9]+\\)")
	var text []string
	_ = text
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	points := r.FindAllString(strings.Join(text[:], ""), -1)

	fmt.Println(points)
	ans := 0
	for _, point := range points {
		s := strings.Split(point, "mul(")[1]
		s1 := strings.Split(s, ")")[0]
		s2 := strings.Split(s1, ",")
		num1, err1 := strconv.Atoi(s2[0])

		if err1 != nil {
			fmt.Println("error1")
		}
		num2, err2 := strconv.Atoi(s2[1])

		if err2 != nil {
			fmt.Println("error1")
		}

		ans += (num1 * num2)
	}
	fmt.Println(ans)
}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var text []string
	_ = text
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	fullText := strings.Join(text[:], "")

	do := regexp.MustCompile("^do\\(\\)$")
	dont := regexp.MustCompile("^don't\\(\\)$")
	mul := regexp.MustCompile("mul\\([0-9]+,[0-9]+\\)")

	idx := 0
	what2do := 0
	// 0 normal
	// 1 do
	// -1 dont
	ans := 0
	for idx < len(fullText) {
		if idx+4 < len(fullText) {
			matchedDo := do.MatchString(fullText[idx : idx+4])
			if matchedDo {
				what2do = 1
				idx += 4
				continue
			}
		}
		if idx+7 < len(fullText) {
			matchedDont := dont.MatchString(fullText[idx : idx+7])
			if matchedDont {
				what2do = -1
				idx += 7
				continue
			}
		}
		if idx+12 < len(fullText) {
			matchedStr := mul.FindString(fullText[idx : idx+12])
			if len(matchedStr) > 0 && what2do >= 0 {
				s := strings.Split(matchedStr, "mul(")[1]
				s1 := strings.Split(s, ")")[0]
				s2 := strings.Split(s1, ",")
				num1, err1 := strconv.Atoi(s2[0])

				if err1 != nil {
					fmt.Println("error1")
				}
				num2, err2 := strconv.Atoi(s2[1])

				if err2 != nil {
					fmt.Println("error1")
				}
				ans += (num1 * num2)
				what2do = 0
				idx += len(matchedStr)
				continue
			}
		}
		idx += 1
	}

	fmt.Println(ans)

}
