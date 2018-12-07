package main

import (
	"fmt"
	"io/ioutil"
)

type Traveler struct {
	x, y int
}

func (t *Traveler) travel(c byte) {
	switch c {
	case '>':
		t.x++
	case '<':
		t.x--
	case '^':
		t.y++
	case 'v':
		t.y--
	}
}

func main() {
	moves, err := ioutil.ReadFile("day03.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var houses = make(map[Traveler]int)
	santa := Traveler{0, 0}
	robot := Traveler{0, 0}
	houses[santa]++
	houses[robot]++

	for i, m := range moves {
		if i%2 == 0 {
			santa.travel(m)
			houses[santa]++
		} else {
			robot.travel(m)
			houses[robot]++
		}
		// fmt.Println(string(m), h, houses[h])
	}

	fmt.Println("Traveler:", len(houses))
	//fmt.Println("Together:", len(santa)+len(robot))
}
