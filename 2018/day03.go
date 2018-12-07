package main

import (
	"github.com/kravciak/aoc/helpers"
	"fmt"
	"strings"
	"strconv"
)

type Area struct {
	id int
	x, y int
	w, h int
}

var fabric [1000][1000]int

// #1 @ 896,863: 29x19
func parseLine(line string) Area {
	l := strings.Fields(line)
	id, _ := strconv.Atoi(strings.TrimPrefix(l[0], "#"))
	xy := strings.Split(strings.TrimSuffix(l[2], ":"), ",")
	hw := strings.Split(l[3], "x")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	w, _ := strconv.Atoi(hw[0])
	h, _ := strconv.Atoi(hw[1])
	return Area{id, x, y, w, h}
}

func soloArea(area Area) bool {
	for i := area.x; i < area.x + area.w; i++ {
		for j := area.y; j < area.y + area.h; j++ {
			if fabric[i][j] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	lines := helpers.ReadLines("day03.txt")
	var areas []Area

	for _, line := range lines {
		areas = append(areas, parseLine(line))
	}

	for _, area := range areas {
		for i := area.x; i < area.x + area.w; i++ {
			for j := area.y; j < area.y + area.h; j++ {
				fabric[i][j]++
			}
		}
	}

	var total int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				total++
			}
		}
	}
	fmt.Println(total)

	for _, area := range areas {
		if soloArea(area) {
			fmt.Println(area.id)
			break
		}
	}
}
