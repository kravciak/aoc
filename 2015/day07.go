package main

import (
	"bufio"
	"fmt"
	"os"
)

import (
	"strconv"
	"strings"
)

var instr = make(map[string]string)
var wires = make(map[string]uint16)
var count int

func solveWire(wid string) uint16 {
	// Return if it's already solved
	if val, err := strconv.ParseUint(wid, 10, 16); err == nil {
		return uint16(val)
	}
	if _, ok := wires[wid]; ok {
		return wires[wid]
	}
	// Prevent eternal recursion
	if count++; count > 10000 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", count))
	}

	in := instr[wid]
	fmt.Printf("Solving: %s = %s ", wid, in)

	if strings.ContainsRune(in, ' ') {
		// gate
		in := strings.Fields(in)
		fmt.Println("(gate)")
		switch in[1] {
		case "AND":
			wires[wid] = and(solveWire(in[0]), solveWire(in[2]))
		case "OR":
			wires[wid] = or(solveWire(in[0]), solveWire(in[2]))
		case "LSHIFT":
			bits, _ := strconv.ParseUint(in[2], 10, 32)
			wires[wid] = lshift(solveWire(in[0]), uint(bits))
		case "RSHIFT":
			bits, _ := strconv.ParseUint(in[2], 10, 32)
			wires[wid] = rshift(solveWire(in[0]), uint(bits))
		default:
			wires[wid] = not(solveWire(in[1]))
		}
		// fmt.Println(wid, wires[wid])
	} else if val, err := strconv.ParseUint(in, 10, 16); err == nil {
		// value
		fmt.Println("(value)")
		wires[wid] = uint16(val)
	} else {
		// wire
		fmt.Println("(wire)")
		wires[wid] = solveWire(in)
	}
	return wires[wid]
}

func lshift(w uint16, n uint) uint16 {
	return w << n
}
func rshift(w uint16, n uint) uint16 {
	return w >> n
}
func and(w, v uint16) uint16 {
	return w & v
}
func or(w, v uint16) uint16 {
	return w | v
}
func not(w uint16) uint16 {
	return ^w
}

func main() {
	f, err := os.Open("day07.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		in := strings.Split(scanner.Text(), "->")
		in[0] = strings.TrimSpace(in[0])
		in[1] = strings.TrimSpace(in[1])
		instr[in[1]] = in[0]
	}

	wires["b"] = 46065
	fmt.Println("a:", solveWire("a"))
}
