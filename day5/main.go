package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const size = 5

type point struct {
	x, y int
}

type entry struct {
	p1, p2 point
}

// func newBoard(nums []int) *board {
// 	b := new(board)
// 	k := 0
// 	for i := 0; i < len(b.cells); i++ {
// 		for j := 0; j < len(b.cells[i]); j++ {
// 			b.cells[i][j].num = nums[k]
// 			k++
// 		}
// 	}
// 	return b
// }

// func (b *board) mark(num int) (bool, int) {
// 	for i := 0; i < len(b.cells); i++ {
// 		for j := 0; j < len(b.cells[i]); j++ {
// 			if b.cells[i][j].num == num {
// 				b.cells[i][j].mark = true
// 				win, score := b.check(num, i, j)
// 				if win {
// 					return true, score
// 				}
// 			}
// 		}
// 	}
// 	return false, -1
// }

// func (b *board) check(num, i, j int) (bool, int) {
// 	win := true
// 	score := 0
// 	for k := 0; k < len(b.cells[i]); k++ {
// 		win = win && b.cells[i][k].mark
// 	}
// 	if !win {
// 		win = true
// 		for k := 0; k < len(b.cells); k++ {
// 			win = win && b.cells[k][j].mark
// 		}
// 	}
// 	if win {
// 		for i = 0; i < len(b.cells); i++ {
// 			for j = 0; j < len(b.cells[i]); j++ {
// 				if !b.cells[i][j].mark {
// 					score += b.cells[i][j].num
// 				}
// 			}
// 		}
// 		score *= num
// 	}
// 	return win, score
// }
//
// func solve(file io.Reader) (int, int) {
// 	var boards []*board
// 	var win bool
// 	var answer, answer1, answer2, winCount int
// 	boardNums := make([]int, 0, size*size)
//
// 	scanner := bufio.NewScanner(file)
//
// 	// draws
// 	scanner.Scan()
// 	drawsStr := strings.Split(scanner.Text(), ",")
// 	draws := make([]int, 0, len(drawsStr))
// 	for _, v := range drawsStr {
// 		num, err := strconv.Atoi(v)
// 		if err != nil {
// 			log.Fatalf("can't parse draw")
// 		}
// 		draws = append(draws, num)
// 	}
//
// 	// boards
// 	for i := 0; scanner.Scan(); i++ {
// 		for _, v := range strings.Fields(scanner.Text()) {
// 			num, err := strconv.Atoi(v)
// 			if err != nil {
// 				log.Fatalf("can't parse board")
// 			}
// 			boardNums = append(boardNums, num)
// 		}
// 		if len(boardNums) == size*size {
// 			boards = append(boards, newBoard(boardNums))
// 			boardNums = make([]int, 0, size*size)
// 		}
// 	}
//
// 	// solve
// 	for _, draw := range draws {
// 		for i := range boards {
// 			if boards[i] == nil {
// 				continue
// 			}
// 			if win, answer = boards[i].mark(draw); win {
// 				winCount++
// 				boards[i] = nil
// 				if winCount == 1 {
// 					answer1 = answer
// 				}
// 				if winCount == len(boards) {
// 					answer2 = answer
// 				}
// 			}
// 		}
// 	}
// 	return answer1, answer2
// }

func parse(file io.Reader) []entry {
	entries := make([]entry, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		point1 := strings.Split(points[0], ",")
		p1 := point{
			x: mustAtoi(point1[0]),
			y: mustAtoi(point1[1]),
		}
		point2 := strings.Split(points[1], ",")
		p2 := point{
			x: mustAtoi(point2[0]),
			y: mustAtoi(point2[1]),
		}
		entries = append(entries, entry{p1: p1, p2: p2})
	}
	return entries
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("can't parse int from input")
	}
	return i
}

func printField(field map[point]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if v, ok := field[point{x: j, y: i}]; ok {
				fmt.Print(v)
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}
}

func part1(entries []entry) int {
	var answer int
	field := make(map[point]int)

	for _, e := range entries {
		var start, end int
		if e.p1.x == e.p2.x {
			start, end = e.p1.y, e.p2.y
			if e.p1.y > e.p2.y {
				start, end = end, start
			}
			for i := start; i <= end; i++ {
				field[point{e.p2.x, i}]++
			}
		}
		if e.p1.y == e.p2.y {
			start, end = e.p1.x, e.p2.x
			if e.p1.x > e.p2.x {
				start, end = end, start
			}
			for i := start; i <= end; i++ {
				field[point{i, e.p2.y}]++
			}
		}
	}
	for _, i := range field {
		if i > 1 {
			answer++
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
	entries := parse(file)

	fmt.Printf("First answer: %d\n", part1(entries))
}
