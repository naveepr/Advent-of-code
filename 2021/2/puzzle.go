package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SolvePart1(scanner *bufio.Scanner) int {
	x, y := 0, 0

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		dis, err := strconv.Atoi(s[1])
		check(err)

		switch s[0] {
		case "forward":
			x += dis
		case "up":
			y -= dis
		case "down":
			y += dis
		default:
			log.Fatal("Input File unexpected entry")
		}
	}
	return x * y
}

func SolvePart2(scanner *bufio.Scanner) int {
	x, y, aim := 0, 0, 0

	for scanner.Scan() {
		s := strings.Fields(scanner.Text())
		dis, err := strconv.Atoi(s[1])
		check(err)

		switch s[0] {
		case "forward":
			x += dis
			y += (aim * dis)
		case "up":
			aim -= dis
		case "down":
			aim += dis
		default:
			log.Fatal("Input File unexpected entry")
		}
	}
	return x * y

}

func main() {
	part1 := true
	var ans int
	arg := os.Args
	if len(arg) == 3 && arg[2] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	if part1 {
		ans = SolvePart1(scanner)
	} else {
		ans = SolvePart2(scanner)
	}

	fmt.Println("Answer is ", ans)
}
