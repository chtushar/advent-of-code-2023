package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
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


func solution1()  {
	lines := readLines("./input.txt");
	spaceRegex := regexp.MustCompile("\\s+")
	times := convToIntSlice(spaceRegex.Split(strings.TrimSpace(strings.Split(lines[0], "Time:")[1]), -1))
	distances := convToIntSlice(spaceRegex.Split(strings.TrimSpace(strings.Split(lines[1], "Distance:")[1]), -1))

	ans := 1
	for i, t := range times {
		d := distances[i]
		nums := 0
		for holdFor := 0; holdFor <= t; holdFor++ {
			rem := t - holdFor;
			covered := holdFor * rem
			if d < covered {
				nums += 1
			}
		}
		ans *= nums
	}

	fmt.Println(ans)
}

func solution2()  {
	lines := readLines("./input.txt");
	spaceRegex := regexp.MustCompile("\\s+")

	time := convStrToInt(spaceRegex.ReplaceAllString(strings.TrimSpace(strings.Split(lines[0], "Time:")[1]), ""))
	distance := convStrToInt(spaceRegex.ReplaceAllString(strings.TrimSpace(strings.Split(lines[1], "Distance:")[1]), ""))

	nums := 0
	for holdFor := 0; holdFor <= time; holdFor++ {
		rem := time - holdFor;
		covered := holdFor * rem
		if distance < covered {
			nums += 1
		}
	}
	fmt.Println(nums)
}

func main()  {
	solution2()
}
