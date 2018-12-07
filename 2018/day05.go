package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isPair(a, b byte) bool {
	var d byte
	if a > b {
		d = a - b
	} else {
		d = b - a
	}
	return d == 32
}

func isEqual(a, b byte) bool {
	return strings.EqualFold(string(a), string(b))
}

func prevI(i int) int {
	for i--; i >= 0; i-- {
		if polymer[i] != 0 {
			return i
		}
	}
	return -1
}

func nextJ(j int) int {
	for j++; j < len(polymer); j++ {
		if polymer[j] != 0 {
			return j
		}
	}
	return len(polymer)
}

func initIJ() (ii, jj int){
	ii = -1
	for i, p := range polymer {
		if p != 0 {
			if ii == -1 {
				ii = i
			} else {
				jj = i
				break
			}
		}
	}
	return
}

func react(ignore byte) (units int) {
	n := len(polymer)

	for i := 0; i < n; i++ {
		if isEqual(polymer[i], ignore) {
			polymer[i] = 0
		}
	}

	for i, j := initIJ(); j < n; j = nextJ(j) {
		if isPair(polymer[i], polymer[j]) {
			polymer[i], polymer[j] = 0, 0
			if pi := prevI(i); pi != -1 {
				i = pi
			} else {
				j = nextJ(j)
				i = j
			}
		} else {
			i = j
		}
	}

	for _, u := range polymer {
		if u != 0 {
			units++
		}
	}
	return
}

var polymer []byte

func main() {
	data, _ := ioutil.ReadFile("day05.txt")
//	data = []byte("dabAcCaCBAcCcaDA")
	polymer = make([]byte, len(data))

	min := len(data)
	for i := byte('a'); i < byte('z'); i++ {
		copy(polymer, data)
		r := react(i)
		if r < min {
			min = r
		}
	}
	fmt.Println(min)
}
