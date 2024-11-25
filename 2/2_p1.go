package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0
	i := 1
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		s = strings.Split(s[1], ";")
		isGood := true
		for _, substring := range s {
			a := strings.Split(substring, ",")
			for _, play := range a {
				chance := strings.Split(play[1:], " ")
				num, err := strconv.Atoi(chance[0])
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					if chance[1] == "red" && num > 12 {
						isGood = false
					} else if chance[1] == "green" && num > 13 {
						isGood = false
					} else if chance[1] == "blue" && num > 14 {
						isGood = false
					}
				}
			}
		}
		if isGood {
			ans += i
		}
		i++
	}
	fmt.Printf("Answer to p1: %d\n", ans)
}

func p2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := 0
	i := 1
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		s = strings.Split(s[1], ";")
		// isGood := true
		minRed := 0
		minBlue := 0
		minGreen := 0
		for _, substring := range s {
			a := strings.Split(substring, ",")
			for _, play := range a {
				chance := strings.Split(play[1:], " ")
				num, err := strconv.Atoi(chance[0])
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					if chance[1] == "red" && num > minRed {
						minRed = num
					} else if chance[1] == "green" && num > minGreen {
						minGreen = num
					} else if chance[1] == "blue" && num > minBlue {
						minBlue = num
					}
				}
			}
		}
		ans += minBlue * minGreen * minRed
		i++
	}
	fmt.Printf("Answer to p2: %d\n", ans)
}
