package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	// num of segments
	oneSeg   = 2
	fourSeg  = 4
	sevenSeg = 3
	eightSeg = 7
)

type dTable struct {
	segments map[rune]int
}

func parse(file io.Reader) ([][]string, [][]string) {
	ins := make([][]string, 0, 10)
	outs := make([][]string, 0, 4)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		init := strings.Split(scanner.Text(), " | ")
		in := strings.Split(init[0], " ")
		out := strings.Split(init[1], " ")
		ins = append(ins, in)
		outs = append(outs, out)
	}

	return ins, outs
}

func uniqSegments(v int) bool {
	if v == oneSeg || v == fourSeg || v == sevenSeg || v == eightSeg {
		return true
	}
	return false
}

func decode(in map[int][]string) dTable {
	cnt := make(map[rune]int)
	segments := make(map[rune]int)
	for _, v := range in {
		for _, s := range v {
			for _, r := range s {
				cnt[r]++
			}
		}
	}
	/*
			 1111
			2    3
			2    3
			 4444
			5    6
			5    6
			 7777
		segments 2, 5, 6 appear a unique number of times
	*/
	for r, c := range cnt {
		switch c {
		case 4:
			segments[r] = 5
		case 6:
			segments[r] = 2
		case 7:
			if len(diff(string(r), in[fourSeg][0])) == 1 {
				segments[r] = 7
				continue
			}
			segments[r] = 4

		case 8:
			if len(diff(string(r), in[oneSeg][0])) == 1 {
				segments[r] = 1
				continue
			}
			segments[r] = 3
		case 9:
			segments[r] = 6
		}
	}
	return dTable{segments: segments}
}

func diff(a, b string) []rune {
	m := make(map[rune]struct{}, len(b))
	for _, x := range b {
		m[x] = struct{}{}
	}
	var d []rune
	for _, x := range a {
		if _, ok := m[x]; !ok {
			d = append(d, x)
		}
	}
	return d
}

func pow(i, j int) int {
	if j == 0 {
		return 1
	}
	result := i
	for k := 2; k <= j; k++ {
		result *= i
	}
	return result
}

func part2(ins, outs [][]string) int {
	var answer int
	check := map[int][7]int{
		0: {1, 1, 1, 0, 1, 1, 1},
		1: {0, 0, 1, 0, 0, 1, 0},
		2: {1, 0, 1, 1, 1, 0, 1},
		3: {1, 0, 1, 1, 0, 1, 1},
		4: {0, 1, 1, 1, 0, 1, 0},
		5: {1, 1, 0, 1, 0, 1, 1},
		6: {1, 1, 0, 1, 1, 1, 1},
		7: {1, 0, 1, 0, 0, 1, 0},
		8: {1, 1, 1, 1, 1, 1, 1},
		9: {1, 1, 1, 1, 0, 1, 1},
	}
	for i, in := range ins {
		d := make(map[int][]string)
		for _, v := range in {
			d[len(v)] = append(d[len(v)], v)
		}
		dec := decode(d)
		o := outs[i]
		j := len(o)
		var num int
		for _, v := range o {
			var tmp [7]int
			for _, r := range v {
				pos := dec.segments[r]
				tmp[pos-1] = 1
			}
			for k, ints := range check {
				if ints == tmp {
					j--
					num += k * pow(10, j)
				}
			}
		}
		answer += num
	}
	return answer
}

func part1(outs [][]string) int {
	var answer int
	for _, out := range outs {
		for _, v := range out {
			if uniqSegments(len(v)) {
				answer++
			}
		}
	}
	return answer
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	ins, outs := parse(file)
	a1 := part1(outs)
	a2 := part2(ins, outs)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
