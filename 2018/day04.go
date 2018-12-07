package main

import (
	"github.com/kravciak/aoc/helpers"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

var guards = make(map[int][]int)

func main() {
	lines := helpers.ReadLines("day04.txt")
	sort.Strings(lines)

	var g, td int
	for _, line := range lines {
		t , _ := strconv.Atoi(line[15:17])
		switch line[19:] {
			case "falls asleep":
				td = t
			case "wakes up":
				for ; td < t; td++ {
					guards[g][td]++
				}
			default:
				g, _ = strconv.Atoi(strings.Fields(line[26:])[0])
				if _, ok := guards[g]; !ok {
					guards[g] = make([]int, 60)
				}
		}
	}

	var sgid, smax int
	for guard, minutes := range guards {
		sleep := 0
		for _, m := range minutes { sleep += m }
		if sleep > smax {
			smax = sleep
			sgid = guard
		}
	}

	var maxwin, wintime int
	for i := 0; i < 60; i++ {
		if guards[sgid][i] > maxwin {
			maxwin = guards[sgid][i]
			wintime = i
		}
	}
	fmt.Println(sgid, wintime, sgid*wintime)

	var fgid, fmin, fmax int
	for guard, minutes := range guards {
		for i, m := range minutes { 
			if m > fmax {
				fmax = m
				fmin = i
				fgid = guard
			}
		}
	}
	fmt.Println(fgid, fmin, fgid * fmin)

}
