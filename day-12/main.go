package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func readLines(path string, sep string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), sep)
}

func convStrToInt (s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return num
}

func convToIntSlice(strs []string) []int  {
	nums := []int{}
	for _, s := range strs {
		num := convStrToInt(s)
		nums = append(nums, num)
	}
	return nums
}

var dp [104][104][104]int

func getCount(nums []int, pattern string, times int) int {
	allnums := make([]int, 0)
	var sb strings.Builder
	for times > 0 {
		times--
		sb.WriteString(pattern)
		if times != 0 {
			sb.WriteRune('?')
		}
		allnums = append(allnums, nums...)
	}
	dp = [104][104][104]int{}
	fullpattern := sb.String()
	return getCountRecursively(fullpattern, 0, allnums, 0, 0)
}

func getCountRecursively(pattern string, p int, nums []int, n int, g int) int {
	if len(pattern) == p {
		if (n == len(nums)-1 && nums[n] == g) || (n == len(nums) && g == 0) {
			return 1
		}
		return 0
	}

	if dp[p][n][g] != 0 {
		return dp[p][n][g] - 1
	}
	sum := 0
	char := pattern[p]

	if strings.Contains("?#", string(char)) {
		sum += getCountRecursively(pattern, p+1, nums, n, g+1)
	}
	if strings.Contains("?.", string(char)) {
		if g > 0 && n < len(nums) && nums[n] == g {
			sum += getCountRecursively(pattern, p+1, nums, n+1, 0)
		}
		if g == 0 {
			sum += getCountRecursively(pattern, p+1, nums, n, 0)
		}
	}

	dp[p][n][g] = sum + 1
	return sum
}

func main()  {
	lines := readLines("./input.txt", "\n")
	sum := 0
	
	for _, line := range lines {
		splits := strings.Split(line, " ")
		nums := convToIntSlice(strings.Split(splits[1], ","))
		sum += getCount(nums, splits[0], 5)
	}

	fmt.Println(sum)
}
