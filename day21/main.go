package main

import (
	"fmt"
)

type state struct {
	players [2]player
	turn    int
}

type player struct {
	pos, score int
}

func (p *player) move(die int) {
	p.pos += die
	p.pos %= 10
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

func part1() (answer1 int) {
	players := [2]player{{4, 0}, {10, 0}}
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

func winCounter(state state, cache map[state][2]int) [2]int {
	if v, ok := cache[state]; ok {
		return v
	}
	var result [2]int
	for die1 := 1; die1 <= 3; die1++ {
		for die2 := 1; die2 <= 3; die2++ {
			for die3 := 1; die3 <= 3; die3++ {
				die := die1 + die2 + die3
				s := state
				s.players[s.turn].move(die)
				if s.players[s.turn].score >= 21 {
					result[s.turn]++
					continue
				}
				s.turn++
				s.turn %= 2
				win := winCounter(s, cache)
				result[0] += win[0]
				result[1] += win[1]
			}
		}
	}
	cache[state] = result
	return result
}

func part2() (answer2 int) {
	s := state{[2]player{{4, 0}, {10, 0}}, 0}
	cache := make(map[state][2]int)
	cnt := winCounter(s, cache)
	answer2 = cnt[0]
	if cnt[1] > cnt[0] {
		answer2 = cnt[1]
	}
	return answer2
}

func main() {
	a1, a2 := part1(), part2()
	fmt.Println("First answer: ", a1)
	fmt.Println("Second answer: ", a2)
}
