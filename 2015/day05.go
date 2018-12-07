package main

import (
	"bufio"
	"fmt"
	"os"
)

import (
	"strings"
)

func vovels(word string) bool {
	var count int
	for _, c := range word {
		if strings.ContainsRune("aeiou", c) {
			count++
		}
		if count == 3 {
			return true
		}
	}
	return false
}

func twice(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		if word[i] == word[i+1] {
			return true
		}
	}
	return false
}

func badboys(word string) bool {
	bad := [...]string{"ab", "cd", "pq", "xy"}
	for _, b := range bad {
		if strings.Contains(word, b) {
			return false
		}
	}
	return true
}

/*
It contains a pair of any two letters that appears at least twice
in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa),
but not like aaa (aa, but it overlaps).
*/
func pair(word string) bool {
	for i := 0; i < len(word)-1; i++ {
		touple := string(word[i]) + string(word[i+1])
		fin := strings.Index(word, touple)
		lin := strings.LastIndex(word, touple)
		if lin-fin > 1 {
			return true
		}
	}
	return false
}

func repeat(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("day05.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var counter int
	for scanner.Scan() {
		word := scanner.Text()
		//		if vovels(word) && twice(word) && badboys(word) {
		//			counter++
		//		}
		fmt.Println(word)
		if pair(word) && repeat(word) {
			counter++
		}
		//break
	}
	fmt.Println(counter)
}
