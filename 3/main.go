package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var counter = make(map[int]map[string]int)
	for input.Scan() {
		line := input.Text()
		for k, v := range line {
			if _, ok := counter[k]; !ok {
				counter[k] = map[string]int{"0": 0, "1": 0}
			}
			val := counter[k]
			val[string(v)] += 1
		}
	}

	fmt.Printf("Simple: %d\n", simple(counter))
}

func simple(counter map[int]map[string]int) int {
	keys := make([]int, 0, len(counter))
	for k := range counter {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	gamma := make([]int, len(counter))
	epsilon := make([]int, len(counter))
	for k := range keys {
		if counter[k]["0"] > counter[k]["1"] {
			gamma[k] = 0
			epsilon[k] = 1
		} else {
			gamma[k] = 1
			epsilon[k] = 0
		}
	}
	return binaryToInt(gamma) * binaryToInt(epsilon)
}

func binaryToInt(binary []int) int {
	result := 0
	pow := len(binary) - 1
	for i := 0; i < len(binary); i++ {
		result += int(math.Pow(2, float64(pow))) * binary[i]
		pow -= 1
	}
	return result
}
