package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	var p, c int
	for scanner.Scan() {
		p++
		switch scanner.Text() {
		case "(":
			c++
		case ")":
			c--
		}
		if c == -1 {
			fmt.Println("Basement:", p)
		}
	}
	fmt.Println("Floor:", c)
}
