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

func parseInput(scanner *bufio.Scanner) []uint8 {
	input := make([]uint8, 0, math.MaxInt32)
	var str []string

	for scanner.Scan() {
		str = strings.Split(scanner.Text(), ",")
	}

	for _, val := range str {
		num, err := strconv.Atoi(val)
		check(err)
		input = append(input, uint8(num))
	}

	return input
}

func Solve(inputptr *[]uint8, days int, part1 bool) int {
	input := *inputptr
	var rem_days [9]int

	var res int
	if part1 {
		for i := 1; i <= days; i++ {
			for index, val := range input {
				if val == 0 {
					input[index] = 6
					input = append(input, 8)
				} else {
					input[index]--
				}
			}
		}
		res = len(input)
	} else {
		for _, val := range input {
			rem_days[val]++
		}

		for i := 0; i < days; i++ {
			var tempDays [9]int
			for index := range rem_days {
				if index == 0 {
					tempDays[6] += rem_days[0]
					tempDays[8] += rem_days[0]

				} else {
					tempDays[index-1] += rem_days[index]
				}
			}
			rem_days = tempDays
		}
		for _, val := range rem_days {
			res += val
		}
	}

	return res
}

func main() {
	arg := os.Args
	part1 := true

	if len(arg) < 3 {
		log.Fatal("Too few args")
	}

	days, err := strconv.Atoi(arg[2])
	check(err)
	if len(arg) == 4 && arg[3] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := parseInput(scanner)

	ans := Solve(&input, days, part1)

	fmt.Println("Answer is ", ans)
}
