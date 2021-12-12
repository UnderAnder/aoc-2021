package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

type graph struct {
	vertices []*vertex
}

type vertex struct {
	key      string
	adjacent []*vertex
	visited  bool
}

func (g *graph) addVertex(k string) {
	if g.contains(k) {
		return
	}
	g.vertices = append(g.vertices, &vertex{key: k})
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

func (g *graph) getVertex(key string) (*vertex, error) {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i], nil
		}
	}
	return &vertex{}, fmt.Errorf("vertex %v do not exist", key)
}

func (g *graph) contains(k string) bool {
	for _, v := range g.vertices {
		if k == v.key {
			return true
		}
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
		fmt.Printf("\nVertex %s : ", v.key)
		for _, vv := range v.adjacent {
			fmt.Printf(" %v ", vv.key)
		}
	}
	fmt.Println()
}

func (g *graph) traverse(start *vertex) int {
	var cnt int
	if isLower(start.key) {
		start.visited = true
	}
	if start.key == "end" {
		cnt++
	}
	for _, v := range start.adjacent {
		if v.visited {
			continue
		}
		v.visited = false
		cnt += g.traverse(v)
	}
	start.visited = false
	return cnt
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func solve(lines []string) int {
	var answer int
	g := graph{}
	for _, line := range lines {
		s := strings.Split(line, "-")
		g.addVertex(s[0])
		g.addVertex(s[1])
		g.addEdge(s[0], s[1])
	}
	g.String()
	vert, err := g.getVertex("start")
	if err != nil {
		return 0
	}
	answer = g.traverse(vert)
	return answer
}

func parse(file io.Reader) []string {
	p := make([]string, 0, 20)
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
	a1 := solve(ps)
	fmt.Printf("First answer: %d\n", a1)
}
