package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func simple(commands []string) int {
	depth := 0
	horizontal := 0

	for _, command := range commands {
		split := strings.Split(command, " ")
		action := split[0]
		value, _ := strconv.Atoi(split[1])
		if action == "forward" {
			horizontal += value
		} else if action == "up" {
			depth -= value
		} else if action == "down" {
			depth += value
		}
	}
	return depth * horizontal
}

func aim(commands []string) int {
	var depth, horizontal, aim = 0, 0, 0
	for _, command := range commands {
		split := strings.Split(command, " ")
		action := split[0]
		value, _ := strconv.Atoi(split[1])
		if action == "forward" {
			horizontal += value
			depth += aim * value
		} else if action == "up" {
			aim -= value
		} else if action == "down" {
			aim += value
		}
	}
	return depth * horizontal
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var commands []string
	for scanner.Scan() {
		command := scanner.Text()
		commands = append(commands, command)
	}

	fmt.Printf("Simple: %d\n", simple(commands))
	fmt.Printf("Complex: %d\n", aim(commands))
}
