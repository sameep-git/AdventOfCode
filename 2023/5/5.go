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
}

func p1() {
	seeds := strings.Split("432563865 39236501 1476854973 326201032 1004521373 221995697 2457503679 46909145 603710475 11439698 1242281714 12935671 2569215463 456738587 3859706369 129955069 3210146725 618372750 601583464 1413192", " ")
	_ = seeds

	soils := parseInput("soils.txt")
	fmt.Println(soils)
}

func parseInput(input string) [][]int {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fIn := [][]int{}

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		rowInt := []int{}
		for _, el := range row {
			elInt, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println("error converting")
			} else {
				rowInt = append(rowInt, elInt)
			}
		}
		fIn = append(fIn, rowInt)
	}

	return fIn
}
