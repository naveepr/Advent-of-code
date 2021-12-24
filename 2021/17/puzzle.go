package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func parseInput(scanner *bufio.Scanner) ([]int, []int) {
	var x, y []int
	var x1, x2, y1, y2 int
	var err error

	for scanner.Scan() {
		s := scanner.Text()
		reg := regexp.MustCompile("x=(-?[0-9]+)..(-?[0-9]+), y=(-?[0-9]+)..(-?[0-9]+)")
		match := reg.FindStringSubmatch(s)
		fmt.Println(match)
		x1, err = strconv.Atoi(match[1])
		check(err)
		x2, err = strconv.Atoi(match[2])
		check(err)
		y1, err = strconv.Atoi(match[3])
		check(err)
		y2, err = strconv.Atoi(match[4])
		check(err)

		x = append(x, x1, x2)
		y = append(y, y1, y2)
	}
	return x, y
}

func Solve(x, y []int) int {
	var res, val, xmin, xmax, ymin, ymax int

	val = max(abs(y[0]), abs(y[1]))
	ymin = min(y[0], y[1])
	ymax = (val - 1)
	xmax = x[1]

	fval := math.Sqrt(float64(x[0] * 2))
	val = int(math.Floor(fval))

	for val*(val+1)/2 < x[0] {
		val++
	}

	xmin = val

	fmt.Println(xmin, xmax, ymin, ymax)
	for j := ymin; j <= ymax; j++ {
		for i := xmin; i <= xmax; i++ {
			startx, starty := 0, 0
			dx, dy := i, j
			for startx <= max(x[1], x[0]) && starty >= min(y[1], y[0]) {
				startx += dx
				starty += dy

				if dx > 0 {
					dx--
				} else if dx < 0 {
					dx++
				}
				dy--

				if startx >= min(x[0], x[1]) && startx <= max(x[0], x[1]) && starty <= max(y[0], y[1]) && starty >= min(y[0], y[1]) {
					res++
					break
				}
			}
		}
	}
	return res
}

func main() {
	arg := os.Args
	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	x, y := parseInput(scanner)
	fmt.Println(x, y)
	val := min(y[0], y[1])
	fmt.Println(val * (val + 1) / 2)

	fmt.Println(Solve(x, y))
}
