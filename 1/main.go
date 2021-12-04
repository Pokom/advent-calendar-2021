package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var readings []int
	for input.Scan() {
		current, err := strconv.Atoi(input.Text())
		if err != nil {
			continue
		}
		readings = append(readings, current)
	}
	fmt.Printf("Simple: %d\n", simple(readings))
	fmt.Printf("Three Window: %d\n", three_window(readings))
}

func simple(readings []int) int {
	increases := 0
	prev := 0
	for _, current := range readings {
		if current > prev && prev != 0 {
			increases += 1
		}
		prev = current
	}
	return increases
}

func three_window(readings []int) int {
	increases := 0
	currentWindowTotal := 0
	prevWindowTotal := 0

	for i := 0; i < len(readings)-2; i++ {
		a := readings[i]
		b := readings[i+1]
		c := readings[i+2]
		currentWindowTotal = a + b + c
		if prevWindowTotal != 0 && currentWindowTotal > prevWindowTotal {
			increases += 1
		}
		prevWindowTotal = currentWindowTotal
	}
	return increases
}
