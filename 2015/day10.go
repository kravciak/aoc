package main

import (
	"fmt"
	"strconv"
	"strings"
)

func lookAndSay(in string) string {
	var c int
	var b byte
	var sb strings.Builder

	for i := 0; i < len(in); i++ {
		if b == in[i] {
			c++
		} else {
			if c > 0 {
				sb.WriteString(strconv.Itoa(c))
				sb.WriteByte(b)
			}
			b = in[i]
			c = 1
		}
	}
	sb.WriteString(strconv.Itoa(c))
	sb.WriteByte(b)
	return sb.String()
}

func main() {
	input := "1113122113"
	for i := 0; i < 50; i++ {
		input = lookAndSay(input)
	}
	fmt.Println(len(input))
}
