package main

import (
	"fmt"
//	"strconv"
	"strings"
)

func inc(pw string) string {
	for i := len(pw) - 1; i >= 0; i-- {
		if pw[i] == 'z' {
			pw = pw[:i] + "a" + pw[i+1:]
		} else {
			pw = pw[:i] + string(pw[i]+1) + pw[i+1:]
			break
		}
	}
	return pw
}

// Passwords must include one increasing straight of at least three letters
// like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
func rule1(pw string) bool {
	for i := 0; i < len(pw)-2; i++ {
		if pw[i] == pw[i+1]-1 && pw[i] == pw[i+2]-2 {
			return true
		}
	}
	return false
}

// Passwords may not contain the letters i, o, or l
func rule2(pw string) bool {
	return !strings.ContainsAny(pw, "iol")
}

// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
func rule3(pw string) bool {
	var p bool
	for i := 0; i < len(pw)-1; i++ {
		if pw[i] == pw[i+1] {
			if p {
				return true
			} else {
				p = true
				i++
			}
		}
	}
	return false
}

func rules(pw string) bool {
	return rule1(pw) && rule2(pw) && rule3(pw)
}

func nextPW(pw string) string {
	pw = inc(pw)
	for i := 0; rules(pw) == false; i++ {
		pw = inc(pw)
		if i == 1000000 {
			fmt.Println("NOT FOUND")
			return ""
		}
	}
	return pw
}

func main() {
	input := "hepxcrrq"

	input = nextPW(input)
	fmt.Println(input)

	input = nextPW(input)
	fmt.Println(input)
}
