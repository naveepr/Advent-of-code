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

func parseInput(scanner *bufio.Scanner) ([][]string, []string) {
	var str []string
	var fold []string
	var coordinates [][]int
	var rowsz, colsz int

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		s = strings.TrimSpace(s)
		str = strings.Split(s, ",")
		x, err := strconv.Atoi(str[0])
		check(err)
		y, err := strconv.Atoi(str[1])
		check(err)
		if y > rowsz {
			rowsz = y
		}
		if x > colsz {
			colsz = x
		}
		coordinates = append(coordinates, []int{x, y})
	}

	for scanner.Scan() {
		s := scanner.Text()
		s = strings.TrimSpace(s)
		str = strings.Split(s, " ")
		fold = append(fold, str[2])
	}

	matrix := make([][]string, rowsz+1)
	for i := 0; i <= rowsz; i++ {
		matrix[i] = make([]string, colsz+1)
		for j := range matrix[i] {
			matrix[i][j] = "."
		}
	}
	for _, row := range coordinates {
		matrix[row[1]][row[0]] = "#"
	}
	return matrix, fold
}

func Solve(matrix [][]string, fold []string, part1 bool) int {
	var res1, rowsz, colsz int
	var str []string

	rowsz = len(matrix)
	colsz = len(matrix[0])

	for _, f := range fold {
		str = strings.Split(f, "=")
		foldat, err := strconv.Atoi(str[1])
		check(err)
		if str[0] == "x" {
			for j := 1; j+foldat < colsz; j++ {
				for i := 0; i < rowsz; i++ {
					if matrix[i][foldat+j] == "#" {
						matrix[i][foldat-j] = matrix[i][foldat+j]
					}
				}
			}
			colsz = foldat + 1
		} else {
			for i := 1; i+foldat < rowsz; i++ {
				for j := 0; j < colsz; j++ {
					if matrix[foldat+i][j] == "#" {
						matrix[foldat-i][j] = matrix[foldat+i][j]
					}
				}
			}
			rowsz = foldat + 1
		}

		if part1 {
			break
		}
	}

	for i := 0; i < rowsz; i++ {
		for j := 0; j < colsz; j++ {
			fmt.Print(matrix[i][j])
			fmt.Print(" ")
		}
		fmt.Println("")
	}

	for i := 0; i < rowsz; i++ {
		for j := 0; j < colsz; j++ {
			if matrix[i][j] == "#" {
				res1++
			}
		}
	}

	return res1
}

func main() {
	arg := os.Args
	part1 := true

	if len(arg) == 3 && arg[2] == "part2" {
		part1 = false
	}
	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input, fold := parseInput(scanner)

	fmt.Println(fold)
	ans := Solve(input, fold, part1)

	fmt.Println("Answer is ", ans)
}
