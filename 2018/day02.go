package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func process(line string) (two bool, three bool) {
	counts := make(map[rune]int)
	for _, b := range line {
		counts[b]++
	}
	for _, c := range counts {
		switch c {
		case 2: two = true
		case 3: three = true
		}
		if two && three { break }
	}
	return
}

func diffOne(line1, line2 string) bool {
	var count int
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			count++
		}
		if count > 1 { break }
	}
	return count == 1
}

func printDiff(line1, line2 string) {
	var sb strings.Builder
	for i := 0; i < len(line1); i++ {
		if line1[i] == line2[i] {
			sb.WriteByte(line1[i])
		}
	}
	fmt.Println(sb.String())
}

func main() {
	file, err := os.Open("day02.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read file into slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var twos, threes int
	for _, line := range lines {
		two, three := process(line)
		if two { twos++ }
		if three { threes++ }
	}
	fmt.Println("Two:", twos, "Three:", threes, "Sum:", twos * threes)

	for i := 0; i < len(lines)-1; i++ {
		for j := i+1; j < len(lines); j++ {
			if diffOne(lines[i], lines[j]) {
				printDiff(lines[i], lines[j])
				return
			}
		}
	}


}
