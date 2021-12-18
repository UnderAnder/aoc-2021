package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type pair struct {
	x, y, parent *pair
	num          int
	isNum        bool
}

// parse string recursive by 1 symbol
func parse(parent *pair, s string) (*pair, string) {
	p := pair{}
	p.parent = parent
	if s[0] == '[' {
		p.isNum = false
		p.x, s = parse(&p, s[1:])
		p.y, s = parse(&p, s[1:])
		return &p, s[1:]
	}
	p.isNum = true
	p.num = int(s[0] - '0')
	return &p, s[1:]
}

func (p *pair) reduce() {
	p.x.parent = p
	p.y.parent = p
	for {
		if ok := p.explode(0, false); ok {
			continue
		}
		if ok := p.split(); ok {
			continue
		}
		break
	}
}

func (p *pair) explode(depth int, isExplode bool) bool {
	depth++
	if p.isNum {
		return isExplode
	}
	if depth > 4 && p.x.isNum && p.y.isNum {
		x, y := p.x.num, p.y.num
		p.addLeft(x)
		p.addRight(y)
		p.x, p.y = nil, nil

		p.isNum = true
		p.num = 0
		return true
	}
	isExplode = p.x.explode(depth, isExplode)
	if !isExplode {
		isExplode = p.y.explode(depth, isExplode)
	}
	return isExplode

}

func (p *pair) addRight(n int) {
	if p.parent == nil {
		return
	}
	if p.parent.y != nil && p.parent.y != p {
		lt := p.parent.y
		for lt.x != nil {
			lt = lt.x
		}
		lt.num += n
	} else if p.parent.y != nil {
		p.parent.addRight(n)
	} else {
		p.parent.y.num += n
	}
}

func (p *pair) addLeft(n int) {
	if p.parent == nil {
		return
	}
	if p.parent.x != nil && p.parent.x != p {
		rt := p.parent.x
		for rt.y != nil {
			rt = rt.y
		}
		rt.num += n
	} else if p.parent.x != nil {
		p.parent.addLeft(n)
	} else {
		p.parent.x.num += n
	}
}

func (p *pair) split() bool {
	if p.isNum {
		if p.num > 9 {
			*p = pair{
				isNum:  false,
				x:      &pair{isNum: true, num: p.num / 2, parent: p},
				y:      &pair{isNum: true, num: p.num - p.num/2, parent: p},
				parent: p.parent,
			}
			return true
		}
		return false
	}
	return p.x.split() || p.y.split()
}

func (p *pair) magnitude() int {
	if p.isNum {
		return p.num
	}
	n := 3 * p.x.magnitude()
	m := 2 * p.y.magnitude()
	return n + m
}

func (p *pair) String() string {
	if p.isNum {
		return fmt.Sprintf("%d", p.num)
	}
	return fmt.Sprintf("[%s,%s]", p.x, p.y)
}

func part1(file io.Reader) (answer1 int) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	p, _ := parse(nil, scanner.Text())
	for scanner.Scan() {
		p2, _ := parse(nil, scanner.Text())
		log.Println(p.String())
		log.Println(p2.String())
		p = &pair{
			x:      p,
			y:      p2,
			isNum:  false,
			parent: nil,
		}
		p.reduce()
	}
	answer1 = p.magnitude()
	return answer1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	a1 := part1(file)
	fmt.Printf("First answer: %d\n", a1)
}
