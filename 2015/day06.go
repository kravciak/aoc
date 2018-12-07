package main

import (
	"bufio"
	"fmt"
	"os"
)

import (
	"strconv"
	"strings"
)

type Grid [1000][1000]int

type Area struct {
	x, y int
	l, h int
}

func (g *Grid) turn(a Area, action string) {
	fmt.Println(a, action)
	for i := a.x; i <= a.l; i++ {
		for j := a.y; j <= a.h; j++ {
			switch action {
			case "on":
				g[i][j]++
			case "off":
				if g[i][j] > 0 {
					g[i][j]--
				}
			case "toggle":
				g[i][j] += 2
			}
		}
	}
}

func (g *Grid) count() int {
	var count int
	for _, x := range g {
		for _, y := range x {
			count += y
		}
	}
	return count
}

func main() {
	f, err := os.Open("day06.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var g Grid
	var a Area
	for scanner.Scan() {
		word := strings.TrimPrefix(scanner.Text(), "turn ")
		todo := strings.Fields(word)

		// Parse params into Area
		xy := strings.Split(todo[1], ",")
		a.x, _ = strconv.Atoi(xy[0])
		a.y, _ = strconv.Atoi(xy[1])
		xy = strings.Split(todo[3], ",")
		a.l, _ = strconv.Atoi(xy[0])
		a.h, _ = strconv.Atoi(xy[1])

		g.turn(a, todo[0])
	}
	fmt.Println(g.count())
}
