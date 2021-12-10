package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func solve(lines []string) int {
	var answer int
	illigalScore := map[rune]int{
		41:  3,     // )
		93:  57,    // ]
		125: 1197,  // }
		62:  25137, // >
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
	for _, line := range lines {
		tmp := make([]rune, 0, 20)
		for _, char := range line {
			_, ok := pairTable[char]
			if ok {
				tmp = append(tmp, char)
			}
			r, ok := reversTable[char]
			if ok {
				if tmp[len(tmp)-1] != r {
					illigals = append(illigals, char)
					break
				}
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	for _, char := range illigals {
		answer += illigalScore[char]
	}
	return answer
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
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	ps := parse(file)
	a1 := solve(ps)
	fmt.Printf("First answer: %d\n", a1)
}
