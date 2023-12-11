package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction struct {
	x int
	y int
}

func readLines(path string, sep string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), sep)
}

var NORTH = Direction{
	x: 0,
	y: -1,
}
var SOUTH = Direction{
	x: 0,
	y: 1,
}
var EAST = Direction{
	x: 1,
	y: 0,
}
var WEST = Direction{
	x: -1,
	y: 0,
}

type Position struct {
	r int
	c int
}



func getFarthestTile(point Position, visited *map[[2]int]bool, grid []string) {
	if point.r < 0 || point.r >= len(grid) || point.c < 0 || point.c >= len(grid[0]) || (*visited)[[2]int{point.r, point.c}] {
        return
    }
	
	char := string(grid[point.r][point.c])
	(*visited)[[2]int{point.r,point.c}] = true


	// North
	if point.r > 0 && strings.Contains("S|JL", char)   {
		top := grid[point.r + NORTH.y][point.c]

		if strings.Contains("|7F", string(top)) {
			getFarthestTile(Position{
				r: point.r + NORTH.y,
				c: point.c + NORTH.x,
			}, visited, grid)
		}
	}

	// South
	if point.r < len(grid) - 1 && strings.Contains("S|7F", char) {
		bottom := grid[point.r + SOUTH.y][point.c + SOUTH.x]

		if strings.Contains("|JL", string(bottom)) {
			getFarthestTile(Position{
				r: point.r + SOUTH.y,
				c: point.c + SOUTH.x,
			}, visited, grid)
		}
	}

	// West
	if point.c > 0 && strings.Contains("S-J7", char) {
		left := grid[point.r + WEST.y][point.c + WEST.x]

		if strings.Contains("-LF", string(left)) {
			getFarthestTile(Position{
				r: point.r + WEST.y,
				c: point.c + WEST.x,
			}, visited, grid)
		}
	}

	// East
	if point.c < len(grid[point.r]) - 1 && strings.Contains("S-LF", char) {
		right := grid[point.r + EAST.y][point.c + EAST.x]

		if strings.Contains("-J7", string(right)) {
			getFarthestTile(Position{
				r: point.r + EAST.y,
				c: point.c + EAST.x,
			}, visited, grid)
		}
	}
}

func raycastWest(lines []string, visited map[[2]int]bool, r, c int) int {
	intersections := 0
	for i, ch := range lines[r] {
		if i < c {
			if visited[[2]int{r,i}] && strings.Contains("JL|", string(ch)) {
				intersections++
			}
		} else {
			break
		}
	}
	return intersections
}

func main()  {
	lines := readLines("./input.txt", "\n")
	start := Position{}
	for r, line := range lines {
		for c, char := range line {
			if char == 'S' {
				start.r = r
				start.c = c
			}
		}
	}

	visited := make(map[[2]int]bool)

	getFarthestTile(Position{
		r: start.r,
		c: start.c,
	}, &visited, lines)

	fmt.Println(len(visited) / 2)

	visited[[2]int{start.r, start.c}] = true

	lines[start.r] = strings.Replace(lines[start.r], "S", "J", -1) // Hardcoded for the given input

	count := 0
	for r, line := range lines {
		for c, _ := range line {
			if !visited[[2]int{r,c}] {
				intersect := raycastWest(lines, visited, r, c)
				if intersect % 2 == 1 {
					count += 1		
				}
			}
		}
	}
	fmt.Println(count)
}

