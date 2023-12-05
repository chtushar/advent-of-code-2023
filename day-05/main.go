package main

import (
	"fmt"
	"os"
	"strconv"
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

func readLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n\n")
}

func getSrcDestinationLabel(str string) (string, string) {
	split := strings.Split(strings.Split(str, " ")[0], "-to-")
	return split[0], split[1]
}

var rangesMap map[string][][]int
var srcDestMap map[string]string;

func findLocation(num int, src string, dest string) int {
	if src == "location" {
		return num
	}
	nextDest := srcDestMap[dest];
	ranges := rangesMap[dest]
	nextNum := num
	for _, r := range ranges {
		if num >= r[1] && num <= r[1] + r[2] {
			nextNum = r[0] + (num - r[1])
		}
	}

	return findLocation(nextNum, dest, nextDest)
}

func getLocationRanges(seedRanges [][]int, src string, dest string) [][]int {
	unmappedRange := [][]int{}
	mappedRange := [][]int{}
	
	for _, r := range rangesMap[dest] {
		seedRanges = append(seedRanges, unmappedRange...)
		unmappedRange = unmappedRange[:0]

		d, s, offset := r[0], r[1], r[2]
		for _, sR := range seedRanges {
			start, end := sR[0], sR[1]
			// shift the first range
			seedRanges = seedRanges[1:]

			// inside the range
			inS := max(start, s)
			inE := min(end, s + offset)
		
			if inS < inE {
				mappedRange = append(mappedRange, []int{
					inS - s + d,
					inE - s + d,
				})

				if inS > start {
					unmappedRange = append(unmappedRange, []int{
						start,
						inS,
					})
				}

				if end > inE {
					unmappedRange = append(unmappedRange, []int{
						inE,
						end,
					})
				}
			} else {
				unmappedRange = append(unmappedRange, []int{start, end})
			}
		}
	}

	if len(unmappedRange) > 0 {
		mappedRange = append(mappedRange, unmappedRange...)
	}
	
	if dest == "location" {
		return mappedRange
	}
	nextDest := srcDestMap[dest]

	return getLocationRanges(mappedRange, dest, nextDest)
}

func main()  {
	maps := readLines("./input-1.txt")
	seedStr := strings.Split(strings.TrimSpace(strings.Split(maps[0], ":")[1]), " ");
	var seeds []int;
	rangesMap = make(map[string][][]int)
	srcDestMap = make(map[string]string)

	for _, s := range seedStr {
		if s != "" {
			sn, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			seeds = append(seeds, int(sn))
		}
	}

	for _, m := range maps[1:] {
		splits := strings.Split(m, "\n")
		src, dest := getSrcDestinationLabel(splits[0])

		srcDestMap[src] = dest
		var ranges [][]int
		for _, s := range splits[1:] {
			numsStr := strings.Split(s, " ")
			var nums []int
			for _, s := range numsStr {
				num, err := strconv.Atoi(s)
				if err != nil {
					os.Exit(1)
					fmt.Println(err)
				}
				nums = append(nums, int(num))
			}
			ranges = append(ranges, nums)
		}

		rangesMap[dest] = ranges
	}

	l := int(^uint(0) >> 1)
	for _, s := range seeds {
		n := findLocation(s, "seed", "soil")
		if n < l {
			l = n
		}
	}
	fmt.Println("first solution:", l)

	var seedsRange [][]int

	for i := 0; i < len(seeds); i+=2 {
		seedsRange = append(seedsRange, []int{seeds[i], seeds[i] + seeds[i + 1]})
	}

	r := getLocationRanges(seedsRange, "seed", "soil")

	var lowest = r[0][0]
	for i := 1; i < len(r); i++ {
		lowest = min(lowest, r[i][0])
	}
	fmt.Printf("second solution: %d\n", lowest)
}