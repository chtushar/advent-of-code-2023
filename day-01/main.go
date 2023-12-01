package main

import (
	"fmt"
	"os"
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

func txtToNumber(l *string) {
	*l = strings.ReplaceAll(*l, "one", "on1e");
	*l = strings.ReplaceAll(*l, "two", "tw2o");
	*l = strings.ReplaceAll(*l, "three", "thre3e");
	*l = strings.ReplaceAll(*l, "four", "fou4r");
	*l = strings.ReplaceAll(*l, "five", "fiv5e");
	*l = strings.ReplaceAll(*l, "six", "si6x");
	*l = strings.ReplaceAll(*l, "seven", "seve7n");
	*l = strings.ReplaceAll(*l, "eight", "eigh8t");
	*l = strings.ReplaceAll(*l, "nine", "nin9e");
}

func getSum(l *string) int {
	length := len(*l)
	lp := 0
	rp := length - 1
	sum := 0
	foundLeftNumber := false
	foundRightNumber := false

	chars := strings.Split(*l, "")

	for !(foundLeftNumber && foundRightNumber) && lp < length && rp > -1 {
		if !foundLeftNumber {
			num, err := strconv.Atoi(chars[lp])
			if err != nil {
				lp++
			} else {
				sum += num*10
				foundLeftNumber = true
			}
		}

		if !foundRightNumber {
			num, err := strconv.Atoi(chars[rp])
			if err != nil {
				rp--
			} else {
				sum += num
				foundRightNumber = true
			}
		}
	}

	return sum
}

func main()  {
	lines := readLines("./input-1.txt");
	
	sum := 0;
	for _, l := range lines {
		txtToNumber(&l)
		sum += getSum(&l)
	}
	fmt.Println("Sum is: ", sum)
}