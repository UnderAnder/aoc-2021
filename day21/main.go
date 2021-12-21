package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc/pkg/helpers"
)

type player struct {
	pos, score int
}

func (p *player) move(die int) {
	p.pos += die
	p.pos = p.pos % 10
	if p.pos == 0 {
		p.pos = 10
	}
	p.score += p.pos
}

func deterministicDie() func() int {
	start, stop := 1, 101
	return func() int {
		res := 0
		end := start + 3
		for ; start < end; start++ {
			if start == stop {
				end = end - start + 1
				start = 1
			}
			res += start
		}
		return res
	}
}

func solve(file io.Reader) (answer1 int) {
	players := parse(file)
	gen := deterministicDie()
	rolls := 0
	loser := -1
	for i := 0; i < 2; i = (i + 1) % 2 {
		die := gen()
		rolls += 3
		players[i].move(die)
		if players[i].score >= 1000 {
			loser = (i + 1) % 2
			break
		}
	}
	gen = nil
	answer1 = players[loser].score * rolls
	return answer1
}

func answer(m map[string]int) int {
	uniqChars := make(map[uint8]int)
	for k, v := range m {
		uniqChars[k[1]] += v
	}
	min, max := 1<<(strconv.IntSize-1)-1, 0
	for _, v := range uniqChars {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func parse(file io.Reader) []player {
	players := make([]player, 0, 2)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pos := strings.Split(scanner.Text(), " ")
		players = append(players, player{helpers.MustAtoi(pos[4]), 0})
	}
	return players
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	a1 := solve(file)
	fmt.Println("First answer: ", a1)
	// fmt.Println("Second answer: ", a2)
}
