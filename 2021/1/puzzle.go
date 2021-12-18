package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SolvePart1(scanner *bufio.Scanner) int {
	count, prev := 0, math.MaxInt

	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		check(err)

		if cur > prev {
			count++
		}
		prev = cur
	}
	return count
}

func SolvePart2(scanner *bufio.Scanner) int {
	count, sum, prev := 0, 0, math.MaxInt
	var window = make([]int, 0, 3)

	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		check(err)

		if len(window) < 3 {
			window = append(window, cur)
			sum += cur
		} else {
			prev = sum
			sum = sum - window[0] + cur
			if sum > prev {
				count++
			}
			window = window[1:3]
			window = append(window, cur)
		}
	}

	return count

}

func main() {
	part1 := true
	var count int
	arg := os.Args
	if len(arg) == 3 && arg[2] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if part1 {
		count = SolvePart1(scanner)
	} else {
		count = SolvePart2(scanner)
	}

	fmt.Println("count is ", count)
}
