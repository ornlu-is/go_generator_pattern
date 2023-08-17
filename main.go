package main

import "fmt"

func linearSpaceGenerator(start, end, step float64) chan float64 {
	ch := make(chan float64)

	go func(ch chan float64) {
		defer close(ch)

		for n := start; n < end; n += step {
			ch <- n
		}
	}(ch)

	return ch
}

func main() {
	// without generator pattern
	step := 0.1
	start := 0.2
	end := 2.
	for i := start; i < end; i += step {
		fmt.Printf("i: %f\n", i)
	}

	// with generator pattern
	for i := range linearSpaceGenerator(0.1, 2., 0.1) {
		fmt.Printf("i: %f\n", i)
	}
}
