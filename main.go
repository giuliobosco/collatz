package main

import "fmt"

func collatz(a, b, c, d int64) int64 {
	if a == 4 && b == 2 && c == 1 {
		return d
	}

	if c%2 == 0 {
		return collatz(b, c, c/2, d+1)
	} else {
		return collatz(b, c, 3*c+1, d+1)
	}
}

func main() {
	var number, cycles, i int64

	for i = 1; i < 50000000; i++ {
		if a := collatz(0, 0, i, 0); cycles < a {
			number = i
			cycles = a
		}
	}

	fmt.Printf("\n\nBigger\nnumber: %d, cycles: %d\n", number, cycles)
}
