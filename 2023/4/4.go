package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	// p1()
	p2()
}

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ans := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")[1]
		winning := strings.Split(s, "|")[0]
		have := strings.Split(s, "|")[1]

		winningList := strings.FieldsFunc(winning, func(r rune) bool {
			return r == ' '
		})
		haveList := strings.FieldsFunc(have, func(r rune) bool {
			return r == ' '
		})

		c := 0
		for _, num := range haveList {
			if slices.Contains(winningList, num) {
				c++
			}
		}
		if c > 0 {
			ans += (1 << (c - 1))
		}
	}

	fmt.Println(ans)
}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// ans := 0
	cards := make([]int, 214)
	for i := range cards {
		cards[i] = 1
	}
	winningList := []string{}
	_ = winningList
	haveList := []string{}
	_ = haveList
	row := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")[1]
		winning := strings.Split(s, "|")[0]
		have := strings.Split(s, "|")[1]

		winningList = strings.FieldsFunc(winning, func(r rune) bool {
			return r == ' '
		})
		haveList = strings.FieldsFunc(have, func(r rune) bool {
			return r == ' '
		})
		c := 0
		for _, num := range haveList {
			if slices.Contains(winningList, num) {
				c++
			}
		}
		itr := row + 1
		for itr < row+c+1 {
			if itr < len(cards) {
				cards[itr] += cards[row]
			}
			itr += 1
		}
		row += 1
	}
	ans := 0
	for _, num := range cards {
		ans += num
	}
	fmt.Println(ans)
}
