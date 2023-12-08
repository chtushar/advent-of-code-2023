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

const (
	K5 = 6
	K4 = 5
	FH = 4
	K3 = 3
	P2 = 2
	P1 = 1
	H = 0
)

type Hand struct {
	cards string
	bid int
	kind int
}

var cards map[string]string = map[string]string{
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"T": "A",
	"J": "B",
	"Q": "C",
	"K": "D",
	"A": "E",
}

var cards2 map[string]string = map[string]string{
	"J": "*",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"T": "A",
	"Q": "C",
	"K": "D",
	"A": "E",
}

const solution = 2;

func checkKindWithJoker(hand Hand) int {
	maps := make(map[rune]int)
	for _, c := range hand.cards {
		maps[c] += 1
	}

	if maps['J'] >= 4 {
		return K5
	}

	if maps['J'] == 3 {
		if len(maps) == 2{
			return K5
		}
		return K4
	}

	if maps['J'] == 2 {
		if hand.kind == P2 {
			return K4
		}

		if hand.kind == P1 {
			return K3
		}

		if hand.kind == FH {
			return K5	
		}
	}

	if maps['J'] == 1 {
		if hand.kind == K3 {
			return K4
		}
		if hand.kind == P2 {
			return FH
		}
		if hand.kind == P1 {
			return K3
		}
		if hand.kind == K4 {
			return K5
		}
	}

	return P1
}

func checkKind(hand string) int {
	maps := make(map[rune]int)
	for _, c := range hand {
		maps[c] += 1
	}

	if len(maps) == 1 {
		return K5
	}
	
	if len(maps) == 2 {
		for _, c := range maps {
			if c == 4 {
				return K4
			}

			if c == 3 {
				return FH
			}
		}

		return FH
	}

	if len(maps) == 3 {
		for _, c := range maps {
			if c == 3 {
				return K3
			}
		}

		return P2
	}

	if len(maps) == 4 {
		return P1
	}

	return H
}

func checkPower(h1 Hand, h2 Hand) bool {
	if h1.kind == h2.kind {
		for i, c1 := range h1.cards {
			c2 := rune(h2.cards[i])
			if c2 != c1 {
				if solution == 2 {
					return cards2[string(c1)] < cards2[string(c2)]
				}
				return cards[string(c1)] < cards[string(c2)]
			}
		}
	}

	return h1.kind < h2.kind
}

func partitionHandsArray(handsArray *[]Hand, start, end int) int {
    pivot := (*handsArray)[end]
	pI := start

    for i := start; i < end; i++ {
        if checkPower((*handsArray)[i], pivot) {
            (*handsArray)[pI], (*handsArray)[i] = (*handsArray)[i], (*handsArray)[pI]
            pI++
        }
    }

    (*handsArray)[pI], (*handsArray)[end] = (*handsArray)[end], (*handsArray)[pI]

    return pI
}

func sortHands(handsArray *[]Hand, start int, end int)  {
	if start < end {
		pI  := partitionHandsArray(handsArray, start, end)
		
		sortHands(handsArray, start, pI-1)
    	sortHands(handsArray, pI+1, end)
	}
}

func main()  {
	lines := readLines("./input.txt")
	spaceRegex := regexp.MustCompile("\\s+")

	handsArray := []Hand{}

	for _, l := range lines {
		splits := spaceRegex.Split(l, -1)
		hand := splits[0]
		bid := convStrToInt(splits[1])

		kind := checkKind(hand)

		this := Hand{
			cards: hand,
			bid: bid,
			kind: kind,
		}

		if solution == 2 && strings.Contains(hand, "J") {
			newKind := checkKindWithJoker(this)
			this.kind = newKind
		}

		handsArray = append(handsArray, this)
	}

	sortHands(&handsArray, 0, len(handsArray) - 1)

	winning := 0
	for i, h := range handsArray {
		fmt.Println(h.cards, h.bid, h.kind)
		winning += (i + 1) * h.bid
	}
	fmt.Println(winning)
}