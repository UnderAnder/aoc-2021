package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

func solve(lines []string) (int, int) {
	var answer1 int
	illigalScore := map[rune]int{
		41:  3,     // )
		93:  57,    // ]
		125: 1197,  // }
		62:  25137, // >
	}
	completeScore := map[rune]int{
		41:  1, // )
		93:  2, // ]
		125: 3, // }
		62:  4, // >
	}
	pairTable := map[rune]rune{
		40:  41,  // ()
		91:  93,  // []
		123: 125, // {}
		60:  62,  // <>
	}
	reversTable := map[rune]rune{
		41:  40,  // )(
		93:  91,  // ][
		125: 123, // }{
		62:  60,  // ><
	}

	illigals := make([]rune, 0)
	completes := make([]int, 0)
lines:
	for _, line := range lines {
		tmp := make([]rune, 0, 20)
		for _, char := range line {
			_, ok := pairTable[char]
			if ok {
				tmp = append(tmp, char)
			}
			r, ok := reversTable[char]
			if ok {
				// part1
				if tmp[len(tmp)-1] != r {
					illigals = append(illigals, char)
					continue lines
				}
				tmp = tmp[:len(tmp)-1]
			}
		}
		// part2
		score := 0
		for i := len(tmp) - 1; i >= 0; i-- {
			r := pairTable[tmp[i]]
			score *= 5
			score += completeScore[r]
		}
		completes = append(completes, score)
	}

	for _, char := range illigals {
		answer1 += illigalScore[char]
	}
	sort.Ints(completes)
	answer2 := completes[len(completes)/2]
	return answer1, answer2
}

func parse(file io.Reader) []string {
	p := make([]string, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		p = append(p, s)
	}
	return p
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	ps := parse(file)
	a1, a2 := solve(ps)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
