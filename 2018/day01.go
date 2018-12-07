package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
	file, err := os.Open("day01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var changes []int
	for scanner.Scan() {
		line := scanner.Text()
		f, _ := strconv.Atoi(line)
		changes = append(changes, f)
	}

	// Find total frequency
	var freq int
	for _, f := range changes {
		freq += f
	}
	fmt.Println(freq)

	// Find repeatign frequency
	freq = 0
	var safety int
	var reached = map[int]bool{0:true}

	for i := 0; ; i++ {
		freq += changes[i]
		if _, ok := reached[freq]; ok {
			fmt.Println("Reached twice:", freq)
			break
		} else {
			reached[freq] = true
		}
		if i == len(changes)-1 { i = -1 }
		if safety++; safety > 1000000 { break }
	}
}
