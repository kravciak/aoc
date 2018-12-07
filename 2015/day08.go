package main

import (
	"bufio"
	"fmt"
	"os"
)

func encode(in string) string {
	out := "\""
	for _, c := range in {
		switch c {
		case '"':
			out += "\\\""
		case '\\':
			out += "\\\\"
		default:
			out += string(c)
		}
	}
	return out + "\""
}

func decode(in string) (out string) {
	for i := 1; i < len(in)-1; i++ {
		if in[i] == '\\' {
			switch in[i+1] {
			case '"':
				out += "\""
			case '\\':
				out += "\\"
			case 'x':
				out += "%"
				i += 2
			}
			i++
		} else {
			out += string(in[i])
		}
	}
	return
}

func main() {
	f, err := os.Open("day08.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var enc, dec int
	for scanner.Scan() {
		line := scanner.Text()
		lenc := encode(line)
		ldec := decode(line)
		enc += len(lenc) - len(line)
		dec += len(line) - len(ldec)
	}
	fmt.Printf("dec: %d, enc: %d", dec, enc)
}
