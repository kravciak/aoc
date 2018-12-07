package main

import (
	"github.com/kravciak/aoc/helpers"
	"fmt"
	"strings"
	"strconv"
//	"os"
)

type Point struct {
	x, y int
}

var points []Point

func parseLine(line string) Point {
	l := strings.Split(line, ", ")
	x, _ := strconv.Atoi(l[0])
	y, _ := strconv.Atoi(l[1])
	return Point{x,y}
}

func findInfinites() []Point {
	var a,b Point
	c := Point{999, 999}
	d := Point{999, 999}

	for _, p := range points {
		if p.x > a.x { a = p }
		if p.y > b.y { b = p }
		if p.x < c.x { c = p }
		if p.y < d.y { d = p }
	}
	panic("FIXME")
	return []Point{a,b,c,d}
}

// var grid = make(map[Point]Point)
var grid [1000][1000]Point

func main() {
	lines := helpers.ReadLines("day06s.txt")
	points = make([]Point, len(lines))

//	fmt.Println(os.Args)

	for i, line := range lines {
		points[i] = parseLine(line)
	}
	inf := findInfinites()

	fmt.Println(inf)
//	[][]int
	
}
