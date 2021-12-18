package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseInput(scanner *bufio.Scanner) (map[string]string, string) {
	var input string
	pol_template := make(map[string]string)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		input = s
	}

	for scanner.Scan() {
		s := scanner.Text()
		regCompile := regexp.MustCompile("(.+) -> (.+)")
		match := regCompile.FindStringSubmatch(s)
		pol_template[match[1]] = match[2]
	}

	return pol_template, input
}

func Solve(input string, pol_template map[string]string) uint64 {
	var res uint64
	var tempStr string
	charMap := make(map[rune]uint64)
	var maxVal, minVal uint64

	// fmt.Println(input)
	maxVal, minVal = 0, math.MaxUint64
	for i := 0; i < 10; i++ {
		tempStr = ""
		l := len(input)
		for j := 0; j <= l-2; j++ {
			substr := input[j : j+2]
			tempStr += input[j : j+1]
			if s, found := pol_template[substr]; found {
				tempStr += s
			}
			// fmt.Println(tempStr)
		}
		tempStr += input[l-1:]
		input = tempStr
		// fmt.Println(input)
	}

	for _, r := range input {
		charMap[r]++
		if charMap[r] > maxVal {
			maxVal = charMap[r]
		}
	}

	for _, v := range charMap {
		if v < minVal {
			minVal = v
		}

	}
	res = maxVal - minVal
	return res
}

func SolvePart2(input string, pol_template map[string]string) uint64 {
	var res uint64
	c1 := make(map[string]uint64)
	var maxVal, minVal uint64

	// fmt.Println(input)
	maxVal, minVal = 0, math.MaxUint64

	l := len(input)
	for j := 0; j <= l-2; j++ {
		substr := input[j : j+2]
		if _, found := pol_template[substr]; found {
			c1[substr]++
		}
	}

	for i := 0; i < 40; i++ {
		c2 := make(map[string]uint64)
		for k, v := range c1 {
			c2[k[0:1]+pol_template[k]] += v
			c2[pol_template[k]+k[1:2]] += v
		}
		c1 = make(map[string]uint64)
		for k, v := range c2 {
			c1[k] = v
		}
	}

	charMap := make(map[byte]uint64)

	for k, v := range c1 {
		charMap[k[0]] += v
		// charMap[k[1]] += v
	}
	charMap[input[l-1]]++
	for _, v := range charMap {
		if v > maxVal {
			maxVal = v
		}
		if v < minVal {
			minVal = v
		}
	}
	res = maxVal - minVal
	return res
}

func main() {
	arg := os.Args
	// part1 := true

	// if len(arg) == 3 && arg[2] == "part2" {
	// 	part1 = false
	// }
	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	pol_template, input := parseInput(scanner)
	ans1 := Solve(input, pol_template)
	ans2 := SolvePart2(input, pol_template)

	fmt.Println("Answer is ", ans1, ans2)
}
