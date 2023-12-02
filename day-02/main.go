package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


const maxR = 12;
const maxG = 13;
const maxB = 14;

func readLines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
}

func isGamePossible(game string) bool {
	splits := strings.Split(game, ":");
	sets := strings.Split(strings.TrimSpace(splits[1]), ";")

	for _, set := range sets {
		set = strings.TrimSpace(set)
		balls := strings.Split(set, ", ");

		for _, ball := range balls {
			b := strings.Split(ball, " ")
			amt, err := strconv.Atoi(b[0])
			color := b[1];
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			
			if color == "green" && amt > maxG {
				return false
			} else if color == "red" && amt > maxR {
				return false
			} else if color == "blue" && amt > maxB {
				return false
			}
		}
	}

	return true
}

func calculatePower(game string) int {
	splits := strings.Split(game, ":");
	sets := strings.Split(strings.TrimSpace(splits[1]), ";")
	var max []int = []int{0, 0, 0};

	for _, set := range sets {
		set = strings.TrimSpace(set)
		balls := strings.Split(set, ", ");
		for _, ball := range balls {
			b := strings.Split(ball, " ")
			amt, err := strconv.Atoi(b[0])
			color := b[1];
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if color == "red" && amt > max[0] {
				max[0] = amt
			} else if color == "green" && amt > max[1] {
				max[1] = amt
			} else if color == "blue" && amt > max[2] {
				max[2] = amt
			}
		}
	}

	return max[0]*max[1]*max[2]
}

func solution1(games []string)  {
	sum := 0
	for i, game := range games {
		gameId := i + 1;
		possible := isGamePossible(game)
		if  possible {
			sum += gameId
		}
	}
	fmt.Println(sum)
}

func solution2(games []string)  {
	sum := 0
	for _, game := range games {
		power := calculatePower(game)
		sum += power
	}
	fmt.Println(sum)
}

func main()  {
	games := readLines("./input-1.txt");
	solution2(games)
}
