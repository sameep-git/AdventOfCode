package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//p1()
	p2()
}

func p1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num1, err1 := strconv.Atoi(nums[0])

		if err1 != nil {
			fmt.Println("error1")
		}
		num2, err2 := strconv.Atoi(nums[1])

		if err2 != nil {
			fmt.Println("error1")
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	ans := 0

	for idx := range len(list1) {
		if list1[idx] > list2[idx] {
			ans += (list1[idx] - list2[idx])
		} else {
			ans += (list2[idx] - list1[idx])
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

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		num1, err1 := strconv.Atoi(nums[0])

		if err1 != nil {
			fmt.Println("error1")
		}
		num2, err2 := strconv.Atoi(nums[1])

		if err2 != nil {
			fmt.Println("error1")
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	freqMap := make(map[int]int)
	for _, num := range list2 {
		freqMap[num] = freqMap[num] + 1
	}

	ans := 0

	for _, num := range list1 {
		count, ok := freqMap[num]
		if ok {
			ans += (num * count)
		} else {
			fmt.Println("not in")
		}
	}

	fmt.Println(ans)
}
