package main

import (
	"fmt"
	"os"
	"strings"
)

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func readLines(path string, sep string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), sep)
}

func insertString(slice []string, index int, value string) []string {
    if index < 0 || index > len(slice) {
        return slice
    }

    result := make([]string, 0, len(slice)+1)

    result = append(result, slice[:index]...)

    result = append(result, value)

    result = append(result, slice[index:]...)

    return result
}

func main()  {
	lines := readLines("./input.txt", "\n")
	rowsWithNoGal := make(map[int]bool)
	colsWithNoGal := make(map[int]bool)
	galaxies := [][2]int{}
	galaxiesCount := 0

	for r, line := range lines {
		if !strings.Contains(line, "#") {
			rowsWithNoGal[r] = true
		} else {
			galaxiesCount += strings.Count(line, "#")
		}
		for c := range line {
			if lines[r][c] == '#' {
				galaxies = append(galaxies, [2]int{r,c})
			}
			for i := range lines {
				if lines[i][c] == '#' {
					break
				}
				if i == len(lines) - 1 {
					colsWithNoGal[c] = true
				}
			}
		}
	}

	count := 0
	expandBy := 1000000
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			for k := min(g1[0], g2[0]); k < max(g1[0], g2[0]); k++ {
				if rowsWithNoGal[k] {
					count += expandBy 
				} else {
					count += 1
				}
			}

			for k := min(g1[1], g2[1]); k < max(g1[1], g2[1]); k++ {
				if colsWithNoGal[k] {
					count += expandBy
				} else {
					count += 1
				}
			}
		}
	}

	fmt.Println(count)
}