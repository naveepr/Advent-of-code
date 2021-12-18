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

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func parseInput(scanner *bufio.Scanner) ([][4]int, int, int) {
	var matrix [][4]int
	var input string
	var max_x, max_y int

	for scanner.Scan() {
		regCompile := regexp.MustCompile("(.+),(.+) -> (.+),(.+)")
		input = scanner.Text()
		match := regCompile.FindStringSubmatch(input)
		x1, err := strconv.Atoi(match[1])
		check(err)
		y1, err := strconv.Atoi(match[2])
		check(err)
		x2, err := strconv.Atoi(match[3])
		check(err)
		y2, err := strconv.Atoi(match[4])
		check(err)

		max_x = max(max_x, max(x1, x2))
		max_y = max(max_y, max(y1, y2))
		matrix = append(matrix, [4]int{x1, y1, x2, y2})
	}
	return matrix, max_x, max_y
}

func Solve(inputptr *[][4]int, matrixptr *[][]int, part1 bool) int {
	input := *inputptr
	matrix := *matrixptr
	var x1, y1, x2, y2, res int

	for i := range input {
		x1 = input[i][0]
		y1 = input[i][1]
		x2 = input[i][2]
		y2 = input[i][3]

		if (x1 != x2) && (y1 != y2) {
			if !part1 {
				slope := float64((y2 - y1) / (x2 - x1))
				if math.Abs(slope) == 1 {
					if x1 < x2 {
						if slope == 1 { //+1,1
							for i, j := x1, y1; i <= x2 && j <= y2; i++ {
								matrix[j][i]++
								if matrix[j][i] == 2 {
									res++
								}
								j++
							}
						} else { //+1,-1
							for i, j := x1, y1; i <= x2 && j >= y2; i++ {
								matrix[j][i]++
								if matrix[j][i] == 2 {
									res++
								}
								j--
							}
						}
					} else {
						if slope == 1 { //-1,-1
							for i, j := x1, y1; i >= x2 && j >= y2; i-- {
								matrix[j][i]++
								if matrix[j][i] == 2 {
									res++
								}
								j--
							}
						} else { //-1,+1
							for i, j := x1, y1; i >= x2 && j <= y2; i-- {
								matrix[j][i]++
								if matrix[j][i] == 2 {
									res++
								}
								j++
							}
						}
					}
				}
			}
			continue
		}

		if (x1 == x2) && (y1 == y2) {
			matrix[x1][y1]++
			if matrix[x1][y1] == 2 {
				res++
			}
			continue

		}

		if x1 == x2 {
			for i := min(y1, y2); i <= max(y1, y2); i++ {
				matrix[i][x1]++
				if matrix[i][x1] == 2 {
					res++
				}
			}
			continue
		}

		if y1 == y2 {
			for i := min(x1, x2); i <= max(x1, x2); i++ {
				matrix[y1][i]++
				if matrix[y1][i] == 2 {
					res++
				}
			}
			continue
		}

	}
	return res
}

func main() {
	var matrix [][]int
	arg := os.Args
	part1 := true
	if len(arg) == 3 && arg[2] == "2" {
		part1 = false
	}

	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input, colsz, rowsz := parseInput(scanner)
	matrix = make([][]int, rowsz+1)
	for i := range matrix {
		matrix[i] = make([]int, colsz+1)
	}

	ans := Solve(&input, &matrix, part1)

	fmt.Println("Answer is ", ans)
}
