package main

import (
	"fmt"
	"math"
	"os"
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

func getNumbersMap(str string) map[string]bool {
	nMap := make(map[string]bool)
	for _, wN := range strings.Split(str, " ") {
		if wN != "" {
			nMap[wN] = true
		}
	}
	return nMap
}

func solution1()  {
	cards := readLines("./input-1.txt");
	finalScore := 0;
	totalCards := 0
	cardCopiesMap := make(map[int]int)

	
	for i, l := range cards {
		cardCopiesMap[i] += 1;
		l = strings.Split(l, ":")[1];
		card := strings.Split(l, "|");
		
		card[0] = strings.TrimSpace(card[0]);
		card[1] = strings.TrimSpace(card[1]);
		
		
		wNumMap := getNumbersMap(card[0])
		numsMap := getNumbersMap(card[1])
	
		score := 0;
		for n := range numsMap {
			if wNumMap[n] {
				score += 1
			}
		}
		
		pow := score - 1  
		finalScore += int(math.Pow(float64(2), float64(pow)))

		for j := 0; j < score; j++ {
			cardCopiesMap[i + j + 1] += cardCopiesMap[i]
		}
	}

	for _, v := range cardCopiesMap {
		totalCards += v
	}
	
	// Solution - 1
	fmt.Println(finalScore)
	// Solution - 2
	fmt.Println(totalCards)
}

func main()  {
	solution1()
}