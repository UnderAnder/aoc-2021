package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"aoc/pkg/helpers"
)

type graph struct {
	vertices map[string]*vertex
}

type vertex struct {
	key      string
	adjacent []*vertex
	visited  int
}

func (g *graph) addVertex(k string) {
	if g.contains(k) {
		return
	}
	g.vertices[k] = &vertex{key: k}
}

func (g *graph) addEdge(key1, key2 string) {
	k1, err := g.getVertex(key1)
	if err != nil {
		log.Println(err)
		return
	}
	k2, err := g.getVertex(key2)
	if err != nil {
		log.Println(err)
		return
	}
	k1.add(k2)
	k2.add(k1)
}

func (v *vertex) add(vert *vertex) {
	if v.contains(vert.key) {
		return
	}
	v.adjacent = append(v.adjacent, vert)
}

func (g *graph) getVertex(k string) (*vertex, error) {
	if _, ok := g.vertices[k]; ok {
		return g.vertices[k], nil
	}
	return &vertex{}, fmt.Errorf("vertex %v do not exist", k)
}

func (g *graph) contains(k string) bool {
	if _, ok := g.vertices[k]; ok {
		return true
	}
	return false
}

func (v *vertex) contains(k string) bool {
	for _, val := range v.adjacent {
		if k == val.key {
			return true
		}
	}
	return false
}

func (g graph) String() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %s: ", v.key)
		for _, vv := range v.adjacent {
			fmt.Printf("%s ", vv.key)
		}
	}
	fmt.Println()
}

func (g *graph) traverse(vert *vertex, twice bool) int {
	var cnt int
	if vert.key == "end" {
		return cnt + 1
	}
	for _, v := range vert.adjacent {
		notToVisit := helpers.IsLower(v.key) && (v.key == "start" || twice) && v.visited != 0
		if notToVisit {
			continue
		}
		v.visited++
		cnt += g.traverse(v, twice || helpers.IsLower(v.key) && v.visited == 2)
		v.visited--
	}
	return cnt
}

func solve(file io.Reader) (answer1, answer2 int) {
	g := graph{vertices: make(map[string]*vertex)}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "-")
		g.addVertex(s[0])
		g.addVertex(s[1])
		g.addEdge(s[0], s[1])
	}
	g.String()
	vert, err := g.getVertex("start")
	if err != nil {
		return 0, 0
	}
	vert.visited = 1
	answer1 = g.traverse(vert, true)
	answer2 = g.traverse(vert, false)
	return answer1, answer2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	a1, a2 := solve(file)
	fmt.Printf("First answer: %d\n", a1)
	fmt.Printf("Second answer: %d\n", a2)
}
