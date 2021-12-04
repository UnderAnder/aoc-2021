package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const size = 5

type cell struct {
	num  int
	mark bool
}
type board struct {
	cells [size][size]cell
}

func newBoard(nums []int) *board {
	b := new(board)
	k := 0
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[i]); j++ {
			b.cells[i][j].num = nums[k]
			k++
		}
	}
	return b
}

func (b *board) mark(num int) (bool, int) {
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[i]); j++ {
			if b.cells[i][j].num == num {
				b.cells[i][j].mark = true
				win, score := b.check(num, i, j)
				if win {
					return true, score
				}
			}
		}
	}
	return false, -1
}

func (b *board) check(num, i, j int) (bool, int) {
	win := true
	score := 0
	for k := 0; k < len(b.cells[i]); k++ {
		win = win && b.cells[i][k].mark
	}
	if !win {
		win = true
		for k := 0; k < len(b.cells); k++ {
			win = win && b.cells[k][j].mark
		}
	}
	if win {
		for i = 0; i < len(b.cells); i++ {
			for j = 0; j < len(b.cells[i]); j++ {
				if !b.cells[i][j].mark {
					score += b.cells[i][j].num
				}
			}
		}
		score *= num
	}
	return win, score
}

func part1(file io.Reader) int {
	var draws, boardNums []int
	var boards []*board
	var win bool
	var answer int

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			for _, v := range strings.Split(scanner.Text(), ",") {
				num, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalf("can't parse draw")
				}
				draws = append(draws, num)
			}
			continue
		}
		if scanner.Text() == "" {
			if len(boardNums) == 0 {
				continue
			}
			boards = append(boards, newBoard(boardNums))
			boardNums = make([]int, 0, size*size)
			continue
		}

		for _, v := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("can't parse board")
			}
			boardNums = append(boardNums, num)
		}
	}
	boards = append(boards, newBoard(boardNums))

	for _, draw := range draws {
		for _, b := range boards {
			win, answer = b.mark(draw)
			if win {
				return answer
			}
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf1 := &bytes.Buffer{}
	buf2 := io.TeeReader(file, buf1)
	fmt.Printf("First answer: %d\n", part1(buf2))
}
