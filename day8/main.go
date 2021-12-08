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
	ins := make([][]string, 0)
	outs := make([][]string, 0)

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
	var three, five, nine string
	one := in[oneSeg][0]
	seven := in[sevenSeg][0]
	seg1 := diff(in[3][0], one)
	for i, v := range in[5] {
		d := diff(v, seven)
		if len(d) == 2 {
			three = in[5][i]
		}
	}
	for i, v := range in[6] {
		d := diff(v, three)
		if len(d) == 1 {
			nine = in[6][i]
		}
	}
	for i, v := range in[5] {
		d := diff(nine, v)
		if len(d) == 1 && v != three {
			five = in[5][i]
		}
	}
	seg2 := diff(in[fourSeg][0], three)
	seg3 := diff(in[fourSeg][0], five)
	seg4 := diff(diff(in[fourSeg][0], in[3][0]), seg2)
	seg5 := diff(diff(in[eightSeg][0], five), seg3)
	seg6 := diff(in[oneSeg][0], seg3)
	seg7 := diff(diff(three, in[3][0]), seg4)
	return dTable{
		segments: map[rune]int{
			rune(seg1[0]): 1,
			rune(seg2[0]): 2,
			rune(seg3[0]): 3,
			rune(seg4[0]): 4,
			rune(seg5[0]): 5,
			rune(seg6[0]): 6,
			rune(seg7[0]): 7,
		},
	}
}

func diff(a, b string) string {
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
	return string(d)
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
