package main

import (
	"fmt"
	"os"
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

type Map struct {
	this string
	left string 
	right string
}

func countSteps(path string, index int, maps map[string]Map, current Map) int {

	if  solution == 1 && current.this == "ZZZ" {
		return 0
	}

	if  solution == 2 && strings.HasSuffix(current.this, "Z") {
		return 0
	}

	if path[index] == 'L' {
		current = maps[current.left]
	} else if (path[index] == 'R') {
		current = maps[current.right]
	}

	if index >= len(path) - 1 {
		index = -1
	}

	return 1 + countSteps(path, index + 1, maps, current)
}

func countSteps2(path string, index int, maps map[string]Map, current []Map) int {
	
	zCount := 0
	for _, c := range current {
		if strings.HasSuffix(c.this, "Z") {
			zCount += 1
		}
	}

	if len(current) == zCount {
		return 0
	}

	newCurrent := []Map{}
	if path[index] == 'L' {
		for _, m := range current {
			newCurrent = append(newCurrent, maps[m.left])
		}
	} else if (path[index] == 'R') {
		for _, m := range current {
			newCurrent = append(newCurrent, maps[m.right])
		}
	}

	if index >= len(path) - 1 {
		index = -1
	}


	return 1 + countSteps2(path, index + 1, maps, newCurrent)
}

func gcd(num1, num2 uint64) uint64 {
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	return num1
}

func lcm(num1, num2 uint64, rest ...uint64) uint64 {
	result := num1 * num2 / gcd(num1, num2)
	for i := 0; i < len(rest); i++ {
		result = lcm(result, rest[i])
	}
	return result
}


const solution = 2

func main()  {
	lines := readLines("./input.txt", "\n\n")
	maps := make(map[string]Map)

	for _, s := range strings.Split(lines[1], "\n") {
		splits := strings.Split(s, " = ")
		this := splits[0]
		direction := strings.Split(strings.Trim(strings.Trim(splits[1], ")"), "("), ", ")

		maps[this] = Map{
			this: this,
			left: direction[0],
			right: direction[1],
		}
	}

	if solution == 1 {
		count := countSteps(lines[0], 0, maps, maps["AAA"])
		fmt.Println(count)
		
	}

	if solution == 2 {
		steps := []uint64{}
		for m := range maps {
			if strings.HasSuffix(m, "A") {
				steps = append(steps, uint64(countSteps(lines[0], 0, maps, maps[m])))
			}
		}
		fmt.Println(lcm(steps[0], steps[1], steps[2:]...))
	}
}
