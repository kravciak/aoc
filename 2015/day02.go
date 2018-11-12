package main

import (
	"bufio"
	"fmt"
	"os"
)

import (
	"sort"
	"strconv"
	"strings"
)

type Present struct {
	x, y, z int
}

func (p *Present) area() int {
	return 2*(p.x*p.y+p.x*p.z+p.y*p.z) + p.x*p.y
}

func (p *Present) ribbon() int {
	r := 2 * (p.x + p.y)
	t := p.x * p.y * p.z
	return r + t
}

func main() {
	f, err := os.Open("day02.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var paper, ribbon int
	ints := make([]int, 3, 3)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "x")
		for i, v := range vals {
			ints[i], err = strconv.Atoi(v)
		}
		sort.Ints(ints)
		p := Present{ints[0], ints[1], ints[2]}

		paper += p.area()
		ribbon += p.ribbon()
	}
	fmt.Println("Need paper:", paper)
	fmt.Println("Need ribbon:", ribbon)
}
