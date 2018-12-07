package main

import (
	"github.com/kravciak/aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

import "container/ring"

type Person struct {
	name string
	wishes []Wish
//	left, right string
//	happy map[string] int
}

type Wish struct {
	name string
	wname string
	happy int
}

func (w Wish) String() string {
	return fmt.Sprintf("(%s -> %s = %d)", w.name, w.wname, w.happy)
	return w.name
}

func NewWish(line string) Wish {
	l := strings.Fields(line)
	p, q := l[0], strings.TrimSuffix(l[10], ".")
	h, _ := strconv.Atoi(l[3])
	if l[2] == "lose" { h = -h }
	return Wish{p,q,h}
}

func calcHappy() {

}

func printRing(r *ring.Ring) {
	r.Do(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println()
}

func main() {
	lines := helpers.ReadLines("day13s.txt")
	
	people := make(map[string]bool)
	wishes := make([]Wish, len(lines))
	for i, line := range lines {
		wish := NewWish(line)
		wishes[i] = wish
		people[wish.name] = true
	}

	fmt.Println(people)

	r := ring.New(len(wishes))
	for i, wish := range wishes {
		r.Value = wish
		r = r.Next()
		printRing(r)
		
		if i == 3 {break
		}
	}
}
