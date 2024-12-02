package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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

	ans := 0
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		isSafe := true
		increasing := 0
		// 0 for nil
		// 1 for increasing
		// -1 for decreasing
		for idx := range len(nums) - 1 {
			num1, err1 := strconv.Atoi(nums[idx])

			if err1 != nil {
				fmt.Println("error1")
			}
			num2, err2 := strconv.Atoi(nums[idx+1])

			if err2 != nil {
				fmt.Println("error1")
			}

			if num1 > num2 {
				if increasing == 0 {
					increasing = -1
				} else if increasing == 1 {
					isSafe = false
					break
				}
			} else {
				if increasing == 0 {
					increasing = 1
				} else if increasing == -1 {
					isSafe = false
					break
				}
			}
			if abs(num1, num2) < 1 || abs(num1, num2) > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			ans += 1
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

	ans := 0
	for scanner.Scan() {
		nums := strings.Fields(scanner.Text())
		// 0 for nil
		// 1 for increasing
		// -1 for decreasing
		tempAns := 0
		for i := range len(nums) {
			increasing := 0
			tmp := slices.Clone(nums)
			tmp2 := append(tmp[:i], tmp[i+1:]...)
			isGood := true
			for idx := range len(tmp2) - 1 {
				num1, err1 := strconv.Atoi(tmp2[idx])

				if err1 != nil {
					fmt.Println("error1")
				}
				num2, err2 := strconv.Atoi(tmp2[idx+1])

				if err2 != nil {
					fmt.Println("error1")
				}

				if num1 > num2 {
					if increasing == 0 {
						increasing = -1
					} else if increasing == 1 {
						isGood = false
					}
				} else {
					if increasing == 0 {
						increasing = 1
					} else if increasing == -1 {
						isGood = false
					}
				}
				if abs(num1, num2) < 1 || abs(num1, num2) > 3 {
					isGood = false
				}
			}
			if isGood {
				tempAns += 1
			}
		}
		if tempAns > 0 {
			ans += 1
		}
	}
	fmt.Println(ans)
}

func abs(n1 int, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	} else {
		return n2 - n1
	}
}
