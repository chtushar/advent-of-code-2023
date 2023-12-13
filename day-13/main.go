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

func readLines(path string, sep string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), sep)
}

func getReflectionIndex(patterns []string) int {
	for i := 1; i < len(patterns); i++ {
		if patterns[i] != patterns[i-1] {
			continue
		}

		fwd, bck := i, i-1
		isValid := true

		for j := 0; j < min(i, len(patterns)-i)-1; j++ {
			fwd++
			bck--
			if patterns[fwd] != patterns[bck] {
				isValid = false
				break
			}
		}

		if isValid {
			return i
		}
	}

	return 0
}

func main()  {
	lines := readLines("./input.txt", "\n\n")
	
	sum := 0
	for _, r := range lines {
		l := strings.Split(r, "\n")
		cols := []string{}
		rows := []string{}
		
		rN, cN := len(l), len(l[0])

		for c := 0; c < cN; c++ {
			var col string
			for r := 0; r < rN; r++ {
				col += string(l[r][c])
			}
			cols = append(cols, col)
		}

		for r := 0; r < rN; r++ {
			var row string
			for c := 0; c < cN; c++ {
				row += string(l[r][c])
			}
			rows = append(rows, row)
		}

		sum += getReflectionIndex(rows) * 100
		sum += getReflectionIndex(cols)
	}


	fmt.Println(sum)
}