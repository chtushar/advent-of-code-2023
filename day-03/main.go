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

var dotCode rune = 46;
var zeroCode rune = 48;
var nineCode rune = 57;
var gearCode rune = 42;
var coords [][]int = [][]int {{-1,-1},{0,-1},{1,-1},{-1,0},{1,0},{-1,1},{0,1},{1,1}};

func charIsNumber (ch rune) bool{
	return ch >= zeroCode && ch <= nineCode 
}

func getCoordKey(x int,y int) string {
	var b strings.Builder
    fmt.Fprintf(&b, "%d,%d", x, y)
    return b.String()
}

func parseCoordKey(coordKey string) []int {
	var coordString []string = strings.Split(coordKey, ",")
	x, err := strconv.Atoi(coordString[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	y, err := strconv.Atoi(coordString[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	return []int{x,y}
}

func getStartingDigitCoord(x int, y int, lines []string) []int {
	j := y
	for j > -1 && charIsNumber(rune(lines[x][j])) {
		j--
	}
	return []int{x,j+1}
}

func getNumberFromCoord(coord []int, lines []string) int {
	i := 0
	numStr := ""
	for coord[1] + i < len(lines[0]) && charIsNumber(rune(lines[coord[0]][coord[1] + i])) {
		numStr += string(lines[coord[0]][coord[1] + i])
		i++
	}

	num, err := strconv.Atoi(numStr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return num
}

func solution1 () {
	lines := readLines("./input-1.txt")
	height := len(lines)
	width := len(lines[0])
	set := make(map[string]bool);

	for j, r := range lines {
		for i, c := range r {
			if c == dotCode || charIsNumber(c) {
				continue
			}
			
			for _, coord := range coords {
				y := i + coord[0];
				x := j + coord[1];

				if x < 0 && x > width || y < -1 && y > height {
					return
				}
				
				if charIsNumber(rune(lines[x][y])) {
					startingCoord := getStartingDigitCoord(x,y,lines)
					coordKey := getCoordKey(startingCoord[0], startingCoord[1])
					set[coordKey] = true
				}
			}
		}
	}

	sum := 0
	for c := range set {
		coord := parseCoordKey(c)
		num := getNumberFromCoord(coord, lines)
		sum += num
	}

	fmt.Println(sum)
}

func solution2()  {
	lines := readLines("./input-1.txt")
	height := len(lines)
	width := len(lines[0])
	sum := 0

	for j, r := range lines {
		for i, c := range r {
			if c == dotCode || charIsNumber(c) {
				continue
			}

			if c == gearCode {
				var set = make(map[string]bool);
				for _, coord := range coords {
					x := j + coord[1];
					y := i + coord[0];
	
					if x < 0 && x > width || y < -1 && y > height {
						return
					}
					
					if charIsNumber(rune(lines[x][y])) {
						startingCoord := getStartingDigitCoord(x,y,lines)
						key := getCoordKey(startingCoord[0], startingCoord[1])
						set[key] = true
					}
				}
				if len(set) == 2 {
					ratio := 1
					for key := range set {
						startingCoord := parseCoordKey(key)
						num := getNumberFromCoord(startingCoord, lines)

						ratio *= num
					}
					sum += ratio
				}
			}

		}
	}

	fmt.Println(sum)
}

func main()  {
	solution2()	
}