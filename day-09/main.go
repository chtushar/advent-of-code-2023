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

func convToIntSlice(strs []string) []int  {
	nums := []int{}
	for _, s := range strs {
		num := convStrToInt(s)
		nums = append(nums, num)
	}
	return nums
}

func convStrToInt (s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return num
}

func getNextNumber(seq []int) int {
	zeroCount := 0
	for _, s := range seq {
		if s == 0 {
			zeroCount += 1
		}
	}
	if len(seq) == zeroCount { 
		return 	0
	}

	// Finding the diffs
	seq2 := seq[1:]
	diffs := []int{}
	for i, v := range seq2 {
		diffs = append(diffs, v - seq[i])
	}

	return seq[len(seq) - 1] + getNextNumber(diffs)
}

func getPrevNumber(seq []int) int {
	zeroCount := 0
	for _, s := range seq {
		if s == 0 {
			zeroCount += 1
		}
	}
	if len(seq) == zeroCount { 
		return 	0
	}

	// Finding the diffs
	seq2 := seq[1:]
	diffs := []int{}
	for i, v := range seq2 {
		diffs = append(diffs, v - seq[i])
	}

	return seq[0] - getPrevNumber(diffs)
}

const solution = 2
func main() {
	lines := readLines("./input.txt", "\n")
	sequences := [][]int{}
	for _, l := range lines {
		splits := strings.Split(l, " ")
		sequences = append(sequences, convToIntSlice(splits))
	}

	ans := 0
	if solution == 1 {
		for _, s := range sequences {
			nextNum := getNextNumber(s)
			ans += nextNum
		}
	}

	if solution == 2 {
		for _, s := range sequences {
			nextNum := getPrevNumber(s)
			ans += nextNum
		}
	}

	fmt.Println(ans)
}
