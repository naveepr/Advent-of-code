package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) []int {
	var input []int
	var str []string

	for scanner.Scan() {
		str = strings.Split(scanner.Text(), ",")
	}

	for _, val := range str {
		num, err := strconv.Atoi(val)
		check(err)
		input = append(input, num)
	}

	return input
}

func Solve(inputptr *[]int, part1 bool) int {
	input := *inputptr
	maxVal, minVal, minCost, cost := 0, math.MaxInt, math.MaxInt, 0

	for _, val := range input {
		if val > maxVal {
			maxVal = val
		}

		if val < minVal {
			minVal = val
		}
	}

	for val1 := minVal; val1 <= maxVal; val1++ {
		cost = 0
		for _, val2 := range input {
			val := val1 - val2
			if val < 0 {
				val *= -1
			}
			if part1 {
				cost += val
			} else {
				cost += (val * (val + 1)) / 2
			}
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func main() {
	arg := os.Args
	part1 := true

	if len(arg) == 3 && arg[2] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := parseInput(scanner)

	ans := Solve(&input, part1)

	fmt.Println("Answer is ", ans)
}
