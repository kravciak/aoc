package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(a []string, i1, i2 int) {
	a[i1], a[i2] = a[i2], a[i1]
}

func heap(a []string, size int) {
	if size == 1 {
		total := 0
		for i := 0; i < len(a)-1; i++ {
			total += get(a[i], a[i+1])
		}
		if min == 0 || min > total {
			min = total
		}
		if max < total {
			max = total
		}
		return
	}
	for i := 0; i < size; i++ {
		heap(a, size-1)
		if size%2 == 1 {
			swap(a, 0, size-1)
		} else {
			swap(a, i, size-1)
		}
	}
}

type Path struct {
	t1, t2 string
	dst    int
}

var min, max int
var towns = make(map[string]bool)
var paths = make([]Path, 0)

func get(t1, t2 string) int {
	for _, p := range paths {
		if p.t1 == t1 && p.t2 == t2 || p.t1 == t2 && p.t2 == t1 {
			return p.dst
		}
	}
	return 0
}

func main() {
	f, err := os.Open("day09.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// fmt.Println(line, dst)
		line := strings.Fields(scanner.Text())
		t1, t2 := string(line[0]), string(line[2])
		dst, _ := strconv.Atoi(line[4])

		towns[t1] = true
		towns[t2] = true
		paths = append(paths, Path{t1, t2, dst})
	}

	n := 0
	towns_slice := make([]string, len(towns))
	for t := range towns {
		towns_slice[n] = t
		n++
	}
	heap(towns_slice, n)
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
