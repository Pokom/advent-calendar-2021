package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	increases := 0
	var prev int
	var err error
	prev, err = strconv.Atoi(input.Text())
	if err != nil {
		log.Fatal(err)
	}
	for input.Scan() {
		current, err := strconv.Atoi(input.Text())
		if err != nil {
			continue
		}
		if current > prev {
			increases += 1
		}
		prev = current
		fmt.Printf("prev(%d) curr(%d)\n", prev, current)
	}
	fmt.Printf("There were %d increases.\n", increases)
}
